package logic

import (
	"context"

	"SimpleTikTok/internal_proto/microservices/miniomanage/internal/svc"
	"SimpleTikTok/internal_proto/microservices/miniomanage/types/miniomanageserver"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutFileUploaderByteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutFileUploaderByteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutFileUploaderByteLogic {
	return &PutFileUploaderByteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// byte形式文件上传
func (l *PutFileUploaderByteLogic) PutFileUploaderByte(in *miniomanageserver.PutFileUploaderByteRequest) (*miniomanageserver.PutFileUploaderByteponse, error) {
	// todo: add your logic here and delete this line

	return &miniomanageserver.PutFileUploaderByteponse{}, nil
}
