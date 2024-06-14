package logic

import (
	"GoBao/common/globalKey"
	"GoBao/common/kqOrder"
	"GoBao/common/tool/sn"
	"GoBao/common/xerr"
	"GoBao/server/order/model"
	"GoBao/server/order/rpc/order"
	"GoBao/server/user/rpc/user"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"

	model2 "GoBao/server/pay/model"
	"GoBao/server/pay/rpc/internal/svc"
	"GoBao/server/pay/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderPaymentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderPaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderPaymentLogic {
	return &OrderPaymentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderPaymentLogic) OrderPayment(in *pb.OrderPaymentReq) (*pb.OrderPaymentResp, error) {
	// todo: add your logic here and delete this line

	//判断订单是否存在
	detail, err := l.svcCtx.OrderRpc.GetOrderDetail(l.ctx, &order.GetOrderDetailReq{
		UserID:  in.UserID,
		OrderSn: in.OrderSn,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("payRpc use OrderRpc ERROR"), "payRpc use OrderRpc ERROR:%+v,userid:%v,ordersn:%v", err, in.UserID, in.OrderSn)
	}

	//判断订单状态
	if detail.Status != globalKey.OrderWaitPay {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.PAYMENT_ORDER_STATUS_ERROR), "userID:%v,orderSn:%v,price:%v", in.UserID, in.OrderSn, detail.TotalPrice)
	}

	var order = model.Order{
		Id:          detail.ID,
		UserId:      detail.UserID,
		ProductId:   detail.ProductID,
		ProductName: detail.ProductName,
		OrderSn:     detail.OrderSn,
		UnitPrice:   detail.UnitPrice,
		Quantity:    detail.Quantity,
		TotalPrice:  detail.TotalPrice,
		Status:      detail.Status,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
		Remark:      detail.Remark,
	}

	PaySn, err := l.walletPay(order, in.UserID)
	if err != nil {
		return nil, err
	}

	return &pb.OrderPaymentResp{
		PayTotalPrice: order.TotalPrice,
		PaySn:         PaySn,
	}, nil

}

func (l *OrderPaymentLogic) walletPay(order model.Order, userID int64) (string, error) {
	//获取用户信息
	detail, err := l.svcCtx.UserRpc.UserDetail(l.ctx, &user.UserDetailReq{UserID: userID})
	if err != nil {
		return "", errors.Wrapf(xerr.NewErrMsg("PayRpc walletRpc USE UserRpc UserDetail ERROR"),
			"PayRpc walletRpc USE UserRpc UserDetail ERROR:%+v,userID:%v", err, userID)
	}
	//判断用户ID是否正确
	if userID != detail.ID {
		return "", errors.Wrapf(xerr.NewErrCode(xerr.PAYMENT_USER_ERROR), "orderSn:%v,userID:%v", order.OrderSn, userID)
	}
	//判断钱包余额够不够
	if detail.Money < order.TotalPrice {
		return "", errors.Wrapf(xerr.NewErrCode(xerr.PAYMENT_MONEY_NOT_ENOUGH_ERROR),
			"user Money:%v,order TotalPrice:%v", detail.Money, order.TotalPrice)
	}
	//够就生成支付流水
	paymentSn := sn.GenerateSn(sn.PAYMENT_PREFIX)
	//开启事务
	err = l.svcCtx.PayDB.Transaction(func(tx *gorm.DB) error {

		//生成kq消息
		kqMsg, err := json.Marshal(kqOrder.PaymentUpdateOrderState{
			OrderSn:    order.OrderSn,
			UserID:     userID,
			OrderState: globalKey.OrderPayed,
		})
		if err != nil {
			return errors.Wrapf(xerr.NewErrMsg("PaymentRPC USE MqKqOrder Generate message ERROR"),
				"PaymentRPC USE MqKqOrder Generate message ERROR:%+v", err)
		}

		//金额足够，更新用户余额信息
		_, err = l.svcCtx.UserRpc.UpdateUserMoney(l.ctx, &user.UpdateUserMoneyReq{
			UserID: userID,
			Money:  detail.Money - order.TotalPrice,
		})
		if err != nil {
			return errors.Wrapf(xerr.NewErrMsg("PaymentRpc USE UserRpc.UpdateUserMoney ERROR"),
				"PaymentRpc USE UserRpc.UpdateUserMoney ERROR:%+v", err)
		}
		//创建支付信息
		err = tx.Create(&model2.Pay{
			Id:             l.svcCtx.SnowflakeNode.Generate().Int64(),
			UserId:         userID,
			PaySn:          paymentSn,
			OrderSn:        order.OrderSn,
			Paytotal:       order.TotalPrice,
			TradeStateDesc: order.Remark,
			PayTime:        time.Now(),
			CreateTime:     time.Now(),
			UpdateTime:     time.Now(),
		}).Error
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL CREATE payment ERROR:%+v", err)
		}

		//通知修改订单信息(kafka消息队列)
		err = l.svcCtx.MqKqOrder.Push(string(kqMsg))
		if err != nil {
			return errors.Wrapf(xerr.NewErrMsg("PayRPC USE MqKqOrder.Push ERROR"), "PayRPC USE MqKqOrder.Push ERROR:%+v", err)
		}

		return nil
	})
	if err != nil {
		return "", err
	}
	return paymentSn, nil
}
