// Code generated by goctl. DO NOT EDIT.
// Source: cart.proto

package server

import (
	"context"

	"GoBao/server/cart/rpc/internal/logic"
	"GoBao/server/cart/rpc/internal/svc"
	"GoBao/server/cart/rpc/pb"
)

type CartServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedCartServer
}

func NewCartServer(svcCtx *svc.ServiceContext) *CartServer {
	return &CartServer{
		svcCtx: svcCtx,
	}
}

func (s *CartServer) AddProductToCart(ctx context.Context, in *pb.AddProductToCartRequest) (*pb.AddProductToCartResponse, error) {
	l := logic.NewAddProductToCartLogic(ctx, s.svcCtx)
	return l.AddProductToCart(in)
}

func (s *CartServer) DeleteProductFromCart(ctx context.Context, in *pb.DeleteProductFromCartRequest) (*pb.DeleteProductFromCartResponse, error) {
	l := logic.NewDeleteProductFromCartLogic(ctx, s.svcCtx)
	return l.DeleteProductFromCart(in)
}

func (s *CartServer) UpdateCartProductDetail(ctx context.Context, in *pb.UpdateCartProductDetailRequest) (*pb.UpdateCartProductDetailResponse, error) {
	l := logic.NewUpdateCartProductDetailLogic(ctx, s.svcCtx)
	return l.UpdateCartProductDetail(in)
}

func (s *CartServer) GetCartList(ctx context.Context, in *pb.GetCartListRequest) (*pb.GetCartListResponse, error) {
	l := logic.NewGetCartListLogic(ctx, s.svcCtx)
	return l.GetCartList(in)
}
