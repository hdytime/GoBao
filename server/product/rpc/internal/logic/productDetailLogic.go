package logic

import (
	"GoBao/common/xerr"
	"GoBao/server/product/model"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"GoBao/server/product/rpc/internal/svc"
	"GoBao/server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductDetailLogic {
	return &ProductDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductDetailLogic) ProductDetail(in *pb.ProductDetailReq) (*pb.ProductDetailResp, error) {
	// todo: add your logic here and delete this line

	//使用singleflight来预防缓存击穿
	v, err, _ := l.svcCtx.SingleGroup.Do(fmt.Sprintf("product_id:%d", in.ProductID), func() (interface{}, error) {
		var p model.Product
		err := l.svcCtx.ProductDB.Where("id=?", in.ProductID).Take(&p).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, errors.Wrapf(xerr.NewErrCode(xerr.PRODUCT_NOT_EXISTS_ERROR), "ProductID:%v", in.ProductID)
			}
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE product ERROR:%+v,productID:%v", err, in.ProductID)
		}
		var product = &pb.Product{
			ID:         p.Id,
			Name:       p.Name,
			Price:      p.Price,
			Stock:      p.Stock,
			Status:     p.Status,
			CreateTime: p.CreateTime.Unix(),
			UpdateTime: p.UpdateTime.Unix(),
		}
		return product, nil
	})
	if err != nil {
		return nil, err
	}

	return &pb.ProductDetailResp{Product: v.(*pb.Product)}, nil
}
