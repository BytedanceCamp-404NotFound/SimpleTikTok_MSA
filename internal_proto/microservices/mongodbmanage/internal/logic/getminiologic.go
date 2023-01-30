package logic

import (
	"context"

	"SimpleTikTok/internal_proto/microservices/mongodbmanage/internal/svc"
	"SimpleTikTok/internal_proto/microservices/mongodbmanage/types/mongodbmanageserver"

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

func (l *GetMinioLogic) GetMinio(in *mongodbmanageserver.IdRequest) (*mongodbmanageserver.MinioResponse, error) {
	// todo: add your logic here and delete this line

	return &mongodbmanageserver.MinioResponse{}, nil
}
