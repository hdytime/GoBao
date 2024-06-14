package svc

import (
	"GoBao/server/product/api/internal/config"
	"GoBao/server/product/rpc/productrpc"
	"GoBao/server/user/rpc/user"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	ProductRpc productrpc.ProductRpc
	UserRpc    user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ProductRpc: productrpc.NewProductRpc(zrpc.MustNewClient(c.ProductRpcConf)),
		UserRpc:    user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
