// Code generated by goctl. DO NOT EDIT.
// Source: product.proto

package productrpc

import (
	"context"

	"GoBao/server/product/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateProductReq       = pb.CreateProductReq
	CreateProductResp      = pb.CreateProductResp
	DeductSeckillStockReq  = pb.DeductSeckillStockReq
	DeductSeckillStockResp = pb.DeductSeckillStockResp
	DeleteProductReq       = pb.DeleteProductReq
	DeleteProductResp      = pb.DeleteProductResp
	PreloadCacheReq        = pb.PreloadCacheReq
	PreloadCacheResp       = pb.PreloadCacheResp
	Product                = pb.Product
	ProductDetailReq       = pb.ProductDetailReq
	ProductDetailResp      = pb.ProductDetailResp
	RecommendReq           = pb.RecommendReq
	RecommendResp          = pb.RecommendResp
	SearchProductReq       = pb.SearchProductReq
	SearchProductResp      = pb.SearchProductResp
	SeckillDetailReq       = pb.SeckillDetailReq
	SeckillDetailResp      = pb.SeckillDetailResp
	SeckillListReq         = pb.SeckillListReq
	SeckillListResp        = pb.SeckillListResp
	SeckillProduct         = pb.SeckillProduct
	SmallProduct           = pb.SmallProduct

	ProductRpc interface {
		// commonProduct
		Recommend(ctx context.Context, in *RecommendReq, opts ...grpc.CallOption) (*RecommendResp, error)
		SearchProduct(ctx context.Context, in *SearchProductReq, opts ...grpc.CallOption) (*SearchProductResp, error)
		ProductDetail(ctx context.Context, in *ProductDetailReq, opts ...grpc.CallOption) (*ProductDetailResp, error)
		// seckillProduct
		SeckillList(ctx context.Context, in *SeckillListReq, opts ...grpc.CallOption) (*SeckillListResp, error)
		SeckillDetail(ctx context.Context, in *SeckillDetailReq, opts ...grpc.CallOption) (*SeckillDetailResp, error)
		PreloadCache(ctx context.Context, in *PreloadCacheReq, opts ...grpc.CallOption) (*PreloadCacheResp, error)
		DeductSeckillStock(ctx context.Context, in *DeductSeckillStockReq, opts ...grpc.CallOption) (*DeductSeckillStockResp, error)
		// storeProduct
		CreateProduct(ctx context.Context, in *CreateProductReq, opts ...grpc.CallOption) (*CreateProductResp, error)
		DeleteProduct(ctx context.Context, in *DeleteProductReq, opts ...grpc.CallOption) (*DeleteProductResp, error)
	}

	defaultProductRpc struct {
		cli zrpc.Client
	}
)

func NewProductRpc(cli zrpc.Client) ProductRpc {
	return &defaultProductRpc{
		cli: cli,
	}
}

// commonProduct
func (m *defaultProductRpc) Recommend(ctx context.Context, in *RecommendReq, opts ...grpc.CallOption) (*RecommendResp, error) {
	client := pb.NewProductRpcClient(m.cli.Conn())
	return client.Recommend(ctx, in, opts...)
}

func (m *defaultProductRpc) SearchProduct(ctx context.Context, in *SearchProductReq, opts ...grpc.CallOption) (*SearchProductResp, error) {
	client := pb.NewProductRpcClient(m.cli.Conn())
	return client.SearchProduct(ctx, in, opts...)
}

func (m *defaultProductRpc) ProductDetail(ctx context.Context, in *ProductDetailReq, opts ...grpc.CallOption) (*ProductDetailResp, error) {
	client := pb.NewProductRpcClient(m.cli.Conn())
	return client.ProductDetail(ctx, in, opts...)
}

// seckillProduct
func (m *defaultProductRpc) SeckillList(ctx context.Context, in *SeckillListReq, opts ...grpc.CallOption) (*SeckillListResp, error) {
	client := pb.NewProductRpcClient(m.cli.Conn())
	return client.SeckillList(ctx, in, opts...)
}

func (m *defaultProductRpc) SeckillDetail(ctx context.Context, in *SeckillDetailReq, opts ...grpc.CallOption) (*SeckillDetailResp, error) {
	client := pb.NewProductRpcClient(m.cli.Conn())
	return client.SeckillDetail(ctx, in, opts...)
}

func (m *defaultProductRpc) PreloadCache(ctx context.Context, in *PreloadCacheReq, opts ...grpc.CallOption) (*PreloadCacheResp, error) {
	client := pb.NewProductRpcClient(m.cli.Conn())
	return client.PreloadCache(ctx, in, opts...)
}

func (m *defaultProductRpc) DeductSeckillStock(ctx context.Context, in *DeductSeckillStockReq, opts ...grpc.CallOption) (*DeductSeckillStockResp, error) {
	client := pb.NewProductRpcClient(m.cli.Conn())
	return client.DeductSeckillStock(ctx, in, opts...)
}

// storeProduct
func (m *defaultProductRpc) CreateProduct(ctx context.Context, in *CreateProductReq, opts ...grpc.CallOption) (*CreateProductResp, error) {
	client := pb.NewProductRpcClient(m.cli.Conn())
	return client.CreateProduct(ctx, in, opts...)
}

func (m *defaultProductRpc) DeleteProduct(ctx context.Context, in *DeleteProductReq, opts ...grpc.CallOption) (*DeleteProductResp, error) {
	client := pb.NewProductRpcClient(m.cli.Conn())
	return client.DeleteProduct(ctx, in, opts...)
}