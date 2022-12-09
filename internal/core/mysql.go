package core

import "github.com/zeromicro/go-zero/core/stores/sqlx"

type MysqlConfig struct {
	DataSourceName string
}

// NewMysql 实例化mysql
func NewMysql(conf MysqlConfig) sqlx.SqlConn {
	return sqlx.NewMysql(conf.DataSourceName)
}
