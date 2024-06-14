package commonProduct

import (
	"GoBao/common/tool"
	"GoBao/server/product/rpc/pb"
	"context"

	"GoBao/server/product/api/internal/svc"
	"GoBao/server/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchProductLogic {
	return &SearchProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchProductLogic) SearchProduct(req *types.SearchProductReq) (resp *types.SearchProductResp, err error) {
	// todo: add your logic here and delete this line

	page, size, sort := tool.CheckBasePageAndSort(req.Page, req.Size, req.Sort)
	productResp, err := l.svcCtx.ProductRpc.SearchProduct(l.ctx, &pb.SearchProductReq{
		Keyword: req.Keyword,
		Sort:    sort,
		OnSale:  req.OnSale,
		Page:    page,
		Size:    size,
	})
	if err != nil {
		return nil, err
	}

	var res []types.SmallProduct
	for _, smallproduct := range productResp.SmallProducts {
		var sp = types.SmallProduct{
			ID:            smallproduct.ID,
			Name:          smallproduct.Name,
			Price:         smallproduct.Price,
			DiscountPrice: smallproduct.DiscountPrice,
		}
		res = append(res, sp)
	}
	return &types.SearchProductResp{Products: res}, nil
}
