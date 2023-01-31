package logic

import (
	"context"

	"SimpleTikTok/internal_proto/microservices/mysqlmanage/internal/svc"
	"SimpleTikTok/internal_proto/microservices/mysqlmanage/types/mysqlmanageserver"
	"SimpleTikTok/oprations/mysqlconnect"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRigsterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRigsterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRigsterLogic {
	return &UserRigsterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户注册
func (l *UserRigsterLogic) UserRigster(in *mysqlmanageserver.UserRegisterRequest) (*mysqlmanageserver.UserRegisterResponse, error) {
	db := mysqlconnect.GormDB
	res, err := mysqlconnect.FindUserIsExist(db, in.Username, in.Password)
	if err != nil {
		logx.Error("Rigster rpc error: %v", err)
		return &mysqlmanageserver.UserRegisterResponse{
			UserId: -2,
		}, err
	}
	if res != 0 {
		return &mysqlmanageserver.UserRegisterResponse{
			UserId: -3,
		}, nil
	}

	uid, err := mysqlconnect.CreateUser(db, in.Username, in.Password)
	if err != nil {
		logx.Error("Rigster rpc error: %v", err)
		return &mysqlmanageserver.UserRegisterResponse{
			UserId: int32(uid),
		}, err
	}

	return &mysqlmanageserver.UserRegisterResponse{
		UserId: int32(uid),
	}, nil
}
