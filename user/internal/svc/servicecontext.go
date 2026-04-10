package svc

import (
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-demo/user/internal/config"
	"go-zero-demo/user/model"
	"sync"
)

type ServiceContext struct {
	Config config.Config
	DB
	Redis redis.UniversalClient
}

var m2 sync.Map

func NewServiceContext(c config.Config) *ServiceContext {
	//var oom2 func()
	//oom2 = func() {
	//	tick := time.Tick(time.Second)
	//
	//	stime := time.Now()
	//	for range tick {
	//		// 1秒1M内存
	//		var buf []byte
	//		buf = append(buf, make([]byte, 1024*1024)...)
	//		m2.Store(time.Now().UnixMilli(), buf)
	//		fmt.Printf("%f\n", time.Now().Sub(stime).Seconds())
	//	}
	//}
	//go oom2()
	return &ServiceContext{
		Config: c,
		DB:     NewDB(sqlx.NewMysql(c.MysqlConf.DataSource)),
		Redis: redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs: []string{c.Redis.Host},
		}),
	}
}

type DB struct {
	User model.UserModel
}

func NewDB(conn sqlx.SqlConn) DB {
	return DB{
		User: model.NewUserModel(conn),
	}
}
