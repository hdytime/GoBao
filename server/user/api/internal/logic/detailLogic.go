package logic

import (
	"GoBao/common/ctxData"
	"GoBao/server/user/api/internal/svc"
	"GoBao/server/user/api/internal/types"
	"GoBao/server/user/rpc/pb"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail() (resp *types.DetailResp, err error) {
	// todo: add your logic here and delete this line

	userID := ctxData.GetUserIDFromCtx(l.ctx)
	userDetailResp, err := l.svcCtx.UserRpc.UserDetail(l.ctx, &pb.UserDetailReq{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}

	resp = new(types.DetailResp)
	resp.UserInfo.Username = userDetailResp.Username
	resp.UserInfo.Password = userDetailResp.Password
	resp.UserInfo.Money = userDetailResp.Money
	resp.UserInfo.Sex = userDetailResp.Sex
	resp.UserInfo.PhoneNumber = userDetailResp.PhoneNumber
	resp.UserInfo.Email = userDetailResp.Email
	resp.UserInfo.Sign = userDetailResp.Sign

	return
}
