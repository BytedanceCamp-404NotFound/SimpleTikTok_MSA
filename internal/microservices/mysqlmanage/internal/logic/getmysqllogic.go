package logic

import (
	"context"

	"SimpleTikTok/internal/microservices/mysqlmanage/internal/svc"
	"SimpleTikTok/internal/microservices/mysqlmanage/pkg/MySQLManageServer"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMySQLLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMySQLLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMySQLLogic {
	return &GetMySQLLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMySQLLogic) GetMySQL(in *MySQLManageServer.IdRequest) (*MySQLManageServer.MySQLResponse, error) {
	// todo: add your logic here and delete this line

	return &MySQLManageServer.MySQLResponse{}, nil
}
