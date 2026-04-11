package svc

import (
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-demo/like/internal/config"
	"go-zero-demo/like/internal/model"
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
	fmt.Println("redisClient:", c.MysqlConf.DataSource)
	return &ServiceContext{
		Config: c,
		DB:     NewDB(sqlx.NewMysql(c.MysqlConf.DataSource)),
		Redis:  redisClient,
	}
}

type DB struct {
	Like model.LikeModel
}

func NewDB(conn sqlx.SqlConn) DB {
	return DB{
		Like: model.NewLikeModel(conn),
	}
}
