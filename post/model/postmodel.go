package model

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PostModel = (*customPostModel)(nil)

const (
	// MaxPageSize 最大分页大小，防止查询过多数据
	MaxPageSize = 100
)

type (
	// PostModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPostModel.
	PostModel interface {
		postModel
		withSession(session sqlx.Session) PostModel
		GetPostByUser(ctx context.Context, userId int64) ([]Post, error)
		GetPostByIds(ctx context.Context, ids []int64) ([]Post, error)
		FindOne(ctx context.Context, id int64) (*Post, error)
		FindOneWithoutStatus(ctx context.Context, id int64) (*Post, error)
	}

	customPostModel struct {
		*defaultPostModel
		redis *redis.Client
	}
)

// NewPostModel returns a model for the database table.
func NewPostModel(conn sqlx.SqlConn, redisClient *redis.Client) PostModel {
	return &customPostModel{
		defaultPostModel: newPostModel(conn),
		redis:            redisClient,
	}
}

func (m *customPostModel) withSession(session sqlx.Session) PostModel {
	return NewPostModel(sqlx.NewSqlConnFromSession(session), m.redis)
}

func (m *customPostModel) GetPostByUser(ctx context.Context, userId int64) ([]Post, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? and `status` = 0 order by created_at desc", postRows, m.table)
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
	// 使用参数化查询防止 SQL 注入
	placeholders := make([]string, len(ids))
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id
	}
	query := fmt.Sprintf("select %s from %s where `id` in (%s) and `status` = 0", postRows, m.table, strings.Join(placeholders, ","))
	var resp []Post
	err := m.conn.QueryRowsCtx(ctx, &resp, query, args...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// FindOne 通过 ID 查找推文（带缓存）
func (m *customPostModel) FindOne(ctx context.Context, id int64) (*Post, error) {
	// 尝试从缓存获取
	cacheKey := fmt.Sprintf("post:%d", id)
	if m.redis != nil {
		data, err := m.redis.Get(ctx, cacheKey).Bytes()
		if err == nil {
			var post Post
			if json.Unmarshal(data, &post) == nil {
				return &post, nil
			}
		}
	}

	// 缓存未命中，从数据库查询
	query := fmt.Sprintf("select %s from %s where `id` = ? and `status` = 0 limit 1", postRows, m.table)
	var resp Post
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		// 写入缓存
		if m.redis != nil {
			if data, err := json.Marshal(&resp); err == nil {
				m.redis.Set(ctx, cacheKey, data, 1*time.Minute)
			}
		}
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// FindOneWithoutStatus 通过 ID 查找推文（不带状态过滤，用于删除等操作）
func (m *customPostModel) FindOneWithoutStatus(ctx context.Context, id int64) (*Post, error) {
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

// UpdatePostCounts 更新推文计数
func (m *customPostModel) UpdatePostCounts(ctx context.Context, id int64, field string, delta int) error {
	query := fmt.Sprintf("update %s set `%s` = `%s` + ? where `id` = ?", m.table, field, field)
	_, err := m.conn.ExecCtx(ctx, query, delta, id)
	return err
}

// StringToInt64s converts string slice to int64 slice
func StringToInt64s(strs []string) []int64 {
	result := make([]int64, 0, len(strs))
	for _, s := range strs {
		if id, err := strconv.ParseInt(s, 10, 64); err == nil {
			result = append(result, id)
		}
	}
	return result
}

// SearchPosts 搜索推文（标题/内容模糊匹配）
func (m *customPostModel) SearchPosts(ctx context.Context, keyword string, limit, offset int) ([]Post, error) {
	// 限制分页大小，防止查询过多数据
	if limit <= 0 || limit > MaxPageSize {
		limit = MaxPageSize
	}
	if offset < 0 {
		offset = 0
	}
	searchPattern := "%" + escapeLikePattern(keyword) + "%"
	query := fmt.Sprintf("select %s from %s where `status` = 0 and (`title` like ? or `content` like ?) order by created_at desc limit ? offset ?", postRows, m.table)
	var resp []Post
	err := m.conn.QueryRowsCtx(ctx, &resp, query, searchPattern, searchPattern, limit, offset)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SearchPostsCount 搜索推文数量
func (m *customPostModel) SearchPostsCount(ctx context.Context, keyword string) (int64, error) {
	searchPattern := "%" + escapeLikePattern(keyword) + "%"
	query := fmt.Sprintf("select count(*) from %s where `status` = 0 and (`title` like ? or `content` like ?)", m.table)
	var count int64
	err := m.conn.QueryRowCtx(ctx, &count, query, searchPattern, searchPattern)
	return count, err
}

// escapeLikePattern 转义 LIKE 模式中的特殊字符，防止模糊搜索时被误解释为通配符
// 转义 %、_、\ 字符
func escapeLikePattern(s string) string {
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "%", "\\%")
	s = strings.ReplaceAll(s, "_", "\\_")
	return s
}
