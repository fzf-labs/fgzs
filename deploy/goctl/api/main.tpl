package main

import (
	"flag"
	"fmt"

	{{.importPackages}}
)

var configFile = flag.String("f", "etc/{{.serviceName}}.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	// 初始化多服务启动器
	svcGroup := service.NewServiceGroup()
	defer svcGroup.Stop()

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	//全局中间件
	server.Use(middleware.NewArgumentMiddleware(c.Mode, &c.WebHook).Handle)
	handler.RegisterHandlers(server, ctx)
	//错误处理
	httpx.SetErrorHandler(interceptor.HttpErrorHandler)
	//添加服务
	svcGroup.Add(server)
	//添加监控服务
	svcGroup.Add(monitor.NewPyroscopeServer(&c.Pyroscope))
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	//启动多服务
	svcGroup.Start()
}
