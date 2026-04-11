package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-demo/pkg/interceptor"
	"go-zero-demo/user/internal/config"
	"go-zero-demo/user/internal/server"
	"go-zero-demo/user/internal/svc"
	"go-zero-demo/user/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "user/etc/user.yaml", "the config file")
var Version = "dev"

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	// 拦截器--记录错误日志
	s.AddUnaryInterceptors(interceptor.Logger)

	defer s.Stop()

	fmt.Printf("Starting rpc server at %s... (version: %s)\n", c.ListenOn, Version)
	s.Start()
}
