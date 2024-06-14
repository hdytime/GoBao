package logic

import (
	"GoBao/common/xerr"
	"GoBao/server/user/model"
	"GoBao/server/user/rpc/internal/svc"
	"GoBao/server/user/rpc/pb"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserDetailLogic {
	return &UpdateUserDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserDetailLogic) UpdateUserDetail(in *pb.UpdateUserDetailReq) (*pb.UpdateUserDetailResp, error) {
	// todo: add your logic here and delete this line
	var u model.User
	//检查用户是否存在！
	err := l.svcCtx.DB.Where("id=?", in.UserID).Take(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_NOT_EXISTS_ERROR), "UserID:%+v", in.UserID)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE user ERROR")
	}
	//用户存在
	if in.Sex == 1 || in.Sex == 2 {
		u.Sex = in.Sex
	}
	u.Email = in.Email
	u.Sign = in.Sign
	//更新
	err = l.svcCtx.DB.Updates(&u).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_UPDATE_ZERO_ERROR), "MYSQL UPDATE user detail ERROR:%+v,userID:%+v", err, in.UserID)
	}

	return &pb.UpdateUserDetailResp{}, nil
}
