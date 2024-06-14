package order

import (
	"GoBao/common/ctxData"
	"GoBao/server/order/rpc/order"
	"context"

	"GoBao/server/order/api/internal/svc"
	"GoBao/server/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOrderLogic {
	return &DeleteOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteOrderLogic) DeleteOrder(req *types.DeleteOrderReq) error {
	// todo: add your logic here and delete this line

	uid := ctxData.GetUserIDFromCtx(l.ctx)
	_, err := l.svcCtx.OrderRpc.DeleteOrder(l.ctx, &order.DeleteOrderReq{UserID: uid, OrderSn: req.OrderSn})
	if err != nil {
		return err
	}
	return nil
}
