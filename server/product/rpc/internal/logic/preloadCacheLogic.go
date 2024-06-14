package logic

import (
	"GoBao/common/globalKey"
	"GoBao/common/xerr"
	"GoBao/server/product/model"
	"GoBao/server/product/rpc/internal/svc"
	"GoBao/server/product/rpc/pb"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"math/rand"
	"strconv"
	"time"
)

type PreloadCacheLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPreloadCacheLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PreloadCacheLogic {
	return &PreloadCacheLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PreloadCacheLogic) PreloadCache(in *pb.PreloadCacheReq) (*pb.PreloadCacheResp, error) {
	// todo: add your logic here and delete this line

	//mysql查询商品详情
	var seckillProducts []model.SeckillProduct
	err := l.svcCtx.ProductDB.Find(&seckillProducts).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL FIND seckillProducts ERROR:%+v", err)
	}

	//缓存预热
	for _, seckillProduct := range seckillProducts {
		//缓存商品库存
		productStockKey := globalKey.SeckillProductStock + strconv.FormatInt(seckillProduct.Id, 10)
		err := l.svcCtx.RedisDB.Set(l.ctx, productStockKey, seckillProduct.Stock, time.Duration(seckillProduct.Time)).Err()
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_REDIS_ERROR), "REDIS SET seckillProductStock ERROR:%+v", err)
		}
		//缓存商品详情
		productDetailKey := globalKey.SeckillProductDetail + strconv.FormatInt(seckillProduct.Id, 10)
		productDetail := map[string]interface{}{
			"ID":           seckillProduct.Id,
			"Name":         seckillProduct.Name,
			"Price":        seckillProduct.Price,
			"Stock":        seckillProduct.Stock,
			"Status":       seckillProduct.Status,
			"CreateTime":   seckillProduct.CreateTime,
			"UpdateTime":   seckillProduct.UpdateTime,
			"SeckillPrice": seckillProduct.SeckillPrice,
			"StockCount":   seckillProduct.StockCount,
			"StartTime":    seckillProduct.StartTime,
			"Time":         seckillProduct.Time,
		}
		err = l.svcCtx.RedisDB.HMSet(l.ctx, productDetailKey, productDetail).Err()
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_REDIS_ERROR), "REDIS HMSet seckillProductDetail ERROR:%+v", err)
		}
		//缓存加上随机过期时间，避免大量缓存同时失效引起缓存击穿
		l.svcCtx.RedisDB.Expire(l.ctx, productDetailKey, time.Hour+time.Duration(rand.Intn(60))*time.Minute)
	}
	return &pb.PreloadCacheResp{
		Success: true,
	}, nil
}
