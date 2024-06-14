package svc

import (
	"GoBao/server/order/api/internal/config"
	"GoBao/server/order/rpc/order"
	"GoBao/server/product/rpc/productrpc"
	"GoBao/server/user/rpc/user"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	UserRpc    user.User
	ProductRpc productrpc.ProductRpc
	OrderRpc   order.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UserRpc:    user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		ProductRpc: productrpc.NewProductRpc(zrpc.MustNewClient(c.ProductRpcConf)),
		OrderRpc:   order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
	}
}
