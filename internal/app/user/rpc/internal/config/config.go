package config

import (
	"fgzs/internal/core"
	"fgzs/internal/monitor"
	"fgzs/pkg/apple"
	"fgzs/pkg/qq"
	"fgzs/pkg/wechat"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Pyroscope             monitor.PyroscopeConfig
	Mysql                 core.MysqlConfig
	Cache                 cache.CacheConf
	IdentityRpc           zrpc.RpcClientConf
	CommonRpc             zrpc.RpcClientConf
	QQConfig              qq.AuthConfig
	SignInWithApple       apple.SignInWithAppleConfig
	WechatOfficialAccount wechat.OfficialAccountConfig
}
