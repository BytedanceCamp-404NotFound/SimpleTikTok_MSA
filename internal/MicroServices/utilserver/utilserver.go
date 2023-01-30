package main

import (
	"flag"
	"fmt"

	"SimpleTikTok/internal/MicroServices/utilserver/internal/config"
	"SimpleTikTok/internal/MicroServices/utilserver/internal/server"
	"SimpleTikTok/internal/MicroServices/utilserver/internal/svc"
	"SimpleTikTok/internal/MicroServices/utilserver/pkg/Utilserver"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/utilserver.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		Utilserver.RegisterUtilserverServer(grpcServer, server.NewUtilserverServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
