package logic

import (
	"GoBao/common/xerr"
	"GoBao/server/cart/model"
	"GoBao/server/cart/rpc/internal/svc"
	"GoBao/server/cart/rpc/pb"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"time"
)

type UpdateCartProductDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCartProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCartProductDetailLogic {
	return &UpdateCartProductDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCartProductDetailLogic) UpdateCartProductDetail(in *pb.UpdateCartProductDetailRequest) (*pb.UpdateCartProductDetailResponse, error) {
	// todo: add your logic here and delete this line

	//判断是否有该购物车信息
	var CartProduct model.CartProduct
	err := l.svcCtx.CartDB.Where("id=? and user_id=?", in.CartID, in.UserID).Take(&CartProduct).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.CART_NOT_EXISTS_ERROR), "cart not exist,UserID:%+v,CartID:%+v", in.UserID, in.CartID)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE cart product ERROR:%+v", err)
	}

	//有购物车，执行改购物车信息的逻辑
	CartProduct.Quantity = in.Count
	CartProduct.UpdateTime = time.Now()
	err = l.svcCtx.CartDB.Model(&CartProduct).Updates(CartProduct).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL UPDATE cart product ERROR:%+v", err)
	}

	return &pb.UpdateCartProductDetailResponse{}, nil
}
