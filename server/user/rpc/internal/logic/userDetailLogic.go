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

type UserDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserDetailLogic) UserDetail(in *pb.UserDetailReq) (*pb.UserDetailResp, error) {
	// todo: add your logic here and delete this line
	var u model.User
	err := l.svcCtx.DB.Where("id=?", in.UserID).Take(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_NOT_EXISTS_ERROR), "userID:%+v", in.UserID)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE user ERROR")
	}
	return &pb.UserDetailResp{
		ID:          u.ID,
		Username:    u.Username,
		Password:    u.Password,
		Money:       u.Money,
		Sex:         u.Sex,
		PhoneNumber: u.PhoneNumber,
		Email:       u.Email,
		Sign:        u.Sign,
	}, nil
}
