package main

import (
	"flag"
	"fmt"

	"SimpleTikTok/internal/MicroServices/mysqlmanage/internal/config"
	"SimpleTikTok/internal/MicroServices/mysqlmanage/internal/server"
	"SimpleTikTok/internal/MicroServices/mysqlmanage/internal/svc"
	"SimpleTikTok/internal/MicroServices/mysqlmanage/pkg/MySQLManageServer"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/mysqlmanage.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		MySQLManageServer.RegisterMySQLManageServerServer(grpcServer, server.NewMySQLManageServerServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
