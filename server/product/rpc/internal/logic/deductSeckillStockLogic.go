package logic

import (
	"GoBao/common/globalKey"
	"GoBao/common/xerr"
	"GoBao/server/product/rpc/productrpc"
	"context"
	"github.com/pkg/errors"
	"math/rand"
	"strconv"
	"time"

	"GoBao/server/product/rpc/internal/svc"
	"GoBao/server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeductSeckillStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeductSeckillStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeductSeckillStockLogic {
	return &DeductSeckillStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeductSeckillStockLogic) DeductSeckillStock(in *pb.DeductSeckillStockReq) (*pb.DeductSeckillStockResp, error) {
	// todo: add your logic here and delete this line

	//从redis读取库存
	stockKey := globalKey.SeckillProductStock + strconv.FormatInt(in.SeckillProductID, 10)
	stockstring, err := l.svcCtx.RedisDB.Get(l.ctx, stockKey).Result()
	if err != nil {
		if stockstring != "" {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_REDIS_ERROR), "REDIS GET stock ERROR:%+v", err)
		}
		//redis中无库存，查询数据库后写入redis
		detail, err := l.svcCtx.ProductRpc.SeckillDetail(l.ctx, &productrpc.SeckillDetailReq{SeckillProductID: in.SeckillProductID})
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrMsg("ProductRpc ERROR"), "ProductRpc ERROR:%+v")
		}

		stockAfter := detail.SeckillProduct.Stock - in.Stock
		err = l.svcCtx.RedisDB.Set(l.ctx, stockKey, stockAfter, time.Hour+time.Duration(rand.Intn(60))*time.Minute).Err()
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_REDIS_ERROR), "REDIS set stock ERROR:%+v", err)
		}
	}
	stock, _ := strconv.ParseInt(stockstring, 10, 64)
	if stock > in.Stock {
		//更新redis缓存
		stockAfter := stock - in.Stock
		err := l.svcCtx.RedisDB.Set(l.ctx, stockKey, stockAfter, time.Hour+time.Duration(rand.Intn(60))*time.Minute).Err()
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_REDIS_ERROR), "REDIS set stock ERROR:%+v", err)
		}
	}
	return &pb.DeductSeckillStockResp{}, nil
}
