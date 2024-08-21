package svc

import (
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-demo/post/internal/config"
	"go-zero-demo/post/model"
)

type ServiceContext struct {
	Config config.Config
	DB
	Redis redis.UniversalClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     NewDB(sqlx.NewMysql(c.MysqlConf.DataSource)),
		Redis: redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs: []string{c.Redis.Host},
		}),
	}
}

type DB struct {
	Post model.PostModel
}

func NewDB(conn sqlx.SqlConn) DB {
	return DB{
		Post: model.NewPostModel(conn),
	}
}
