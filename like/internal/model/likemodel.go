package model

import (
	"context"
	"database/sql"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LikeModel = (*customLikeModel)(nil)

type (
	LikeModel interface {
		likeModel
		FindByUserAndPost(ctx context.Context, userId, postId int64) (*Like, error)
		FindByPost(ctx context.Context, postId int64) ([]*Like, error)
		Insert(ctx context.Context, data *Like) (sql.Result, error)
		DeleteByUserAndPost(ctx context.Context, userId, postId int64) error
		CountByPost(ctx context.Context, postId int64) (int64, error)
	}

	customLikeModel struct {
		*defaultLikeModel
	}
)

func NewLikeModel(conn sqlx.SqlConn) LikeModel {
	return &customLikeModel{
		defaultLikeModel: newLikeModel(conn),
	}
}

func (m *customLikeModel) FindByUserAndPost(ctx context.Context, userId, postId int64) (*Like, error) {
	query := "select `id`, `user_id`, `post_id`, `created_at` from `like` where `user_id` = ? and `post_id` = ? limit 1"
	var resp Like
	err := m.conn.QueryRowCtx(ctx, &resp, query, userId, postId)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customLikeModel) FindByPost(ctx context.Context, postId int64) ([]*Like, error) {
	query := "select `id`, `user_id`, `post_id`, `created_at` from `like` where `post_id` = ? order by `created_at` desc"
	var resp []*Like
	err := m.conn.QueryRowsCtx(ctx, &resp, query, postId)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customLikeModel) Insert(ctx context.Context, data *Like) (sql.Result, error) {
	query := "insert into `like` (`user_id`, `post_id`, `created_at`) values (?, ?, ?)"
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.PostId, data.CreatedAt)
	return ret, err
}

func (m *customLikeModel) DeleteByUserAndPost(ctx context.Context, userId, postId int64) error {
	query := "delete from `like` where `user_id` = ? and `post_id` = ?"
	_, err := m.conn.ExecCtx(ctx, query, userId, postId)
	return err
}

func (m *customLikeModel) CountByPost(ctx context.Context, postId int64) (int64, error) {
	query := "select count(*) from `like` where `post_id` = ?"
	var count int64
	err := m.conn.QueryRowCtx(ctx, &count, query, postId)
	return count, err
}
