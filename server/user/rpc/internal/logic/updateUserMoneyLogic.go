package logic

import (
	"GoBao/common/xerr"
	"GoBao/server/user/model"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"GoBao/server/user/rpc/internal/svc"
	"GoBao/server/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserMoneyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserMoneyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserMoneyLogic {
	return &UpdateUserMoneyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserMoneyLogic) UpdateUserMoney(in *pb.UpdateUserMoneyReq) (*pb.UpdateUserMoneyResp, error) {
	// todo: add your logic here and delete this line
	var u model.User
	err := l.svcCtx.DB.Where("id=?", in.UserID).Take(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_NOT_EXISTS_ERROR), "userID:%v", in.UserID)
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE user ERROR:%+v", err)
	}
	u.Money = in.Money
	err = l.svcCtx.DB.Where("id=?", in.UserID).Updates(&u).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL UPDATE user ERROR:%+v", err)
	}
	return &pb.UpdateUserMoneyResp{}, nil
}
