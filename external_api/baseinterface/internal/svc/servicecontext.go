package svc

import (
	"SimpleTikTok/external_api/baseinterface/internal/config"
	"SimpleTikTok/internal_proto/microservices/miniomanage/miniomanageserverclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	MinioManageRpc miniomanageserverclient.MinioManageServer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		MinioManageRpc: miniomanageserverclient.NewMinioManageServer(zrpc.MustNewClient(c.MinioManageRpc)),
	}
}
