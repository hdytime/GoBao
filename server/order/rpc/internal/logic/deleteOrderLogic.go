package logic

import (
	"GoBao/common/xerr"
	"GoBao/server/order/model"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"GoBao/server/order/rpc/internal/svc"
	"GoBao/server/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOrderLogic {
	return &DeleteOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteOrderLogic) DeleteOrder(in *pb.DeleteOrderReq) (*pb.DeleteOrderResp, error) {
	// todo: add your logic here and delete this line

	//查询是否有该订单
	var order model.Order
	err := l.svcCtx.OrderDB.Where("order_sn=? and user_id=?", in.OrderSn, in.UserID).Take(&order).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.ORDER_NOT_EXISTS_ERROR), "ordersn:%+v,userid:%+v,ERROR:%+v", in.OrderSn, in.UserID, err)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE order ERROR:%+v", err)
	}

	//有该订单的话，删除
	err = l.svcCtx.OrderDB.Where("order_sn=? and user_id=?", order.OrderSn, order.UserId).Delete(&order).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL DELETE order ERROR:%+v", err)
	}

	return &pb.DeleteOrderResp{}, nil
}
