package svc

import (
	"GoBao/server/order/rpc/internal/config"
	"github.com/hibiken/asynq"
)

func newAsynqClient(c config.Config) *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{
		Addr:     c.RedisConf.Addr,
		DB:       c.RedisConf.DB,
		PoolSize: c.RedisConf.PoolSize,
	})
}
