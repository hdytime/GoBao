// Code generated by goctl. DO NOT EDIT.
// Source: pay.proto

package pay

import (
	"context"

	"GoBao/server/pay/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetPaymentDetailReq  = pb.GetPaymentDetailReq
	GetPaymentDetailResp = pb.GetPaymentDetailResp
	OrderPaymentReq      = pb.OrderPaymentReq
	OrderPaymentResp     = pb.OrderPaymentResp

	Pay interface {
		OrderPayment(ctx context.Context, in *OrderPaymentReq, opts ...grpc.CallOption) (*OrderPaymentResp, error)
		GetPaymentDetail(ctx context.Context, in *GetPaymentDetailReq, opts ...grpc.CallOption) (*GetPaymentDetailResp, error)
	}

	defaultPay struct {
		cli zrpc.Client
	}
)

func NewPay(cli zrpc.Client) Pay {
	return &defaultPay{
		cli: cli,
	}
}

func (m *defaultPay) OrderPayment(ctx context.Context, in *OrderPaymentReq, opts ...grpc.CallOption) (*OrderPaymentResp, error) {
	client := pb.NewPayClient(m.cli.Conn())
	return client.OrderPayment(ctx, in, opts...)
}

func (m *defaultPay) GetPaymentDetail(ctx context.Context, in *GetPaymentDetailReq, opts ...grpc.CallOption) (*GetPaymentDetailResp, error) {
	client := pb.NewPayClient(m.cli.Conn())
	return client.GetPaymentDetail(ctx, in, opts...)
}
