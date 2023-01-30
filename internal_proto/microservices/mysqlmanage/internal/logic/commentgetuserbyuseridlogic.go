package logic

import (
	"context"

	"SimpleTikTok/internal_proto/microservices/mysqlmanage/internal/svc"
	"SimpleTikTok/internal_proto/microservices/mysqlmanage/types/MySQLManageServer"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentGetUserByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentGetUserByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentGetUserByUserIdLogic {
	return &CommentGetUserByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 1
func (l *CommentGetUserByUserIdLogic) CommentGetUserByUserId(in *MySQLManageServer.CommentGetUserByUserIdRequest) (*MySQLManageServer.CommentGetUserByUserIdResponse, error) {
	// todo: add your logic here and delete this line

	return &MySQLManageServer.CommentGetUserByUserIdResponse{}, nil
}
