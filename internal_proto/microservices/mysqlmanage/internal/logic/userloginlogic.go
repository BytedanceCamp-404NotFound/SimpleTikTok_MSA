package logic

import (
	"context"

	"SimpleTikTok/internal_proto/microservices/mysqlmanage/internal/svc"
	"SimpleTikTok/internal_proto/microservices/mysqlmanage/types/mysqlmanageserver"
	"SimpleTikTok/oprations/mysqlconnect"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户登陆校验
func (l *UserLoginLogic) UserLogin(in *mysqlmanageserver.UserLoginRequest) (*mysqlmanageserver.UserLoginResponse, error) {
	uid, err := mysqlconnect.CheckUser(in.Username, in.Password)
	if err != nil {
		logx.Error("Check user rpc error: %v", err)
		return &mysqlmanageserver.UserLoginResponse{
			UserId: -1,
		}, err
	}

	return &mysqlmanageserver.UserLoginResponse{
		UserId: int32(uid),
	}, nil
}
