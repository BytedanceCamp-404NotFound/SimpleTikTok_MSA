// Code generated by goctl. DO NOT EDIT.
// Source: mongodbmanage.proto

package server

import (
	"context"

	"SimpleTikTok/internal_proto/microservices/mongodbmanage/internal/logic"
	"SimpleTikTok/internal_proto/microservices/mongodbmanage/internal/svc"
	"SimpleTikTok/internal_proto/microservices/mongodbmanage/pkg/MongodbManageServer"
)

type MongodbManageServerServer struct {
	svcCtx *svc.ServiceContext
	MongodbManageServer.UnimplementedMongodbManageServerServer
}

func NewMongodbManageServerServer(svcCtx *svc.ServiceContext) *MongodbManageServerServer {
	return &MongodbManageServerServer{
		svcCtx: svcCtx,
	}
}

func (s *MongodbManageServerServer) GetMinio(ctx context.Context, in *MongodbManageServer.IdRequest) (*MongodbManageServer.MinioResponse, error) {
	l := logic.NewGetMinioLogic(ctx, s.svcCtx)
	return l.GetMinio(in)
}