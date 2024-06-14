package logic

import (
	"GoBao/common/globalKey"
	"GoBao/common/xerr"
	"GoBao/server/pay/model"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"

	"GoBao/server/pay/rpc/internal/svc"
	"GoBao/server/pay/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPaymentDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPaymentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPaymentDetailLogic {
	return &GetPaymentDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPaymentDetailLogic) GetPaymentDetail(in *pb.GetPaymentDetailReq) (*pb.GetPaymentDetailResp, error) {
	// todo: add your logic here and delete this line

	var payment model.Pay
	err := l.svcCtx.PayDB.Where("order_sn = ? and del_state = ?", in.OrderSn, globalKey.DelStateNo).Take(&payment).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.PAYMENT_NOT_EXISTS_ERROR), "orderSn: %v", in.OrderSn)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE payment ERROR: %+v, orderSn: %v", err, in.OrderSn)
	}

	return &pb.GetPaymentDetailResp{
		ID:             payment.Id,
		PaySn:          payment.PaySn,
		OrderSn:        payment.OrderSn,
		UserID:         payment.UserId,
		TradeState:     strconv.FormatInt(payment.TradeState, 10),
		PayTotal:       payment.Paytotal,
		TransactionID:  strconv.FormatInt(payment.TransactionID, 10),
		TradeStateDesc: payment.TradeStateDesc,
		PayTime:        payment.CreateTime.Format("2006-01-02 15:04:05"),
	}, nil
}
