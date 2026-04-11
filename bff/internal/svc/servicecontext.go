package svc

import (
	"go-zero-demo/bff/internal/config"
	"go-zero-demo/post/post"
	"go-zero-demo/user/user"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	PostRpc post.PostClient
	UserRpc user.UserClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUserClient(zrpc.MustNewClient(c.UserRpcConf).Conn()),
		PostRpc: post.NewPostClient(zrpc.MustNewClient(c.PostRpcConf).Conn()),
	}
}
