package svc

import (
	"GoBao/common/consts"
	"GoBao/common/goredis"
	gorm2 "GoBao/common/gorm"
	"GoBao/server/product/rpc/internal/config"
	"GoBao/server/product/rpc/productrpc"
	"github.com/bwmarrin/snowflake"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/zrpc"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config        config.Config
	ProductDB     *gorm.DB
	RedisDB       *redis.Client
	SnowflakeNode *snowflake.Node
	SingleGroup   singleflight.Group
	ProductRpc    productrpc.ProductRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	snowflakeNode, _ := snowflake.NewNode(consts.ProductSnowflakeNodeID)
	return &ServiceContext{
		Config:        c,
		ProductDB:     gorm2.ProductDB,
		RedisDB:       goredis.Rdb,
		SnowflakeNode: snowflakeNode,
		ProductRpc:    productrpc.NewProductRpc(zrpc.MustNewClient(c.ProductRpc)),
	}
}
