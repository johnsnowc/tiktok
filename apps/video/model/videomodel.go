package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ VideoModel = (*customVideoModel)(nil)

type (
	// VideoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVideoModel.
	VideoModel interface {
		videoModel

		FindAllByUid(ctx context.Context, uid int64) ([]*Video, error)
	}

	customVideoModel struct {
		*defaultVideoModel
	}
)

func (c *customVideoModel) FindAllByUid(ctx context.Context, uid int64) ([]*Video, error) {
	var resp []*Video

	query := fmt.Sprintf("select %s from %s where `uid` = ?", videoRows, c.table)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, uid)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// NewVideoModel returns a model for the database table.
func NewVideoModel(conn sqlx.SqlConn, c cache.CacheConf) VideoModel {
	return &customVideoModel{
		defaultVideoModel: newVideoModel(conn, c),
	}
}
