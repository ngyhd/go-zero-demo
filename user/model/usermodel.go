package model

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

const (
	// MaxPageSize 最大分页大小，防止查询过多数据
	MaxPageSize = 100
)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		withSession(session sqlx.Session) UserModel
		FindOneByAccount(ctx context.Context, account string) (*User, error)
		FindByIds(ctx context.Context, ids []int64) ([]*User, error)
		FindOne(ctx context.Context, id int64) (*User, error)
	}

	customUserModel struct {
		*defaultUserModel
		redis *redis.Client
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, redisClient *redis.Client) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn),
		redis:           redisClient,
	}
}

func (m *customUserModel) withSession(session sqlx.Session) UserModel {
	return NewUserModel(sqlx.NewSqlConnFromSession(session), m.redis)
}

// FindOneByAccount 通过账号查找用户信息（不带缓存）
func (m *customUserModel) FindOneByAccount(ctx context.Context, account string) (*User, error) {
	query := fmt.Sprintf("select %s from %s where `account` = ? limit 1", userRows, m.table)
	var resp User
	err := m.conn.QueryRowCtx(ctx, &resp, query, account)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// FindOne 通过 ID 查找用户（带缓存）
func (m *customUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	// 尝试从缓存获取
	cacheKey := fmt.Sprintf("user:%d", id)
	if m.redis != nil {
		data, err := m.redis.Get(ctx, cacheKey).Bytes()
		if err == nil {
			var user User
			if json.Unmarshal(data, &user) == nil {
				return &user, nil
			}
		}
	}

	// 缓存未命中，从数据库查询
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
	var resp User
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		// 写入缓存
		if m.redis != nil {
			if data, err := json.Marshal(&resp); err == nil {
				m.redis.Set(ctx, cacheKey, data, 5*time.Minute)
			}
		}
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// FindByIds 批量查询用户
func (m *customUserModel) FindByIds(ctx context.Context, ids []int64) ([]*User, error) {
	if len(ids) == 0 {
		return []*User{}, nil
	}
	placeholders := make([]string, len(ids))
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id
	}
	query := fmt.Sprintf("select %s from %s where `id` in (%s)", userRows, m.table, strings.Join(placeholders, ","))
	var resp []*User
	err := m.conn.QueryRowsCtx(ctx, &resp, query, args...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SearchUsers 搜索用户（账号/昵称模糊匹配）
func (m *customUserModel) SearchUsers(ctx context.Context, keyword string, limit, offset int) ([]*User, error) {
	// 限制分页大小，防止查询过多数据
	if limit <= 0 || limit > MaxPageSize {
		limit = MaxPageSize
	}
	if offset < 0 {
		offset = 0
	}
	searchPattern := "%" + escapeLikePattern(keyword) + "%"
	query := fmt.Sprintf("select %s from %s where `account` like ? or `nickname` like ? limit ? offset ?", userRows, m.table)
	var resp []*User
	err := m.conn.QueryRowsCtx(ctx, &resp, query, searchPattern, searchPattern, limit, offset)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SearchUsersCount 搜索用户数量
func (m *customUserModel) SearchUsersCount(ctx context.Context, keyword string) (int64, error) {
	searchPattern := "%" + escapeLikePattern(keyword) + "%"
	query := fmt.Sprintf("select count(*) from %s where `account` like ? or `nickname` like ?", m.table)
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
