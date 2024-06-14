package logic

import (
	"GoBao/common/globalKey"
	"GoBao/common/xerr"
	"GoBao/server/order/model"
	"GoBao/server/pay/rpc/pay"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"

	"GoBao/server/order/rpc/internal/svc"
	"GoBao/server/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderDetailLogic {
	return &GetOrderDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderDetailLogic) GetOrderDetail(in *pb.GetOrderDetailReq) (*pb.GetOrderDetailResp, error) {
	// todo: add your logic here and delete this line

	var order model.Order
	err := l.svcCtx.OrderDB.Where("user_id=? and order_sn=?", in.UserID, in.OrderSn).Take(&order).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.ORDER_NOT_EXISTS_ERROR), "UserID:%+v,OrderSn:%+v", in.UserID, in.OrderSn)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE order ERROR:%+v", err)
	}

	var res = &pb.GetOrderDetailResp{
		ID:          order.Id,
		UserID:      order.UserId,
		ProductID:   order.ProductId,
		OrderSn:     order.OrderSn,
		ProductName: order.ProductName,
		UnitPrice:   order.UnitPrice,
		Quantity:    order.Quantity,
		TotalPrice:  order.TotalPrice,
		Status:      order.Status,
		Paytime:     0,
		Createtime:  order.CreateTime.Unix(),
		Updatetime:  order.UpdateTime.Unix(),
		Remark:      order.Remark,
	}

	if order.Status >= globalKey.OrderPayed {
		// 已经支付过会有支付信息
		detail, err := l.svcCtx.PayRpc.GetPaymentDetail(l.ctx, &pay.GetPaymentDetailReq{OrderSn: order.OrderSn})
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrMsg("Order Use PayRpc.GetPaymentDetail Fail"),
				"Order Use PaymentRpc.GetPaymentDetail Fail,orderSn:%v,ERROR:%+v", order.OrderSn, err)
		}
		paytime, _ := time.Parse("2006-01-02 15:04:05", detail.PayTime)
		res.Paytime = paytime.Unix()
	}
	return res, nil
}
