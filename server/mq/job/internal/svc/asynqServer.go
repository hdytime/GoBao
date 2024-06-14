package svc

import (
	"GoBao/server/mq/job/internal/config"
	"context"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
)

func newAsynqServer(c config.Config) *asynq.Server {
	return asynq.NewServer(asynq.RedisClientOpt{
		Addr:     c.RedisConf.Addr,
		DB:       c.RedisConf.DB,
		PoolSize: c.RedisConf.PoolSize,
	}, asynq.Config{
		IsFailure: func(err error) bool {
			logx.WithContext(context.Background()).Errorf("asynq server exec task IsFailure =====>>>>>  ERROR: %+v", err)
			return true
		},
		Concurrency: 20,
	})
}
