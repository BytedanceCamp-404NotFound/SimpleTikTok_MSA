package logic

import (
	"context"

	"SimpleTikTok/internal_proto/microservices/miniomanage/internal/svc"
	"SimpleTikTok/internal_proto/microservices/miniomanage/types/miniomanageserver"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutFileUploaderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutFileUploaderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutFileUploaderLogic {
	return &PutFileUploaderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 文件上传
func (l *PutFileUploaderLogic) PutFileUploader(in *miniomanageserver.PutFileUploaderRequest) (*miniomanageserver.PutFileUploaderResponse, error) {
	// todo: add your logic here and delete this line

	return &miniomanageserver.PutFileUploaderResponse{}, nil
}
