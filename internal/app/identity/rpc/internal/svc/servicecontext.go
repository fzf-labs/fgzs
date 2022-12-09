package svc

import (
	"fgzs/internal/app/identity/rpc/internal/config"
	"fgzs/internal/core"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config config.Config
	Redis  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Redis:  core.NewRedis(c.Redis),
	}
}
