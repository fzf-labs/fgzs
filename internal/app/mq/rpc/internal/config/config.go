package config

import (
	"fgzs/internal/monitor"
	"fgzs/pkg/mq"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	MqConfig     mq.MqConfig
	Pyroscope    monitor.PyroscopeConfig
	MemberRpc    zrpc.RpcClientConf
	ShortBookRpc zrpc.RpcClientConf
	MessageRpc   zrpc.RpcClientConf
}
