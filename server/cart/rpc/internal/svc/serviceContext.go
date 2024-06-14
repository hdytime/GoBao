package svc

import (
	"GoBao/common/consts"
	"GoBao/common/goredis"
	gorm2 "GoBao/common/gorm"
	"GoBao/server/cart/rpc/internal/config"
	"GoBao/server/product/rpc/productrpc"
	"GoBao/server/user/rpc/user"
	"github.com/bwmarrin/snowflake"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config        config.Config
	CartDB        *gorm.DB
	RedisDB       *redis.Client
	SnowflakeNode *snowflake.Node
	ProductRpc    productrpc.ProductRpc
	UserRpc       user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	node, _ := snowflake.NewNode(consts.CartSnowflakeNodeID)
	return &ServiceContext{
		Config:        c,
		CartDB:        gorm2.CartDB,
		RedisDB:       goredis.Rdb,
		SnowflakeNode: node,
		ProductRpc:    productrpc.NewProductRpc(zrpc.MustNewClient(c.ProductRpc)),
		UserRpc:       user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
