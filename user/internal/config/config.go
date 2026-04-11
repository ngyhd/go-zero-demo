package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	MysqlConf MysqlConf
	RedisConf RedisConf
	JwtAuth   struct {
		AccessSecret string
		AccessExpire int64
	}
}

type MysqlConf struct {
	DataSource string
}

type RedisConf struct {
	Host         string
	Password     string
	DialTimeout  int // 拨号超时时间，单位毫秒
	ReadTimeout  int // 读取超时时间，单位毫秒
	WriteTimeout int // 写入超时时间，单位毫秒
}
