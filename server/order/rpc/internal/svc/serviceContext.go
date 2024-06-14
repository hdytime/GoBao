package svc

import (
	"GoBao/common/consts"
	"GoBao/common/goredis"
	gorm2 "GoBao/common/gorm"
	"GoBao/server/order/rpc/internal/config"
	"GoBao/server/pay/rpc/pay"
	"GoBao/server/product/rpc/productrpc"
	"GoBao/server/user/rpc/user"
	"github.com/bwmarrin/snowflake"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config         config.Config
	UserRpc        user.User
	ProductRpc     productrpc.ProductRpc
	PayRpc         pay.Pay
	OrderDB        *gorm.DB
	RedisDB        *redis.Client
	SnowflakeNode  *snowflake.Node
	AsynqClient    *asynq.Client
	KqPusherClient *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	node, _ := snowflake.NewNode(consts.OrderSnowflakeNodeID)
	return &ServiceContext{
		Config:         c,
		UserRpc:        user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		ProductRpc:     productrpc.NewProductRpc(zrpc.MustNewClient(c.ProductRpcConf)),
		PayRpc:         pay.NewPay(zrpc.MustNewClient(c.PayRpcConf)),
		OrderDB:        gorm2.OrderDB,
		RedisDB:        goredis.Rdb,
		SnowflakeNode:  node,
		AsynqClient:    newAsynqClient(c),
		KqPusherClient: kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
	}
}
