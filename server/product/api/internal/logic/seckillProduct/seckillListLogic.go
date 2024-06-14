package seckillProduct

import (
	"context"

	"GoBao/server/product/api/internal/svc"
	"GoBao/server/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SeckillListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSeckillListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeckillListLogic {
	return &SeckillListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SeckillListLogic) SeckillList(req *types.GetSeckillListReq) (resp *types.GetSeckillListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
