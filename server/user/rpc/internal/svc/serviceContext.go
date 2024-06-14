package svc

import (
	"GoBao/common/consts"
	"GoBao/common/goredis"
	gorm2 "GoBao/common/gorm"
	"GoBao/server/user/rpc/internal/config"
	"github.com/bwmarrin/snowflake"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config        config.Config
	DB            *gorm.DB
	SnowflakeNode *snowflake.Node
	RedisDB       *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	snowflakeNode, _ := snowflake.NewNode(consts.UserSnowflakeNodeID)
	return &ServiceContext{
		Config:        c,
		DB:            gorm2.UserDB,
		SnowflakeNode: snowflakeNode,
		RedisDB:       goredis.Rdb,
	}
}
