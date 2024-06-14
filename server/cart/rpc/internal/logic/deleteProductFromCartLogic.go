package logic

import (
	"GoBao/common/xerr"
	"GoBao/server/cart/model"
	"GoBao/server/cart/rpc/internal/svc"
	"GoBao/server/cart/rpc/pb"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductFromCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductFromCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductFromCartLogic {
	return &DeleteProductFromCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteProductFromCartLogic) DeleteProductFromCart(in *pb.DeleteProductFromCartRequest) (*pb.DeleteProductFromCartResponse, error) {
	// todo: add your logic here and delete this line

	var cart model.CartProduct
	//查找是否有该购物车
	err := l.svcCtx.CartDB.Where("id=?", in.CartID).Take(&cart).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.CART_NOT_EXISTS_ERROR), "CartID:%+v,UserID:%+v", in.CartID, in.UserID)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE cart ERROR:%+v", err)
	}

	err = l.svcCtx.CartDB.Delete(&cart).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL DELETE cart ERROR:%+v", err)
	}

	return &pb.DeleteProductFromCartResponse{}, nil
}
