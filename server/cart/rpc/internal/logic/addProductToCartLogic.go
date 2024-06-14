package logic

import (
	"GoBao/common/xerr"
	"GoBao/server/cart/model"
	"GoBao/server/product/rpc/productrpc"
	"context"
	"github.com/pkg/errors"
	"time"

	"GoBao/server/cart/rpc/internal/svc"
	"GoBao/server/cart/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProductToCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProductToCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductToCartLogic {
	return &AddProductToCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddProductToCartLogic) AddProductToCart(in *pb.AddProductToCartRequest) (*pb.AddProductToCartResponse, error) {
	// todo: add your logic here and delete this line
	//查询是否有该商品
	detailResp, err := l.svcCtx.ProductRpc.ProductDetail(l.ctx, &productrpc.ProductDetailReq{ProductID: in.ProductID})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("CartRpc USER ProductRpc ERROR"), "CartRpc USE ProductRpc ERROR:%+v", err)
	}
	if err == nil && detailResp.Product == nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.PRODUCT_NOT_EXISTS_ERROR), "PRODUCT NOT EXIST,productID:%+v", in.ProductID)
	}

	l.svcCtx.CartDB.Create(&model.CartProduct{
		Id:         l.svcCtx.SnowflakeNode.Generate().Int64(),
		UserId:     in.UserID,
		ProductId:  in.ProductID,
		Price:      detailResp.Product.Price,
		Quantity:   in.Count,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})

	return &pb.AddProductToCartResponse{}, nil
}
