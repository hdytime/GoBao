package main

import (
	"GoBao/server/mq/job/internal/config"
	"GoBao/server/mq/job/internal/logic"
	"GoBao/server/mq/job/internal/svc"
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"

	"os"
)

var configFile = flag.String("f", "etc/mqueue.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	if err := c.SetUp(); err != nil {
		logx.WithContext(context.Background()).Errorf("mqueue ERROR: %+v", err)
		panic(err)
	}

	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()
	cronJob := logic.NewCronJob(ctx, svcContext)
	mux := cronJob.Register()

	if err := svcContext.AsynqServer.Run(mux); err != nil {
		logx.WithContext(ctx).Errorf("ASYNQ CRONJOB ERROR: %+v", err)
		os.Exit(1)
	}
}
