package svc

import (
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-demo/post/internal/config"
	"go-zero-demo/post/model"
)

type ServiceContext struct {
	Config config.Config
	DB
	Redis *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisClient := redis.NewClient(&redis.Options{
		Addr:         c.RedisConf.Host,
		Password:     c.RedisConf.Password,
		PoolSize:     100,
		MinIdleConns: 10,
		DialTimeout:  time.Duration(c.RedisConf.DialTimeout) * time.Millisecond,
		ReadTimeout:  time.Duration(c.RedisConf.ReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(c.RedisConf.WriteTimeout) * time.Millisecond,
	})

	return &ServiceContext{
		Config: c,
		DB:     NewDB(sqlx.NewMysql(c.MysqlConf.DataSource), redisClient),
		Redis:  redisClient,
	}
}

type DB struct {
	Post model.PostModel
}

func NewDB(conn sqlx.SqlConn, redisClient *redis.Client) DB {
	return DB{
		Post: model.NewPostModel(conn, redisClient),
	}
}
