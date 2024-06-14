package logic

import (
	"GoBao/common/globalKey"
	"GoBao/common/tool"
	"GoBao/common/xerr"
	"GoBao/server/product/model"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"time"

	"GoBao/server/product/rpc/internal/svc"
	"GoBao/server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SeckillDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSeckillDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeckillDetailLogic {
	return &SeckillDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SeckillDetailLogic) SeckillDetail(in *pb.SeckillDetailReq) (*pb.SeckillDetailResp, error) {
	// todo: add your logic here and delete this line

	//singleflight预防缓存击穿
	v, err, _ := l.svcCtx.SingleGroup.Do(fmt.Sprintf("seckill_product_id:%d", in.SeckillProductID), func() (interface{}, error) {
		cacheKey := globalKey.SeckillProductDetail + strconv.FormatInt(in.SeckillProductID, 10)
		ProductDetails, err := l.svcCtx.RedisDB.HGetAll(l.ctx, cacheKey).Result()
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_REDIS_ERROR), "REDIS HGETALL seckillProductDetail ERROR:%+v", err)
		}
		//检查缓存是否命中
		if len(ProductDetails) == 0 {
			//redis没查到缓存，走mysql去查数据库
			var seckill model.SeckillProduct
			err := l.svcCtx.ProductDB.Where("id=?", in.SeckillProductID).Take(&seckill).Error
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					return nil, errors.Wrapf(xerr.NewErrCode(xerr.PRODUCT_SECKILL_NOT_EXISTS_ERROR), "productID:%v", in.SeckillProductID)
				}
				return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE seckillProduct ERROR:%+v", err)
			}
			//回写到redis缓存中
			productDetail := map[string]interface{}{
				"Id":           seckill.Id,
				"Name":         seckill.Name,
				"Price":        seckill.Name,
				"Stock":        seckill.Stock,
				"Status":       seckill.Status,
				"CreateTime":   seckill.CreateTime,
				"UpdateTime":   seckill.UpdateTime,
				"SeckillPrice": seckill.SeckillPrice,
				"StockCount":   seckill.StockCount,
				"StartTime":    seckill.StartTime,
				"Time":         seckill.Time,
			}
			for field, value := range productDetail {
				//使用HSetNX而不是Hset避免并发读写导致缓存不一致
				err := l.svcCtx.RedisDB.HSetNX(l.ctx, cacheKey, field, value).Err()
				if err != nil {
					return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_REDIS_ERROR), "REDIS HSetNX ProductDetail ERROR:%+v", err)
				}
			}
			//加上随机过期时间防止缓存击穿
			err = l.svcCtx.RedisDB.Expire(l.ctx, cacheKey, time.Hour+time.Duration(rand.Intn(60))*time.Minute).Err()
			if err != nil {
				return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_REDIS_ERROR), "REDIS EXPIRE ProductDetail ERROR:%+v", err)
			}

			var res = pb.SeckillProduct{
				ID:           seckill.Id,
				Name:         seckill.Name,
				Price:        seckill.Price,
				Stock:        seckill.Stock,
				Status:       seckill.Status,
				CreateTime:   seckill.CreateTime.Unix(),
				UpdateTime:   seckill.UpdateTime.Unix(),
				SeckillPrice: seckill.SeckillPrice,
				SeckillCount: seckill.StockCount,
				StartTime:    seckill.StartTime,
				Time:         seckill.Time,
			}

			return res, nil
		} else {
			//redis查到缓存，直接返回结果
			var res pb.SeckillProduct

			res.ID = tool.ConvertStringToInt64(ProductDetails["Id"])
			res.Name = ProductDetails["Name"]
			res.Price = tool.ConvertStringToFloat64(ProductDetails["Price"])
			res.Stock = tool.ConvertStringToInt64(ProductDetails["Stock"])
			res.Status = tool.ConvertStringToInt64(ProductDetails["Status"])
			res.CreateTime = tool.ConvertStringToTimeUnix(ProductDetails["CreateTime"])
			res.UpdateTime = tool.ConvertStringToTimeUnix(ProductDetails["UpdateTime"])
			res.SeckillPrice = tool.ConvertStringToFloat64(ProductDetails["SeckillPrice"])
			res.SeckillCount = tool.ConvertStringToInt64(ProductDetails["SeckillCount"])
			res.StartTime = ProductDetails["StartTime"]
			res.Time = tool.ConvertStringToInt64(ProductDetails["Time"])
			return res, nil
		}
	})
	if err != nil {
		return nil, err
	}
	return &pb.SeckillDetailResp{SeckillProduct: v.(*pb.SeckillProduct)}, nil
}
