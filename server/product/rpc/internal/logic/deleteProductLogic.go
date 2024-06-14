package logic

import (
	"GoBao/common/xerr"
	"GoBao/server/product/model"
	"GoBao/server/product/rpc/internal/svc"
	"GoBao/server/product/rpc/pb"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductLogic {
	return &DeleteProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteProductLogic) DeleteProduct(in *pb.DeleteProductReq) (*pb.DeleteProductResp, error) {
	// todo: add your logic here and delete this line

	var products []model.Product
	//检查传入的商品是否都存在
	err := l.svcCtx.ProductDB.Where("id=?", in.ProductIDs).Find(&products).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL FIND products ERROR:%+v", err)
	}
	fmt.Print(products)
	if len(products) != len(in.ProductIDs) {
		return nil, errors.Wrapf(xerr.NewErrMsg("选择删除的商品有误，请重新选择"), "Delete Products Req:%+v", in)
	}

	//更新数据
	err = l.svcCtx.ProductDB.Transaction(func(tx *gorm.DB) error {
		err := tx.Delete(&products).Error
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL DELETE products ERROR:%+v", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &pb.DeleteProductResp{}, nil
}
