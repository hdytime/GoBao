package logic

import (
	"GoBao/common/batcher"
	"GoBao/common/globalKey"
	"GoBao/common/xerr"
	"GoBao/server/mq/job/jobtype"
	"GoBao/server/order/model"
	"GoBao/server/product/rpc/productrpc"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/mr"
	"strconv"
	"strings"
	"time"

	"GoBao/server/order/rpc/internal/svc"
	"GoBao/server/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSeckillOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	batcher *batcher.Batcher
}

func NewCreateSeckillOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSeckillOrderLogic {
	f := &CreateSeckillOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
	//batcher配置
	options := batcher.Options{
		5,
		100,
		100,
		5 * time.Second,
	}
	// 实现batcher
	b := batcher.New(options)
	b.Sharding = func(key string) int {
		pid, _ := strconv.ParseInt(key, 10, 64)
		return int(pid) % options.Worker
	}
	b.Do = func(ctx context.Context, val map[string][]interface{}) {
		var msgs []*model.Order
		for _, vs := range val {
			for _, v := range vs {
				msgs = append(msgs, v.(*model.Order))
			}
		}
		kd, err := json.Marshal(msgs)
		if err != nil {
			logx.Errorf("Batcher.Do json.Marshal msgs: %v error: %v", msgs, err)
		}
		if err = f.svcCtx.KqPusherClient.Push(string(kd)); err != nil {
			logx.Errorf("KafkaPusher.Push kd: %s error: %v", string(kd), err)
		}
	}
	f.batcher = b
	f.batcher.Start()
	return f
}

func (l *CreateSeckillOrderLogic) CreateSeckillOrder(in *pb.CreateSeckillOrderReq) (*pb.CreateSeckillOrderResp, error) {
	// todo: add your logic here and delete this line
	var seckillDetail *productrpc.SeckillDetailResp
	//mr并发判断参数条件
	checkFuncs := []func() error{
		func() error {
			//判断商品是否是秒杀商品
			seckillDetail, err := l.svcCtx.ProductRpc.SeckillDetail(l.ctx, &productrpc.SeckillDetailReq{
				SeckillProductID: in.ProductID,
			})
			if err != nil {
				return err
			} else if seckillDetail.SeckillProduct == nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.PRODUCT_SECKILL_NOT_EXISTS_ERROR), "seckillID:%+v", in.ProductID)
			}
			return nil
		},
		func() error {
			//判断下单时间是否在秒杀时间段内
			seckillDetail, err := l.svcCtx.ProductRpc.SeckillDetail(l.ctx, &productrpc.SeckillDetailReq{
				SeckillProductID: in.ProductID,
			})
			if err != nil {
				return err
			} else if seckillDetail.SeckillProduct == nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.PRODUCT_SECKILL_NOT_EXISTS_ERROR), "seckillID:%+v", in.ProductID)
			}
			startTime := strings.Split(seckillDetail.SeckillProduct.StartTime, " ")[0]
			day := time.Now().Format("2006-01-02")
			hour := int64(time.Now().Hour())
			if day != startTime || hour-seckillDetail.SeckillProduct.Time > 2 || hour-seckillDetail.SeckillProduct.Time < 0 {
				return errors.Wrapf(xerr.NewErrCode(xerr.PRODUCT_SECKILL_NOT_EXISTS_ERROR), "ProductID:%+v,Time:%+v", in.ProductID, time.Now())
			}
			return nil
		},
		func() error {
			//判断用户是否重复下单
			doublekey := fmt.Sprintf("%s_%d", globalKey.DoubleOrder, in.ProductID)
			isDouble, err := l.svcCtx.RedisDB.SIsMember(l.ctx, doublekey, in.UserID).Result()
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_REDIS_ERROR), "REDIS SIsMember Fail,key:%+v,ERROR:%+v", doublekey, err)
			} else if isDouble {
				return errors.Wrapf(xerr.NewErrCode(xerr.ORDER_DOUBLE_SECKILL_ERROR), "userID:%+v,seckillID:%+v", in.UserID, in.ProductID)
			}
			return nil
		},
		func() error {
			//查库存是否充足
			//seckillStockKey := fmt.Sprintf("%s_%s_%d", globalKey.SeckillCount, day, seckillDetail.SeckillProduct.Time)
			//seckillStockField := fmt.Sprintf("%d", seckillDetail.SeckillProduct.ID)
			seckillStockKey := globalKey.SeckillProductStock + strconv.FormatInt(in.ProductID, 10)
			stockstring, err := l.svcCtx.RedisDB.Get(l.ctx, seckillStockKey).Result()
			stock, _ := strconv.ParseInt(stockstring, 10, 64)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_REDIS_ERROR), "REDIS GET Fail,Key:%+v,ERROR:%+v", seckillStockKey, err)
			} else if stock <= 0 {
				return errors.Wrapf(xerr.NewErrCode(xerr.ORDER_SECKILL_STOCK_ERROR), "")
			}
			return nil
		},
	}
	err := mr.Finish(checkFuncs...)
	if err != nil {
		return nil, err
	}

	var order = model.Order{
		Id:          l.svcCtx.SnowflakeNode.Generate().Int64(),
		UserId:      in.UserID,
		ProductId:   seckillDetail.SeckillProduct.ID,
		ProductName: seckillDetail.SeckillProduct.Name,
		UnitPrice:   seckillDetail.SeckillProduct.Price,
		Quantity:    in.ProductCount,
		TotalPrice:  seckillDetail.SeckillProduct.Price * float64(in.ProductCount),
		Status:      globalKey.OrderWaitPay,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	}

	////消息队列限流削峰
	//jsonOrder, err := json.Marshal(order)
	//if err != nil {
	//	return nil, errors.Wrapf(xerr.NewErrMsg("JSON MARSHAL order ERROR"), "JSON MARSHAL order ERROR:%+v", err)
	//}
	//err = l.svcCtx.KqPusherClient.Push(string(jsonOrder))
	//if err != nil {
	//	return nil, errors.Wrapf(xerr.NewErrMsg("KQ PUSH order ERROR"), "KQ PUSH order ERROR:%+v", err)
	//}

	//使用batcher来聚合数据，减少网络和磁盘io，提高系统吞吐量
	err = l.batcher.Add(strconv.FormatInt(in.ProductID, 10), &order)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("kafka异步处理订单 ERROR"), "kafka异步处理订单 ERROR:%+v", err)
	}

	//延迟队列来处理过期订单
	payload, err := json.Marshal(jobtype.DeferCloseSeckillOrderPayload{
		UserID:  in.UserID,
		OrderSn: order.OrderSn,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("CREATE defer close seckill order task json.Marshal Fail,ERROR:%+v", err)
	} else {
		_, err := l.svcCtx.AsynqClient.Enqueue(asynq.NewTask(jobtype.DeferCloseSeckillOrder, payload), asynq.ProcessIn(globalKey.CloseSeckillOrderTimeMinutes*time.Minute))
		if err != nil {
			logx.WithContext(l.ctx).Errorf("CREATE defer close seckill order task AsynqClient.Enqueue, OrderSn:%+v,ERROR:%+v", order.OrderSn, err)
		}
	}

	return &pb.CreateSeckillOrderResp{OrderSn: order.OrderSn}, nil
}
