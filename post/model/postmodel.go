package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strconv"
	"strings"
)

var _ PostModel = (*customPostModel)(nil)

type (
	// PostModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostModel.
	PostModel interface {
		postModel
		withSession(session sqlx.Session) PostModel
		GetPostByUser(ctx context.Context, userId int64) ([]Post, error)
		GetPostByIds(ctx context.Context, Idss []int64) ([]Post, error)
	}

	customPostModel struct {
		*defaultPostModel
	}
)

// NewPostModel returns a model for the database table.
func NewPostModel(conn sqlx.SqlConn) PostModel {
	return &customPostModel{
		defaultPostModel: newPostModel(conn),
	}
}

func (m *customPostModel) withSession(session sqlx.Session) PostModel {
	return NewPostModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customPostModel) GetPostByUser(ctx context.Context, userId int64) ([]Post, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? and `status` = 0", postRows, m.table)
	var resp []Post
	err := m.conn.QueryRowsCtx(ctx, &resp, query, userId)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *customPostModel) GetPostByIds(ctx context.Context, ids []int64) ([]Post, error) {
	if len(ids) == 0 {
		return []Post{}, nil
	}
	idsStrSlice := []string{}
	for _, id := range ids {
		idsStrSlice = append(idsStrSlice, strconv.Itoa(int(id)))
	}

	query := fmt.Sprintf("select %s from %s where `id` in (%s) and `status` = 0", postRows, m.table, strings.Join(idsStrSlice, ","))
	var resp []Post
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
