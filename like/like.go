package main

import (
	"flag"
	"fmt"

	"go-zero-demo/like/internal/config"
	"go-zero-demo/like/internal/server"
	"go-zero-demo/like/internal/svc"
	"go-zero-demo/like/like"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "like/etc/like.yaml", "the config file")
var Version = "dev"

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		like.RegisterLikeServer(grpcServer, server.NewLikeServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	defer s.Stop()

	fmt.Printf("Starting like rpc server at %s... (version: %s)\n", c.ListenOn, Version)
	s.Start()
}
