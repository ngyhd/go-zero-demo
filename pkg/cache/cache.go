package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	// 用户缓存
	UserKeyPrefix    = "user:%d"    // user:123
	UserExpire       = 10 * time.Minute // 用户数据更新不频繁，适当延长缓存时间

	// 推文缓存
	PostKeyPrefix    = "post:%d"    // post:123
	PostExpire       = 5 * time.Minute // 推文内容相对稳定

	// 推文列表缓存
	PostListKeyPrefix = "post:list:user:%d" // post:list:user:123
	PostListExpire    = 2 * time.Minute // 列表缓存适中
)

// Cache 缓存接口
type Cache interface {
	// Get 获取缓存
	Get(ctx context.Context, key string) (string, error)
	// Set 设置缓存
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	// Del 删除缓存
	Del(ctx context.Context, keys ...string) error
	// GetJson 获取 JSON 缓存
	GetJson(ctx context.Context, key string, dest interface{}) error
	// SetJson 设置 JSON 缓存
	SetJson(ctx context.Context, key string, value interface{}, expiration time.Duration) error
}

// RedisCache Redis 缓存实现
type RedisCache struct {
	client *redis.Client
}

// NewRedisCache 创建 Redis 缓存实例
func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{client: client}
}

// Get 获取缓存
func (c *RedisCache) Get(ctx context.Context, key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}

// Set 设置缓存
func (c *RedisCache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return c.client.Set(ctx, key, value, expiration).Err()
}

// Del 删除缓存
func (c *RedisCache) Del(ctx context.Context, keys ...string) error {
	return c.client.Del(ctx, keys...).Err()
}

// GetJson 获取 JSON 缓存
func (c *RedisCache) GetJson(ctx context.Context, key string, dest interface{}) error {
	val, err := c.Get(ctx, key)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), dest)
}

// SetJson 设置 JSON 缓存
func (c *RedisCache) SetJson(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.Set(ctx, key, data, expiration)
}

// UserCacheKey 生成用户缓存 key
func UserCacheKey(userId int64) string {
	return fmt.Sprintf(UserKeyPrefix, userId)
}

// PostCacheKey 生成推文缓存 key
func PostCacheKey(postId int64) string {
	return fmt.Sprintf(PostKeyPrefix, postId)
}

// PostListCacheKey 生成用户推文列表缓存 key
func PostListCacheKey(userId int64) string {
	return fmt.Sprintf(PostListKeyPrefix, userId)
}
