package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-demo/user/internal/config"
	"go-zero-demo/user/model"
)

type ServiceContext struct {
	Config config.Config
	DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     NewDB(sqlx.NewMysql(c.MysqlConf.DataSource)),
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
