package logic

import (
	"genuine_douyin/apps/asynqJob/server/jobtype"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
)

// GetUserFavoriteStatusScheduler 向Redis发送定时消息调用worker进行工作
func (l *MqueueScheduler) GetUserFavoriteStatusScheduler() {

	task := asynq.NewTask(jobtype.ScheduleGetUserFavoriteStatus, nil)
	// every one minute exec
	entryID, err := l.svcCtx.Scheduler.Register("@every 10s", task)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("!!!MqueueSchedulerErr!!! ====> 【ScheduleGetUserFavoriteStatus】 registered  err:%+v , task:%+v", err, task)
	}
	logx.Infof("【ScheduleGetUserFavoriteStatus】 registered an  entry: %q \n", entryID)
}

func (l *MqueueScheduler) GetUserFollowStatusScheduler() {

	task := asynq.NewTask(jobtype.ScheduleGetUserFollowStatus, nil)
	// every one minute exec
	entryID, err := l.svcCtx.Scheduler.Register("@every 10s", task)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("!!!MqueueSchedulerErr!!! ====> 【ScheduleGetUserFollowStatus】 registered  err:%+v , task:%+v", err, task)
	}
	logx.Infof("【ScheduleGetUserFollowStatus】 registered an  entry: %q \n", entryID)
}
