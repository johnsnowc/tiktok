package main

import (
	"context"
	"flag"
	"genuine_douyin/apps/asynqJob/client-scheduler/internal/config"
	"genuine_douyin/apps/asynqJob/client-scheduler/internal/logic"
	"genuine_douyin/apps/asynqJob/client-scheduler/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"os"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/scheduler.yaml", "Specify the config file")

func main() {
	flag.Parse()
	var c config.Config

	conf.MustLoad(*configFile, &c)

	logx.DisableStat()
	// log、prometheus、trace、metricsUrl.
	if err := c.SetUp(); err != nil {
		panic(err)
	}

	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()
	mqueueScheduler := logic.NewCronScheduler(ctx, svcContext)
	mqueueScheduler.Register()

	if err := svcContext.Scheduler.Run(); err != nil {
		logx.Errorf("!!!MqueueSchedulerErr!!!  run err:%+v", err)
		os.Exit(1)
	}
}
