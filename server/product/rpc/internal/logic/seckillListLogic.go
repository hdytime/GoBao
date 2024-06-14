package logic

import (
	"context"

	"GoBao/server/product/rpc/internal/svc"
	"GoBao/server/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SeckillListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSeckillListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeckillListLogic {
	return &SeckillListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// seckillProduct
func (l *SeckillListLogic) SeckillList(in *pb.SeckillListReq) (*pb.SeckillListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SeckillListResp{}, nil
}
