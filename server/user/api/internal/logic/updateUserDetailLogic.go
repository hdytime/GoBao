package logic

import (
	"GoBao/common/ctxData"
	"GoBao/server/user/api/internal/svc"
	"GoBao/server/user/api/internal/types"
	"GoBao/server/user/rpc/pb"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserDetailLogic {
	return &UpdateUserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserDetailLogic) UpdateUserDetail(req *types.UpdateUserInfoReq) (err error) {
	// todo: add your logic here and delete this line

	userID := ctxData.GetUserIDFromCtx(l.ctx)

	_, err = l.svcCtx.UserRpc.UpdateUserDetail(l.ctx, &pb.UpdateUserDetailReq{
		UserID: userID,
		Sex:    req.Sex,
		Email:  req.Email,
		Sign:   req.Sign,
	})
	if err != nil {
		return err
	}
	return nil
}
