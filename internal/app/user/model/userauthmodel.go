package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserAuthModel = (*customUserAuthModel)(nil)

type (
	// UserAuthModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserAuthModel.
	UserAuthModel interface {
		userAuthModel
		FindOneByIdentity(ctx context.Context, IdentityType string, IdentityKey string, IdentityCode string, status int64) (*UserAuth, error)
	}

	customUserAuthModel struct {
		*defaultUserAuthModel
	}
)

func (m *customUserAuthModel) FindOneByIdentity(ctx context.Context, IdentityType string, IdentityKey string, IdentityCode string, status int64) (*UserAuth, error) {
	var resp UserAuth
	query := fmt.Sprintf("SELECT %s FROM %s WHERE `identity_type` = ? AND `identity_key` = ? AND `identity_code` = ? AND status = ? LIMIT 1", userAuthRows, m.table)
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, IdentityType, IdentityKey, IdentityCode, status)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// NewUserAuthModel returns a model for the database table.
func NewUserAuthModel(conn sqlx.SqlConn, c cache.CacheConf) UserAuthModel {
	return &customUserAuthModel{
		defaultUserAuthModel: newUserAuthModel(conn, c),
	}
}
