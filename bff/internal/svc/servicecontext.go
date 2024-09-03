package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-demo/bff/internal/config"
	"go-zero-demo/post/post"
	"go-zero-demo/user/user"
)

type ServiceContext struct {
	Config  config.Config
	PostRpc post.PostClient
	UserRpc user.UserClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		PostRpc: post.NewPostClient(zrpc.MustNewClient(c.PostRpcConf).Conn()),
		UserRpc: user.NewUserClient(zrpc.MustNewClient(c.UserRpcConf).Conn()),
	}
}
