package svc

import (
	"fgzs/internal/app/common/rpc/common"
	"fgzs/internal/app/identity/rpc/identity"
	model2 "fgzs/internal/app/user/model"
	"fgzs/internal/app/user/rpc/internal/config"
	"fgzs/internal/core"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Mysql  sqlx.SqlConn
	Redis  *redis.Redis

	UserModel     model2.UserModel
	UserAuthModel model2.UserAuthModel

	IdentityRpc identity.Identity
	CommonRpc   common.Common
}

func NewServiceContext(c config.Config) *ServiceContext {
	slConn := core.NewMysql(c.Mysql)
	return &ServiceContext{
		Config:        c,
		Mysql:         slConn,
		Redis:         core.NewRedis(c.Redis),
		UserModel:     model2.NewUserModel(slConn, c.Cache),
		UserAuthModel: model2.NewUserAuthModel(slConn, c.Cache),

		IdentityRpc: identity.NewIdentity(zrpc.MustNewClient(c.IdentityRpc)),
		CommonRpc:   common.NewCommon(zrpc.MustNewClient(c.CommonRpc)),
	}
}
