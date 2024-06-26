// Code generated by goctl. DO NOT EDIT.
// Source: order.proto

package order

import (
	"context"

	"GoBao/server/order/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateSeckillOrderReq  = pb.CreateSeckillOrderReq
	CreateSeckillOrderResp = pb.CreateSeckillOrderResp
	DeleteOrderReq         = pb.DeleteOrderReq
	DeleteOrderResp        = pb.DeleteOrderResp
	GetOrderDetailReq      = pb.GetOrderDetailReq
	GetOrderDetailResp     = pb.GetOrderDetailResp
	GetOrderListReq        = pb.GetOrderListReq
	GetOrderListResp       = pb.GetOrderListResp
	SmallOrder             = pb.SmallOrder
	UpdateOrderStatusReq   = pb.UpdateOrderStatusReq
	UpdateOrderStatusResp  = pb.UpdateOrderStatusResp

	Order interface {
		CreateSeckillOrder(ctx context.Context, in *CreateSeckillOrderReq, opts ...grpc.CallOption) (*CreateSeckillOrderResp, error)
		GetOrderList(ctx context.Context, in *GetOrderListReq, opts ...grpc.CallOption) (*GetOrderListResp, error)
		GetOrderDetail(ctx context.Context, in *GetOrderDetailReq, opts ...grpc.CallOption) (*GetOrderDetailResp, error)
		DeleteOrder(ctx context.Context, in *DeleteOrderReq, opts ...grpc.CallOption) (*DeleteOrderResp, error)
		UpdateOrderStatus(ctx context.Context, in *UpdateOrderStatusReq, opts ...grpc.CallOption) (*UpdateOrderStatusResp, error)
	}

	defaultOrder struct {
		cli zrpc.Client
	}
)

func NewOrder(cli zrpc.Client) Order {
	return &defaultOrder{
		cli: cli,
	}
}

func (m *defaultOrder) CreateSeckillOrder(ctx context.Context, in *CreateSeckillOrderReq, opts ...grpc.CallOption) (*CreateSeckillOrderResp, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.CreateSeckillOrder(ctx, in, opts...)
}

func (m *defaultOrder) GetOrderList(ctx context.Context, in *GetOrderListReq, opts ...grpc.CallOption) (*GetOrderListResp, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.GetOrderList(ctx, in, opts...)
}

func (m *defaultOrder) GetOrderDetail(ctx context.Context, in *GetOrderDetailReq, opts ...grpc.CallOption) (*GetOrderDetailResp, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.GetOrderDetail(ctx, in, opts...)
}

func (m *defaultOrder) DeleteOrder(ctx context.Context, in *DeleteOrderReq, opts ...grpc.CallOption) (*DeleteOrderResp, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.DeleteOrder(ctx, in, opts...)
}

func (m *defaultOrder) UpdateOrderStatus(ctx context.Context, in *UpdateOrderStatusReq, opts ...grpc.CallOption) (*UpdateOrderStatusResp, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.UpdateOrderStatus(ctx, in, opts...)
}
