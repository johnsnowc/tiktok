package logic

import (
	"context"
	"genuine_douyin/apps/asynqJob/server/internal/logic/jobs"
	"genuine_douyin/apps/asynqJob/server/internal/svc"
	"genuine_douyin/apps/asynqJob/server/jobtype"
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

// register server
func (l *CronJob) Register() *asynq.ServeMux {

	mux := asynq.NewServeMux()

	// handle
	mux.Handle(jobtype.ScheduleGetUserFavoriteStatus, jobs.NewGetUserFavoriteStatusHandler(l.svcCtx))
	mux.Handle(jobtype.ScheduleGetUserFollowStatus, jobs.NewGetUserFollowStatusHandler(l.svcCtx))

	return mux
}
