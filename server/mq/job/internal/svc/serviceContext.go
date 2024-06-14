package svc

import (
	"GoBao/common/goredis"
	"GoBao/server/mq/job/internal/config"
	"GoBao/server/order/rpc/order"
	"GoBao/server/product/rpc/productrpc"
	"GoBao/server/user/rpc/user"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	AsynqServer *asynq.Server

	RedisDB *redis.Client

	OrderRpc   order.Order
	ProductRpc productrpc.ProductRpc
	UserRpc    user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		AsynqServer: newAsynqServer(c),

		RedisDB: goredis.Rdb,

		OrderRpc:   order.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
		ProductRpc: productrpc.NewProductRpc(zrpc.MustNewClient(c.ProductRpc)),
		UserRpc:    user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
