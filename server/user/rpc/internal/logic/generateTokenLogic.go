package logic

import (
	"GoBao/common/utils"
	"GoBao/common/xerr"
	"context"
	"github.com/pkg/errors"
	"time"

	"GoBao/server/user/rpc/internal/svc"
	"GoBao/server/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateTokenLogic) GenerateToken(in *pb.GenerateTokenReq) (*pb.GenerateTokenResp, error) {
	// todo: add your logic here and delete this line

	now := time.Now().Unix()
	accessSecret := l.svcCtx.Config.JwtAuth.AccessSecret
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	token, err := utils.GenerateJwtToken(accessSecret, now, accessExpire, in.UserID)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.TOKEN_GENERATE_ERROR), "getjwttoken ERROR:%+v,userid:%v", err, in.UserID)
	}
	return &pb.GenerateTokenResp{
		AccessToken:  token,
		AccessExpire: accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}
