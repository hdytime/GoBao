package svc

import (
	"GoBao/common/goredis"
	gorm2 "GoBao/common/gorm"
	"GoBao/server/user/api/internal/config"
	"GoBao/server/user/rpc/user"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.User
	DB      *gorm.DB
	RedisDB *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		DB:      gorm2.UserDB,
		RedisDB: goredis.Rdb,
	}
}
