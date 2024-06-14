package logic

import (
	"GoBao/server/mq/job/internal/svc"
	"GoBao/server/mq/job/jobtype"
	"context"
	"github.com/hibiken/asynq"
)

type CronJob struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCronJob(ctx context.Context, svcCtx *svc.ServiceContext) *CronJob {
	return &CronJob{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CronJob) Register() *asynq.ServeMux {
	mux := asynq.NewServeMux()

	// 延迟任务队列
	mux.Handle(jobtype.DeferCloseSeckillOrder, NewCloseSeckillOrderHandler(l.svcCtx))

	return mux
}
