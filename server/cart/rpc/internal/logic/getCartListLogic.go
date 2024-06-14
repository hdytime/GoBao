package logic

import (
	"GoBao/common/xerr"
	"GoBao/server/cart/model"
	"GoBao/server/cart/rpc/internal/svc"
	"GoBao/server/cart/rpc/pb"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCartListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCartListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCartListLogic {
	return &GetCartListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCartListLogic) GetCartList(in *pb.GetCartListRequest) (*pb.GetCartListResponse, error) {
	// todo: add your logic here and delete this line

	var cartlist []model.CartProduct
	err := l.svcCtx.CartDB.Where("user_id=?", in.UserID).Find(&cartlist).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL FIND cart ERROR:%+v,UserID:%+v", err, in.UserID)
	}

	var res []*pb.CartProduct

	copier.Copy(&res, &cartlist)
	return &pb.GetCartListResponse{
		CartProducts: res,
	}, nil
}
