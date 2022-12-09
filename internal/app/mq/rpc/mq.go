package main

import (
	"context"
	"fgzs/internal/app/mq/rpc/internal/config"
	"fgzs/internal/app/mq/rpc/internal/server"
	"fgzs/internal/app/mq/rpc/internal/svc"
	"fgzs/internal/monitor"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
)

var configFile = flag.String("f", "etc/mq.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	// log、prometheus、trace、metricsUrl.
	if err := c.SetUp(); err != nil {
		panic(err)
	}

	svcGroup := service.NewServiceGroup()
	defer svcGroup.Stop()

	ctx := context.Background()
	svcContext := svc.NewServiceContext(c)

	svcGroup.Add(server.NewMqService(ctx, svcContext))
	//添加监控服务
	svcGroup.Add(monitor.NewPyroscopeServer(&c.Pyroscope))
	fmt.Printf("Starting mq server at %s...\n", c.ListenOn)
	//启动多服务
	svcGroup.Start()
}
