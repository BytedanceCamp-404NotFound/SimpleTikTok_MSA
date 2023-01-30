package logic

import (
	"context"

	"SimpleTikTok/internal/MicroServices/mongodbmanage/internal/svc"
	"SimpleTikTok/internal/MicroServices/mongodbmanage/pkg/MongodbManageServer"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMinioLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMinioLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMinioLogic {
	return &GetMinioLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMinioLogic) GetMinio(in *MongodbManageServer.IdRequest) (*MongodbManageServer.MinioResponse, error) {
	// todo: add your logic here and delete this line

	return &MongodbManageServer.MinioResponse{}, nil
}
