package svc

import (
	"SimpleTikTok/external_api/baseinterface/internal/config"
	"SimpleTikTok/internal_proto/microservices/miniomanage/miniomanageserverclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	MinioManageRpc zrpc.RpcClientConf
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		MinioManageRpc: miniomanageserverclient.PutFileUploader(zrpc.MustNewClient(c.MinioManageRpc)),
	}
}
