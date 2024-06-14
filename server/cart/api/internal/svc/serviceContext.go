package svc

import (
	"GoBao/server/cart/api/internal/config"
	"GoBao/server/cart/rpc/cart"
	"GoBao/server/product/rpc/productrpc"
	"GoBao/server/user/rpc/user"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	CartRpc    cart.Cart
	ProductRpc productrpc.ProductRpc
	UserRpc    user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		CartRpc:    cart.NewCart(zrpc.MustNewClient(c.CartRpcConf)),
		ProductRpc: productrpc.NewProductRpc(zrpc.MustNewClient(c.ProductRpcConf)),
		UserRpc:    user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
