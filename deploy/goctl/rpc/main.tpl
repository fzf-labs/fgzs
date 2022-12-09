package main

import (
	"flag"
	"fmt"

	{{.imports}}

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	svr := server.New{{.serviceNew}}Server(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
{{range .serviceNames}}       {{.Pkg}}.Register{{.Service}}Server(grpcServer, {{.ServerPkg}}.New{{.Service}}Server(ctx))
{{end}}
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	//Rpc错误拦截器
	s.AddUnaryInterceptors(interceptor.RpcErrInterceptor)
	//添加rpc服务
	svcGroup.Add(s)
	//添加监控服务
	svcGroup.Add(monitor.NewPyroscopeServer(&c.Pyroscope))
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	//启动多服务
	svcGroup.Start()
}
