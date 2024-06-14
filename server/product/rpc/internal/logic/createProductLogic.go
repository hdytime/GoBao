package logic

import (
	"GoBao/common/globalKey"
	"GoBao/common/xerr"
	"GoBao/server/product/model"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"

	"GoBao/server/product/rpc/internal/svc"
	"GoBao/server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductLogic {
	return &CreateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// storeProduct
func (l *CreateProductLogic) CreateProduct(in *pb.CreateProductReq) (*pb.CreateProductResp, error) {
	// todo: add your logic here and delete this line

	var p model.Product
	//判断是否已经有同名商品
	err := l.svcCtx.ProductDB.Where("name=?", in.Name).Take(&p).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE create product ERROR:%+v", err)
	} else if err == nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.PRODUCT_EXISTS_ERROR), "Product:%+v", in.Name)
	}

	//没有同名商品，增添商品
	err = l.svcCtx.ProductDB.Create(&model.Product{
		Id:            l.svcCtx.SnowflakeNode.Generate().Int64(),
		Name:          in.Name,
		Price:         in.Price,
		DiscountPrice: in.Price,
		Stock:         in.Stock,
		Status:        globalKey.ProductOnline,
		CreateTime:    time.Now(),
		UpdateTime:    time.Now(),
	}).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL CREATE product ERROR:%+v", err)
	}

	return &pb.CreateProductResp{}, nil
}
