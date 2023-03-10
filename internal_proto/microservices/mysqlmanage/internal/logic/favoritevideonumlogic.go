package logic

import (
	"context"

	"SimpleTikTok/internal_proto/microservices/mysqlmanage/internal/svc"
	"SimpleTikTok/internal_proto/microservices/mysqlmanage/types/mysqlmanageserver"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteVideoNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteVideoNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteVideoNumLogic {
	return &FavoriteVideoNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 2
func (l *FavoriteVideoNumLogic) FavoriteVideoNum(in *mysqlmanageserver.FavoriteVideoNumRequest) (*mysqlmanageserver.FavoriteVideoNumResponse, error) {
	// todo: add your logic here and delete this line

	return &mysqlmanageserver.FavoriteVideoNumResponse{}, nil
}
