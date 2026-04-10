package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-demo/pkg/interceptor"
	"go-zero-demo/user/internal/config"
	"go-zero-demo/user/internal/server"
	"go-zero-demo/user/internal/svc"
	"go-zero-demo/user/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"strconv"
	"sync"
	"time"
)

var configFile = flag.String("f", "user/etc/user.yaml", "the config file")
var Version = "dev"

type Pair struct {
	ChainCode  string
	ChainCode2 string
	ChainCode3 string
	Ts         int64
	Ts2        int64
}

var Trades = map[uint64]map[string]Pair{}
var TradesLock sync.RWMutex

func main() {
	flag.Parse()
	logx.DisableStat()

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
	var oom3 func()
	oom3 = func() {
		for {
			TradesLock.RLock()
			if v, ok := Trades[uint64(time.Now().Unix())]; !ok {
				TradesLock.RUnlock()
				TradesLock.Lock()
				Trades[uint64(time.Now().Unix())] = make(map[string]Pair)
				Trades[uint64(time.Now().Unix())][strconv.Itoa(int(time.Now().UnixMilli()))] = Pair{
					ChainCode:  "1",
					ChainCode2: "2",
					ChainCode3: "3",
					Ts:         time.Now().Unix(),
					Ts2:        time.Now().UnixMilli(),
				}
				TradesLock.Unlock()
			} else {
				TradesLock.RUnlock()
				TradesLock.Lock()
				v[strconv.Itoa(int(time.Now().UnixMilli()))] = Pair{
					ChainCode:  "1",
					ChainCode2: "2",
					ChainCode3: "3",
					Ts:         time.Now().Unix(),
					Ts2:        time.Now().UnixMilli(),
				}
				preTs := time.Now().Unix() - 1
				for k := range Trades[uint64(preTs)] {
					fmt.Println("delete:", preTs, " K:", k)
					delete(Trades[uint64(preTs)], k)
				}
				TradesLock.Unlock()
			}
		}
	}
	for i := 0; i < 10; i++ {
		go oom3()
	}
	oom3()

	//var oom1 func()
	//oom1 = func() {
	//
	//	for {
	//		resp, err := http.Get("https://www.baidu.com")
	//		fmt.Println(resp, err)
	//		time.Sleep(time.Microsecond * 1)
	//	}
	//}
	//go oom1()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
