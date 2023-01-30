package logic

import (
	"context"

	"SimpleTikTok/internal/MicroServices/mysql/internal/svc"
	"SimpleTikTok/internal/MicroServices/pkg/MySQL"

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

func (l *GetMySQLLogic) GetMySQL(in *MySQL.IdRequest) (*MySQL.MySQLResponse, error) {
	// todo: add your logic here and delete this line

	return &MySQL.MySQLResponse{}, nil
}
