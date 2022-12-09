package config

import (
	"fgzs/internal/monitor"
	"fgzs/pkg/jwt"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Pyroscope monitor.PyroscopeConfig

	Jwt map[string]jwt.Config
}
