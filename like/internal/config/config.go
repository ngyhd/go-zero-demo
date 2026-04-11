package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	MysqlConf MysqlConf
	RedisConf RedisConf
}

type MysqlConf struct {
	DataSource string
}

type RedisConf struct {
	Host         string
	Password     string
	DialTimeout  int
	ReadTimeout  int
	WriteTimeout int
}
