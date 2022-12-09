package config

import (
	"fgzs/internal/core"
	"fgzs/internal/monitor"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Pyroscope monitor.PyroscopeConfig
	Mysql     core.MysqlConfig
	Cache     cache.CacheConf
}
