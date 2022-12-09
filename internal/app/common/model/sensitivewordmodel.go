package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SensitiveWordModel = (*customSensitiveWordModel)(nil)

type (
	// SensitiveWordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSensitiveWordModel.
	SensitiveWordModel interface {
		sensitiveWordModel
		FindAllWord(ctx context.Context) ([]string, error)
		FindOneByWord(ctx context.Context, word string) (*SensitiveWord, error)
		DeleteByWord(ctx context.Context, word string) error
	}

	customSensitiveWordModel struct {
		*defaultSensitiveWordModel
	}
)

func (m *customSensitiveWordModel) FindAllWord(ctx context.Context) ([]string, error) {
	var resp []string
	query := fmt.Sprintf("select `word` from %s", m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *customSensitiveWordModel) FindOneByWord(ctx context.Context, word string) (*SensitiveWord, error) {
	var resp SensitiveWord
	query := fmt.Sprintf("select %s from %s where `word` = ? limit 1", sensitiveWordRows, m.table)
	err := m.QueryRowNoCacheCtx(ctx, resp, query, word)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *customSensitiveWordModel) DeleteByWord(ctx context.Context, word string) error {
	sensitiveWord, err := m.FindOneByWord(ctx, word)
	if err != nil {
		return err
	}
	err = m.Delete(ctx, sensitiveWord.Id)
	if err != nil {
		return err
	}
	return nil
}

// NewSensitiveWordModel returns a model for the database table.
func NewSensitiveWordModel(conn sqlx.SqlConn, c cache.CacheConf) SensitiveWordModel {
	return &customSensitiveWordModel{
		defaultSensitiveWordModel: newSensitiveWordModel(conn, c),
	}
}
