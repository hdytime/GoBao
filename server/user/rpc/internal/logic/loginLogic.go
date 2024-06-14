package logic

import (
	tool "GoBao/common/tool/md5"
	"GoBao/common/xerr"
	"GoBao/server/user/model"
	"GoBao/server/user/rpc/internal/svc"
	"GoBao/server/user/rpc/pb"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	// todo: add your logic here and delete this line

	var u model.User
	err := l.svcCtx.DB.Where("username=?", in.Username).Take(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_NOT_EXISTS_ERROR), "username:%+v")
		}
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE user ERROR:%+v", err)
	}
	if tool.Md5ToString(in.Password) != u.Password {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_PASSWORD_ERROR), "password:%+v", in.Password)
	}

	//登录成功，生成token
	generateTokenLogin := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	generateTokenResp, err := generateTokenLogin.GenerateToken(&pb.GenerateTokenReq{
		UserID: u.ID,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.TOKEN_GENERATE_ERROR), "Login Generate TOKEN ERROR")
	}

	return &pb.LoginResp{
		AccessToken:  generateTokenResp.AccessToken,
		AccessExpire: generateTokenResp.AccessExpire,
		RefreshAfter: generateTokenResp.RefreshAfter,
	}, nil
}
