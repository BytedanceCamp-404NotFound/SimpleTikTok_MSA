package BaseInterface

import (
	"context"

	"SimpleTikTok/external_api/baseinterface/internal/svc"
	"SimpleTikTok/external_api/baseinterface/internal/types"
	"SimpleTikTok/internal_proto/microservices/mysqlmanage/types/mysqlmanageserver"
	tools "SimpleTikTok/tools/token"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// TODO优化:
// 1. 大量用户注册测试
// 响应时间：2471.2ms
func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterHandlerRequest) (resp *types.UserRegisterHandlerResponse, err error) {
	logx.Infof("UserRegister UserName: %v PassWord: %v", req.UserName, req.PassWord == "")
	if req.PassWord == "" && req.UserName == "" {
		logx.Errorf("UserName and PassWord is nil %v", err)
		return &types.UserRegisterHandlerResponse{
			StatusCode: 400,
			StatusMsg:  "用户名或者密码错误，注册失败",
			UserID:     -1,
			Token:      "",
		}, err
	}

	uid, err := l.svcCtx.MySQLManageRpc.UserRigster(l.ctx, &mysqlmanageserver.UserRegisterRequest{
		Username: req.UserName,
		Password: req.PassWord,
	})
	if uid.UserId == -2 && err != nil {
		logx.Errorf("UserRegister rpc: FindUserIsExist err: %v", err)
		return &types.UserRegisterHandlerResponse{
			StatusCode: 400,
			StatusMsg:  "查找注册用户是否存在失败",
			UserID:     -1,
			Token:      "",
		}, err
	}
	if uid.UserId == -3 {
		logx.Infof("UserRegister rpc: Find User is exist err: %v", err)
		return &types.UserRegisterHandlerResponse{
			StatusCode: 400,
			StatusMsg:  "用户已存在，请直接登录",
			UserID:     int64(uid.UserId),
			Token:      "",
		}, err
	}

	logx.Infof("%d", uid.UserId)
	if uid.UserId == -1 && err != nil {
		return &types.UserRegisterHandlerResponse{
			StatusCode: -1,
			StatusMsg:  "注册失败",
			UserID:     -1,
			Token:      "",
		}, err
	}
	TokenString, err := tools.CreateToken(int(uid.UserId))
	return &types.UserRegisterHandlerResponse{
		StatusCode: 0,
		StatusMsg:  "注册成功",
		UserID:     int64(uid.UserId),
		Token:      TokenString,
	}, err
}
