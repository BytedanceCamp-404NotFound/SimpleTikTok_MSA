// Code generated by goctl. DO NOT EDIT.
// Source: mongodbmanage.proto

package server

import (
	"context"

	"SimpleTikTok/internal/MicroServices/mongodbmanage/internal/logic"
	"SimpleTikTok/internal/MicroServices/mongodbmanage/internal/svc"
	"SimpleTikTok/internal/MicroServices/mongodbmanage/pkg/MongodbManageServer"
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