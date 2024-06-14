package logic

import (
	tool "GoBao/common/tool/md5"
	"GoBao/common/xerr"
	"GoBao/server/user/model"
	"GoBao/server/user/rpc/user"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"GoBao/server/user/rpc/internal/svc"
	"GoBao/server/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	// todo: add your logic here and delete this line
	var u = new(model.User)

	err := l.svcCtx.DB.Where("username=?", in.Username).Take(&u).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL TAKE user ERROR:%+v", err)
	} else if err == nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_EXISTS_ERROR), "用户已经存在")
	}
	//密码加密
	newPassword := tool.Md5ToString(in.Password)
	var newUser = model.User{
		ID:          l.svcCtx.SnowflakeNode.Generate().Int64(),
		Username:    in.Username,
		Password:    newPassword,
		Money:       10000,
		Sex:         0,
		PhoneNumber: 0,
		Email:       "",
		Sign:        "",
	}
	if err := l.svcCtx.DB.Create(&newUser).Error; err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL CREATE user ERROR:%+v", err)
	}
	GenerateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	token, err := GenerateTokenLogic.GenerateToken(&user.GenerateTokenReq{
		UserID: newUser.ID,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "Register Generate TOKEN ERROR")
	}
	return &pb.RegisterResp{
		AccessToken:  token.AccessToken,
		AccessExpire: token.AccessExpire,
		RefreshAfter: token.RefreshAfter,
	}, nil
}
