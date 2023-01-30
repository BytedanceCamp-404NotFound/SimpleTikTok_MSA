package logic

import (
	"context"

	"SimpleTikTok/internal/MicroServices/minio/internal/svc"
	"SimpleTikTok/internal/MicroServices/pkg/Minio"

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

func (l *GetMinioLogic) GetMinio(in *Minio.IdRequest) (*Minio.MinioResponse, error) {
	// todo: add your logic here and delete this line

	return &Minio.MinioResponse{}, nil
}
