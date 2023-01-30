package logic

import (
	"context"

	"SimpleTikTok/internal_proto/microservices/utilserver/internal/svc"
	"SimpleTikTok/internal_proto/microservices/utilserver/types/utilserver"

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

func (l *GetMinioLogic) GetMinio(in *utilserver.IdRequest) (*utilserver.MinioResponse, error) {
	// todo: add your logic here and delete this line

	return &utilserver.MinioResponse{}, nil
}
