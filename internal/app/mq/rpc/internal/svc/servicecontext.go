package svc

import (
	"fgzs/internal/app/mq/rpc/internal/config"
	"fgzs/internal/app/user/rpc/user"
	"fgzs/pkg/mq"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	MqClient mq.IMQ

	UserClient user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		MqClient:   mq.NewMq(&c.MqConfig),
		UserClient: user.NewUser(zrpc.MustNewClient(c.MemberRpc)),
	}
}
