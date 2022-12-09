package main

import (
	"fgzs/internal/app/user/rpc/internal/config"
	server2 "fgzs/internal/app/user/rpc/internal/server"
	"fgzs/internal/app/user/rpc/internal/svc"
	"fgzs/internal/app/user/rpc/userpb"
	"fgzs/internal/interceptor"
	"fgzs/internal/monitor"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	//禁用统计日志
	logx.DisableStat()
	// 初始化多服务启动器
	svcGroup := service.NewServiceGroup()
	defer svcGroup.Stop()

	ctx := svc.NewServiceContext(c)
	svr := server2.NewUserServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		userpb.RegisterUserServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	//Rpc错误拦截器
	s.AddUnaryInterceptors(interceptor.RpcErrInterceptor)
	//添加rpc服务
	svcGroup.Add(s)
	//添加corn
	svcGroup.Add(server2.NewCornService(ctx))
	//添加监控服务
	svcGroup.Add(monitor.NewPyroscopeServer(&c.Pyroscope))
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	//启动多服务
	svcGroup.Start()
}
