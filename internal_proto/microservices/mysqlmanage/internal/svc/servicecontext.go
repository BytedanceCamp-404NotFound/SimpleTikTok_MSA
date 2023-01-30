package svc

import "SimpleTikTok/internal_proto/microservices/mysqlmanage/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
