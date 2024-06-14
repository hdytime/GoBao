package svc

import (
	"GoBao/common/consts"
	"GoBao/common/goredis"
	gorm2 "GoBao/common/gorm"
	"GoBao/server/order/rpc/order"
	"GoBao/server/pay/rpc/internal/config"
	"GoBao/server/product/rpc/productrpc"
	"GoBao/server/user/rpc/user"
	"github.com/bwmarrin/snowflake"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config        config.Config
	PayDB         *gorm.DB
	RedisDB       *redis.Client
	SnowflakeNode *snowflake.Node
	UserRpc       user.User
	ProductRpc    productrpc.ProductRpc
	OrderRpc      order.Order
	MqKqOrder     *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	snowflakeNode, _ := snowflake.NewNode(consts.PaySnowflakeNodeID)
	return &ServiceContext{
		Config:        c,
		PayDB:         gorm2.PaymentDB,
		RedisDB:       goredis.Rdb,
		SnowflakeNode: snowflakeNode,
		UserRpc:       user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		ProductRpc:    productrpc.NewProductRpc(zrpc.MustNewClient(c.ProductRpcConf)),
		OrderRpc:      order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
		MqKqOrder:     kq.NewPusher(c.KqPaymentUpdateOrderStateConf.Brokers, c.KqPaymentUpdateOrderStateConf.Topic),
	}
}
