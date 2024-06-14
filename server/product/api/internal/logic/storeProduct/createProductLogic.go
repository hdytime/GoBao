package storeProduct

import (
	"GoBao/common/ctxData"
	"GoBao/server/product/rpc/pb"
	"context"

	"GoBao/server/product/api/internal/svc"
	"GoBao/server/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductLogic {
	return &CreateProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateProductLogic) CreateProduct(req *types.CreateProductReq) error {
	// todo: add your logic here and delete this line

	//获取userid
	userID := ctxData.GetUserIDFromCtx(l.ctx)

	_, err := l.svcCtx.ProductRpc.CreateProduct(l.ctx, &pb.CreateProductReq{
		Name:   req.Name,
		Price:  req.Price,
		Stock:  req.Stock,
		UserID: userID,
	})
	if err != nil {
		return err
	}
	return nil
}
