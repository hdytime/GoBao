package logic

import (
	"context"

	"GoBao/server/order/rpc/internal/svc"
	"GoBao/server/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderListLogic {
	return &GetOrderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderListLogic) GetOrderList(in *pb.GetOrderListReq) (*pb.GetOrderListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetOrderListResp{}, nil
}
