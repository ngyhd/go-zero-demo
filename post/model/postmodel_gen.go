// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	postFieldNames          = builder.RawFieldNames(&Post{})
	postRows                = strings.Join(postFieldNames, ",")
	postRowsExpectAutoSet   = strings.Join(stringx.Remove(postFieldNames, "`id`", "`-i`"), ",")
	postRowsWithPlaceHolder = strings.Join(stringx.Remove(postFieldNames, "`id`", "`-i`"), "=?,") + "=?"
)

type (
	postModel interface {
		Insert(ctx context.Context, data *Post) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Post, error)
		Update(ctx context.Context, data *Post) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPostModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Post struct {
		Id        int64  `db:"id"`         // 主键id
		UserId    int64  `db:"user_id"`    // 作者id
		Title     string `db:"title"`      // 标题
		Content   string `db:"content"`    // 推文内容
		Status    int64  `db:"status"`     // 状态 0 正常 1 已删除
		Views     int64  `db:"views"`      // 浏览数
		Likes     int64  `db:"likes"`      // 点赞数
		Comments  int64  `db:"comments"`   // 评论数
		Shares    int64  `db:"shares"`     // 分享数
		Collects  int64  `db:"collects"`   // 收藏数
		CreatedAt int64  `db:"created_at"` // 发表时间
		UpdatedAt int64  `db:"updated_at"` // 更新时间
		DeletedAt int64  `db:"deleted_at"` // 删除时间
	}
)

func newPostModel(conn sqlx.SqlConn) *defaultPostModel {
	return &defaultPostModel{
		conn:  conn,
		table: "`post`",
	}
}

func (m *defaultPostModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultPostModel) FindOne(ctx context.Context, id int64) (*Post, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", postRows, m.table)
	var resp Post
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

func (m *defaultPostModel) Insert(ctx context.Context, data *Post) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, postRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.Title, data.Content, data.Status, data.Views, data.Likes, data.Comments, data.Shares, data.Collects, data.CreatedAt, data.UpdatedAt, data.DeletedAt)
	return ret, err
}

func (m *defaultPostModel) Update(ctx context.Context, data *Post) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, postRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserId, data.Title, data.Content, data.Status, data.Views, data.Likes, data.Comments, data.Shares, data.Collects, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.Id)
	return err
}

func (m *defaultPostModel) tableName() string {
	return m.table
}
