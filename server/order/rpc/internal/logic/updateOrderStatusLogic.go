package logic

import (
	"GoBao/common/xerr"
	"GoBao/server/order/model"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"

	"GoBao/server/order/rpc/internal/svc"
	"GoBao/server/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderStatusLogic {
	return &UpdateOrderStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOrderStatusLogic) UpdateOrderStatus(in *pb.UpdateOrderStatusReq) (*pb.UpdateOrderStatusResp, error) {
	// todo: add your logic here and delete this line

	//查询是否有该订单
	var order model.Order
	err := l.svcCtx.OrderDB.Where("user_id=? and order_sn=?", in.UserID, in.OrderSn).Take(&order).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.ORDER_NOT_EXISTS_ERROR), "UserID:%+v,OrderSn:%+v", in.UserID, in.OrderSn)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE order ERROR:%+v", err)
	}

	//有订单修改订单状态
	order.Status = in.Status
	order.UpdateTime = time.Now()
	err = l.svcCtx.OrderDB.Updates(&order).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL UPDATES order ERROR:%+v", err)
	}
	return &pb.UpdateOrderStatusResp{}, nil
}
