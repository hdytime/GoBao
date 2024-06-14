package seckillProduct

import (
	"GoBao/server/product/rpc/pb"
	"context"

	"GoBao/server/product/api/internal/svc"
	"GoBao/server/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SeckillDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSeckillDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeckillDetailLogic {
	return &SeckillDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SeckillDetailLogic) SeckillDetail(req *types.GetSeckillDetailReq) (resp *types.GetSeckillDetailResp, err error) {
	// todo: add your logic here and delete this line

	seckillDetailResp, err := l.svcCtx.ProductRpc.SeckillDetail(l.ctx, &pb.SeckillDetailReq{SeckillProductID: req.SeckillID})
	if err != nil {
		return nil, err
	}

	var res = types.SeckillProduct{
		Product: types.Product{
			ID:         seckillDetailResp.SeckillProduct.ID,
			Name:       seckillDetailResp.SeckillProduct.Name,
			Price:      seckillDetailResp.SeckillProduct.Price,
			Stock:      seckillDetailResp.SeckillProduct.Stock,
			Status:     seckillDetailResp.SeckillProduct.Status,
			CreateTime: seckillDetailResp.SeckillProduct.CreateTime,
			UpdateTime: seckillDetailResp.SeckillProduct.UpdateTime,
		},
		SeckillPrice: seckillDetailResp.SeckillProduct.SeckillPrice,
		StockCount:   seckillDetailResp.SeckillProduct.SeckillCount,
		StartTime:    seckillDetailResp.SeckillProduct.StartTime,
		Time:         seckillDetailResp.SeckillProduct.Time,
	}
	return &types.GetSeckillDetailResp{SeckillProduct: res}, nil
}
