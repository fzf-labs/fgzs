package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SmsRecordModel = (*customSmsRecordModel)(nil)

type (
	// SmsRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsRecordModel.
	SmsRecordModel interface {
		smsRecordModel
	}

	customSmsRecordModel struct {
		*defaultSmsRecordModel
	}
)

// NewSmsRecordModel returns a model for the database table.
func NewSmsRecordModel(conn sqlx.SqlConn, c cache.CacheConf) SmsRecordModel {
	return &customSmsRecordModel{
		defaultSmsRecordModel: newSmsRecordModel(conn, c),
	}
}
