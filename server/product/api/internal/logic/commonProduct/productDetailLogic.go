package commonProduct

import (
	"GoBao/server/product/rpc/pb"
	"context"

	"GoBao/server/product/api/internal/svc"
	"GoBao/server/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductDetailLogic {
	return &ProductDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductDetailLogic) ProductDetail(req *types.ProductDetailReq) (resp *types.ProductDetailResp, err error) {
	// todo: add your logic here and delete this line

	productDetailResp, err := l.svcCtx.ProductRpc.ProductDetail(l.ctx, &pb.ProductDetailReq{
		ProductID: req.ProductID,
	})
	if err != nil {
		return nil, err
	}
	var pro = types.Product{
		ID:         productDetailResp.Product.ID,
		Name:       productDetailResp.Product.Name,
		Price:      productDetailResp.Product.Price,
		Stock:      productDetailResp.Product.Stock,
		Status:     productDetailResp.Product.Status,
		CreateTime: productDetailResp.Product.CreateTime,
		UpdateTime: productDetailResp.Product.UpdateTime,
	}

	return &types.ProductDetailResp{Product: pro}, nil
}
