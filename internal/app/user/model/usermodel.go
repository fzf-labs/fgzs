package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

var (
	cacheFgzsUsernameAndPltOrZonePrefix = "cache:fgzsUser:user:usernameandpltorzone:"
)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c),
	}
}

func (m *customUserModel) FindOneByUsernameAndPltOrZone(ctx context.Context, username string, plt int, zone string) (*User, error) {
	cacheFgzsUsernameAndPltOrZonekey := fmt.Sprintf("%s%s%d%s", cacheFgzsUsernameAndPltOrZonePrefix, username, plt, zone)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, cacheFgzsUsernameAndPltOrZonekey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `username` = ? and `glt`= ? and `zone`= ? limit 1", userRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, username, plt, zone)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
