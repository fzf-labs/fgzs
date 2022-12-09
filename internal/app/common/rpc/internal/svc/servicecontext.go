package svc

import (
	model2 "fgzs/internal/app/common/model"
	"fgzs/internal/app/common/rpc/internal/config"
	"fgzs/internal/core"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config             config.Config
	Redis              *redis.Redis
	SensitiveWordModel model2.SensitiveWordModel
	SmsRecordModel     model2.SmsRecordModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := core.NewMysql(c.Mysql)
	return &ServiceContext{
		Config:             c,
		Redis:              core.NewRedis(c.Redis),
		SmsRecordModel:     model2.NewSmsRecordModel(sqlConn, c.Cache),
		SensitiveWordModel: model2.NewSensitiveWordModel(sqlConn, c.Cache),
	}
}
