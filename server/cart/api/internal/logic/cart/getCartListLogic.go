package cart

import (
	"GoBao/common/ctxData"
	"GoBao/server/cart/api/internal/svc"
	"GoBao/server/cart/api/internal/types"
	"GoBao/server/cart/rpc/cart"
	"context"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCartListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCartListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCartListLogic {
	return &GetCartListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCartListLogic) GetCartList() (resp *types.GetCartListResp, err error) {
	// todo: add your logic here and delete this line

	UserID := ctxData.GetUserIDFromCtx(l.ctx)

	list, err := l.svcCtx.CartRpc.GetCartList(l.ctx, &cart.GetCartListRequest{
		UserID: UserID,
	})
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&resp, &list)
	if err != nil {
		return nil, err
	}
	return
}
