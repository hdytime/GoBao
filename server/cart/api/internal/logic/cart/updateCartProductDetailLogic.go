package cart

import (
	"GoBao/common/ctxData"
	"GoBao/server/cart/rpc/pb"
	"context"

	"GoBao/server/cart/api/internal/svc"
	"GoBao/server/cart/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCartProductDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCartProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCartProductDetailLogic {
	return &UpdateCartProductDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCartProductDetailLogic) UpdateCartProductDetail(req *types.UpdateCartProductDetailReq) error {
	// todo: add your logic here and delete this line

	UserID := ctxData.GetUserIDFromCtx(l.ctx)
	_, err := l.svcCtx.CartRpc.UpdateCartProductDetail(l.ctx, &pb.UpdateCartProductDetailRequest{
		CartID: req.CartID,
		Count:  req.Count,
		UserID: UserID,
	})
	if err != nil {
		return err
	}
	return nil
}
