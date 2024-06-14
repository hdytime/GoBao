package cart

import (
	"GoBao/common/ctxData"
	"GoBao/server/cart/rpc/pb"
	"context"

	"GoBao/server/cart/api/internal/svc"
	"GoBao/server/cart/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProductToCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddProductToCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductToCartLogic {
	return &AddProductToCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddProductToCartLogic) AddProductToCart(req *types.AddProductToCartReq) error {
	// todo: add your logic here and delete this line

	// 获取用户ID
	userID := ctxData.GetUserIDFromCtx(l.ctx)

	_, err := l.svcCtx.CartRpc.AddProductToCart(l.ctx, &pb.AddProductToCartRequest{
		UserID:    userID,
		ProductID: req.ProductID,
		Count:     req.Count,
	})
	if err != nil {
		return err
	}
	return nil
}
