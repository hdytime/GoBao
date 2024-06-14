package cart

import (
	"GoBao/common/ctxData"
	"GoBao/server/cart/rpc/cart"
	"context"

	"GoBao/server/cart/api/internal/svc"
	"GoBao/server/cart/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductFromCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteProductFromCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductFromCartLogic {
	return &DeleteProductFromCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteProductFromCartLogic) DeleteProductFromCart(req *types.DeleteProductFromCartReq) error {
	// todo: add your logic here and delete this line

	UserID := ctxData.GetUserIDFromCtx(l.ctx)
	_, err := l.svcCtx.CartRpc.DeleteProductFromCart(l.ctx, &cart.DeleteProductFromCartRequest{
		CartID: req.CartID,
		UserID: UserID,
	})
	if err != nil {
		return err
	}
	return nil
}
