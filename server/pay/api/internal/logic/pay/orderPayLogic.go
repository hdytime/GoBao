package pay

import (
	"GoBao/common/ctxData"
	"GoBao/server/pay/rpc/pb"
	"context"

	"GoBao/server/pay/api/internal/svc"
	"GoBao/server/pay/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderPayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderPayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderPayLogic {
	return &OrderPayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderPayLogic) OrderPay(req *types.OrderPayReq) (resp *types.OrderPayResp, err error) {
	// todo: add your logic here and delete this line
	uid := ctxData.GetUserIDFromCtx(l.ctx)
	payment, err := l.svcCtx.PayRpc.OrderPayment(l.ctx, &pb.OrderPaymentReq{OrderSn: req.OrderSn, UserID: uid})
	if err != nil {
		return nil, err
	}

	return &types.OrderPayResp{
		PayTotalPrice: payment.PayTotalPrice,
		Paysn:         payment.PaySn,
	}, nil
}
