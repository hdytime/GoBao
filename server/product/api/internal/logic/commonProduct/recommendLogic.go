package commonProduct

import (
	"GoBao/server/product/api/internal/svc"
	"GoBao/server/product/api/internal/types"
	"GoBao/server/product/rpc/pb"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecommendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecommendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecommendLogic {
	return &RecommendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecommendLogic) Recommend() (resp *types.RecommendProductResp, err error) {
	// todo: add your logic here and delete this line

	recommendResp, err := l.svcCtx.ProductRpc.Recommend(l.ctx, &pb.RecommendReq{})
	if err != nil {
		return nil, err
	}

	var res []types.SmallProduct
	for _, p := range recommendResp.SmallProducts {
		var sp = types.SmallProduct{
			ID:            p.ID,
			Name:          p.Name,
			Price:         p.Price,
			DiscountPrice: p.DiscountPrice,
		}
		res = append(res, sp)
	}
	return &types.RecommendProductResp{Products: res}, err
}
