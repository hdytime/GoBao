package logic

import (
	"GoBao/common/globalKey"
	"GoBao/common/xerr"
	"GoBao/server/product/model"
	"GoBao/server/product/rpc/internal/svc"
	"GoBao/server/product/rpc/pb"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecommendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRecommendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecommendLogic {
	return &RecommendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// commonProduct
func (l *RecommendLogic) Recommend(in *pb.RecommendReq) (*pb.RecommendResp, error) {
	// todo: add your logic here and delete this line

	// 查redis缓存
	resp, err := l.svcCtx.RedisDB.Get(l.ctx, globalKey.Recommend).Result()
	if err == nil {
		//查到缓存
		var p []*pb.SmallProduct
		err := json.Unmarshal([]byte(resp), &p)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "json.Unmarshal ERROR:%+v", err)
		}
		return &pb.RecommendResp{SmallProducts: p}, nil
	}
	//没在redis中查到缓存，走mysql
	var ids []int64
	err = l.svcCtx.ProductDB.Model(&model.ProductRecommend{}).Select("product_id").Scan(&ids).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL SCAN product recommmend ERROR:%+v", err)
	}

	var ps []model.Product
	err = l.svcCtx.ProductDB.Where("id=?", ids).Find(&ps).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL FIND product ERROR:%+v", err)
	}

	var res []*pb.SmallProduct
	for _, p := range ps {
		var sp = pb.SmallProduct{
			ID:            p.Id,
			Name:          p.Name,
			Price:         p.Price,
			DiscountPrice: p.DiscountPrice,
		}
		res = append(res, &sp)
	}

	//mysql走完，写入redis当缓存
	jsonres, _ := json.Marshal(&res)
	resCmd := l.svcCtx.RedisDB.Set(l.ctx, globalKey.Recommend, string(jsonres), 24*time.Hour)
	if resCmd.Err() != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_REDIS_ERROR), "REDIS SET product recommend ERROR:%+v", err)
	}
	return &pb.RecommendResp{
		SmallProducts: res,
	}, nil
}
