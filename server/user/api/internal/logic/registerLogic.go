package logic

import (
	"GoBao/server/user/api/internal/svc"
	"GoBao/server/user/api/internal/types"
	"GoBao/server/user/rpc/pb"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line

	registerResp, err := l.svcCtx.UserRpc.Register(l.ctx, &pb.RegisterReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	resp = new(types.RegisterResp)
	resp.JwtAccess.AccessToken = registerResp.AccessToken
	resp.JwtAccess.AccessExpire = registerResp.AccessExpire
	resp.JwtAccess.RefreshAfter = registerResp.RefreshAfter

	return resp, nil
}
