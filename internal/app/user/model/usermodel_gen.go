// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheFgzsUserUserIdPrefix       = "cache:fgzsUser:user:id:"
	cacheFgzsUserUserEmailPrefix    = "cache:fgzsUser:user:email:"
	cacheFgzsUserUserPhonePrefix    = "cache:fgzsUser:user:phone:"
	cacheFgzsUserUserUidPrefix      = "cache:fgzsUser:user:uid:"
	cacheFgzsUserUserUsernamePrefix = "cache:fgzsUser:user:username:"
)

type (
	userModel interface {
		Insert(ctx context.Context, data *User) (sql.Result, error)
		InsertTx(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*User, error)
		FindOneByEmail(ctx context.Context, email string) (*User, error)
		FindOneByPhone(ctx context.Context, phone string) (*User, error)
		FindOneByUid(ctx context.Context, uid string) (*User, error)
		FindOneByUsername(ctx context.Context, username string) (*User, error)
		Update(ctx context.Context, data *User) error
		UpdateTx(ctx context.Context, session sqlx.Session, data *User) error
		Delete(ctx context.Context, id int64) error
		DeleteTx(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		Id         int64          `db:"id"`          // 用户ID
		Uid        string         `db:"uid"`         // uid
		Username   string         `db:"username"`    // 用户名
		Password   string         `db:"password"`    // 密码
		Nickname   string         `db:"nickname"`    // 昵称
		Phone      string         `db:"phone"`       // 手机号
		Email      string         `db:"email"`       // 邮箱
		Sex        int64          `db:"sex"`         // 性别0未知1男2女
		Other      sql.NullString `db:"other"`       // 其他
		Status     int64          `db:"status"`      // 状态 1 正常 2锁定 3注销
		CreateTime time.Time      `db:"create_time"` // 创建时间
		UpdateTime time.Time      `db:"update_time"` // 更新时间
	}
)

func newUserModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user`",
	}
}

func (m *defaultUserModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	fgzsUserUserEmailKey := fmt.Sprintf("%s%v", cacheFgzsUserUserEmailPrefix, data.Email)
	fgzsUserUserIdKey := fmt.Sprintf("%s%v", cacheFgzsUserUserIdPrefix, id)
	fgzsUserUserPhoneKey := fmt.Sprintf("%s%v", cacheFgzsUserUserPhonePrefix, data.Phone)
	fgzsUserUserUidKey := fmt.Sprintf("%s%v", cacheFgzsUserUserUidPrefix, data.Uid)
	fgzsUserUserUsernameKey := fmt.Sprintf("%s%v", cacheFgzsUserUserUsernamePrefix, data.Username)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, fgzsUserUserEmailKey, fgzsUserUserIdKey, fgzsUserUserPhoneKey, fgzsUserUserUidKey, fgzsUserUserUsernameKey)
	return err
}

func (m *defaultUserModel) DeleteTx(ctx context.Context, session sqlx.Session, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	fgzsUserUserEmailKey := fmt.Sprintf("%s%v", cacheFgzsUserUserEmailPrefix, data.Email)
	fgzsUserUserIdKey := fmt.Sprintf("%s%v", cacheFgzsUserUserIdPrefix, id)
	fgzsUserUserPhoneKey := fmt.Sprintf("%s%v", cacheFgzsUserUserPhonePrefix, data.Phone)
	fgzsUserUserUidKey := fmt.Sprintf("%s%v", cacheFgzsUserUserUidPrefix, data.Uid)
	fgzsUserUserUsernameKey := fmt.Sprintf("%s%v", cacheFgzsUserUserUsernamePrefix, data.Username)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return session.ExecCtx(ctx, query, id)
	}, fgzsUserUserEmailKey, fgzsUserUserIdKey, fgzsUserUserPhoneKey, fgzsUserUserUidKey, fgzsUserUserUsernameKey)
	return err
}

func (m *defaultUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	fgzsUserUserIdKey := fmt.Sprintf("%s%v", cacheFgzsUserUserIdPrefix, id)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, fgzsUserUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
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

func (m *defaultUserModel) FindOneByEmail(ctx context.Context, email string) (*User, error) {
	fgzsUserUserEmailKey := fmt.Sprintf("%s%v", cacheFgzsUserUserEmailPrefix, email)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, fgzsUserUserEmailKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `email` = ? limit 1", userRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, email); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByPhone(ctx context.Context, phone string) (*User, error) {
	fgzsUserUserPhoneKey := fmt.Sprintf("%s%v", cacheFgzsUserUserPhonePrefix, phone)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, fgzsUserUserPhoneKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `phone` = ? limit 1", userRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, phone); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByUid(ctx context.Context, uid string) (*User, error) {
	fgzsUserUserUidKey := fmt.Sprintf("%s%v", cacheFgzsUserUserUidPrefix, uid)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, fgzsUserUserUidKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `uid` = ? limit 1", userRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, uid); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByUsername(ctx context.Context, username string) (*User, error) {
	fgzsUserUserUsernameKey := fmt.Sprintf("%s%v", cacheFgzsUserUserUsernamePrefix, username)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, fgzsUserUserUsernameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `username` = ? limit 1", userRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, username); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, data *User) (sql.Result, error) {
	fgzsUserUserEmailKey := fmt.Sprintf("%s%v", cacheFgzsUserUserEmailPrefix, data.Email)
	fgzsUserUserIdKey := fmt.Sprintf("%s%v", cacheFgzsUserUserIdPrefix, data.Id)
	fgzsUserUserPhoneKey := fmt.Sprintf("%s%v", cacheFgzsUserUserPhonePrefix, data.Phone)
	fgzsUserUserUidKey := fmt.Sprintf("%s%v", cacheFgzsUserUserUidPrefix, data.Uid)
	fgzsUserUserUsernameKey := fmt.Sprintf("%s%v", cacheFgzsUserUserUsernamePrefix, data.Username)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uid, data.Username, data.Password, data.Nickname, data.Phone, data.Email, data.Sex, data.Other, data.Status)
	}, fgzsUserUserEmailKey, fgzsUserUserIdKey, fgzsUserUserPhoneKey, fgzsUserUserUidKey, fgzsUserUserUsernameKey)
	return ret, err
}

func (m *defaultUserModel) InsertTx(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error) {
	fgzsUserUserEmailKey := fmt.Sprintf("%s%v", cacheFgzsUserUserEmailPrefix, data.Email)
	fgzsUserUserIdKey := fmt.Sprintf("%s%v", cacheFgzsUserUserIdPrefix, data.Id)
	fgzsUserUserPhoneKey := fmt.Sprintf("%s%v", cacheFgzsUserUserPhonePrefix, data.Phone)
	fgzsUserUserUidKey := fmt.Sprintf("%s%v", cacheFgzsUserUserUidPrefix, data.Uid)
	fgzsUserUserUsernameKey := fmt.Sprintf("%s%v", cacheFgzsUserUserUsernamePrefix, data.Username)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, data.Uid, data.Username, data.Password, data.Nickname, data.Phone, data.Email, data.Sex, data.Other, data.Status)
	}, fgzsUserUserEmailKey, fgzsUserUserIdKey, fgzsUserUserPhoneKey, fgzsUserUserUidKey, fgzsUserUserUsernameKey)
	return ret, err
}

func (m *defaultUserModel) Update(ctx context.Context, newData *User) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	fgzsUserUserEmailKey := fmt.Sprintf("%s%v", cacheFgzsUserUserEmailPrefix, data.Email)
	fgzsUserUserIdKey := fmt.Sprintf("%s%v", cacheFgzsUserUserIdPrefix, data.Id)
	fgzsUserUserPhoneKey := fmt.Sprintf("%s%v", cacheFgzsUserUserPhonePrefix, data.Phone)
	fgzsUserUserUidKey := fmt.Sprintf("%s%v", cacheFgzsUserUserUidPrefix, data.Uid)
	fgzsUserUserUsernameKey := fmt.Sprintf("%s%v", cacheFgzsUserUserUsernamePrefix, data.Username)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Uid, newData.Username, newData.Password, newData.Nickname, newData.Phone, newData.Email, newData.Sex, newData.Other, newData.Status, newData.Id)
	}, fgzsUserUserEmailKey, fgzsUserUserIdKey, fgzsUserUserPhoneKey, fgzsUserUserUidKey, fgzsUserUserUsernameKey)
	return err
}

func (m *defaultUserModel) UpdateTx(ctx context.Context, session sqlx.Session, newData *User) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	fgzsUserUserEmailKey := fmt.Sprintf("%s%v", cacheFgzsUserUserEmailPrefix, data.Email)
	fgzsUserUserIdKey := fmt.Sprintf("%s%v", cacheFgzsUserUserIdPrefix, data.Id)
	fgzsUserUserPhoneKey := fmt.Sprintf("%s%v", cacheFgzsUserUserPhonePrefix, data.Phone)
	fgzsUserUserUidKey := fmt.Sprintf("%s%v", cacheFgzsUserUserUidPrefix, data.Uid)
	fgzsUserUserUsernameKey := fmt.Sprintf("%s%v", cacheFgzsUserUserUsernamePrefix, data.Username)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
		return session.ExecCtx(ctx, query, newData.Uid, newData.Username, newData.Password, newData.Nickname, newData.Phone, newData.Email, newData.Sex, newData.Other, newData.Status, newData.Id)
	}, fgzsUserUserEmailKey, fgzsUserUserIdKey, fgzsUserUserPhoneKey, fgzsUserUserUidKey, fgzsUserUserUsernameKey)
	return err
}

func (m *defaultUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheFgzsUserUserIdPrefix, primary)
}

func (m *defaultUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
