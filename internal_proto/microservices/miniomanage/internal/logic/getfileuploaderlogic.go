package logic

import (
	"context"

	"SimpleTikTok/internal_proto/microservices/miniomanage/internal/svc"
	"SimpleTikTok/internal_proto/microservices/miniomanage/pkg/MinioManageServer"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileUploaderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFileUploaderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileUploaderLogic {
	return &GetFileUploaderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 文件下载
func (l *GetFileUploaderLogic) GetFileUploader(in *MinioManageServer.GetMinioConnectRequest) (*MinioManageServer.GetMinioConnectResponse, error) {
	// todo: add your logic here and delete this line

	return &MinioManageServer.GetMinioConnectResponse{}, nil
}
