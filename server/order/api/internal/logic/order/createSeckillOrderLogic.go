package order

import (
	"GoBao/common/ctxData"
	"GoBao/server/order/rpc/order"
	"context"

	"GoBao/server/order/api/internal/svc"
	"GoBao/server/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSeckillOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSeckillOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSeckillOrderLogic {
	return &CreateSeckillOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSeckillOrderLogic) CreateSeckillOrder(req *types.CreateSeckillOrderReq) (resp *types.CreateSeckillOrderResp, err error) {
	// todo: add your logic here and delete this line

	uid := ctxData.GetUserIDFromCtx(l.ctx)
	CreateSeckillOrderResp, err := l.svcCtx.OrderRpc.CreateSeckillOrder(l.ctx, &order.CreateSeckillOrderReq{
		UserID:       uid,
		ProductID:    req.ProductID,
		ProductCount: req.ProductCount,
		Remark:       req.Remark,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateSeckillOrderResp{OrderSn: CreateSeckillOrderResp.OrderSn}, nil
}
