package storeProduct

import (
	"GoBao/server/product/rpc/pb"
	"context"

	"GoBao/server/product/api/internal/svc"
	"GoBao/server/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductLogic {
	return &DeleteProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteProductLogic) DeleteProduct(req *types.DeleteProductReq) error {
	// todo: add your logic here and delete this line

	_, err := l.svcCtx.ProductRpc.DeleteProduct(l.ctx, &pb.DeleteProductReq{ProductIDs: req.ProductIDs})
	if err != nil {
		return err
	}

	return nil
}
