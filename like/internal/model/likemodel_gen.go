package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	likeFieldNames          = "`id`, `user_id`, `post_id`, `created_at`"
	likeRows                = likeFieldNames
	likeRowsExpectAutoSet   = "`user_id`, `post_id`, `created_at`"
	likeRowsWithPlaceHolder = "`user_id`=?, `post_id`=?, `created_at`=?"
)

type (
	likeModel interface {
		Insert(ctx context.Context, data *Like) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Like, error)
		Update(ctx context.Context, data *Like) error
		Delete(ctx context.Context, id int64) error
	}

	defaultLikeModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Like struct {
		Id        int64 `db:"id"`
		UserId    int64 `db:"user_id"`
		PostId    int64 `db:"post_id"`
		CreatedAt int64 `db:"created_at"`
	}
)

func newLikeModel(conn sqlx.SqlConn) *defaultLikeModel {
	return &defaultLikeModel{
		conn:  conn,
		table: "`like`",
	}
}

func (m *defaultLikeModel) Insert(ctx context.Context, data *Like) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, likeRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.PostId, data.CreatedAt)
	return ret, err
}

func (m *defaultLikeModel) FindOne(ctx context.Context, id int64) (*Like, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", likeRows, m.table)
	var resp Like
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultLikeModel) Update(ctx context.Context, data *Like) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, likeRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserId, data.PostId, data.CreatedAt, data.Id)
	return err
}

func (m *defaultLikeModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultLikeModel) tableName() string {
	return m.table
}
