package main

import (
	"flag"
	"fmt"

	"SimpleTikTok/internal/MicroServices/minio/internal/config"
	"SimpleTikTok/internal/MicroServices/minio/internal/server"
	"SimpleTikTok/internal/MicroServices/minio/internal/svc"
	"SimpleTikTok/internal/MicroServices/pkg/Minio"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/minio.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		Minio.RegisterMinioServer(grpcServer, server.NewMinioServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
