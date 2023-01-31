package BaseInterface

import (
	"context"

	"SimpleTikTok/external_api/baseinterface/internal/svc"
	"SimpleTikTok/external_api/baseinterface/internal/types"
	"SimpleTikTok/oprations/mysqlconnect"
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

	db := mysqlconnect.GormDB
	// if err != nil {
	// 	logx.Errorf("SqlConnect err: %v", err)
	// 	return &types.UserRegisterHandlerResponse{
	// 		StatusCode: 400,
	// 		StatusMsg:  "用户已存在，请直接登录",
	// 		UserID:     -1,
	// 		Token:      "",
	// 	}, err
	// }
	res, err := mysqlconnect.FindUserIsExist(db, req.UserName, req.PassWord)
	if err != nil {
		logx.Errorf("UserRegisterLogic FindUserIsExist err: %v", err)
		return &types.UserRegisterHandlerResponse{
			StatusCode: 400,
			StatusMsg:  "查找注册用户是否存在失败",
			UserID:     -1,
			Token:      "",
		}, err
	}
	if res != 0 {
		logx.Infof("UserRegisterLogic Find User is exist err: %v", err)
		return &types.UserRegisterHandlerResponse{
			StatusCode: 400,
			StatusMsg:  "用户已存在，请直接登录",
			UserID:     int64(res),
			Token:      "",
		}, err
	}

	uid, err := mysqlconnect.CreateUser(db, req.UserName, req.PassWord)
	logx.Infof("%d", uid)
	if uid == -1 {
		return &types.UserRegisterHandlerResponse{
			StatusCode: -1,
			StatusMsg:  "注册失败",
			UserID:     -1,
			Token:      "",
		}, err
	}

	// yzx:go-zero 自带的jwt鉴权，有问题
	// payloads := make(map[string]any)
	// payloads["userIdentity"] = uid
	// TokenString, tokenErr := l.GetToken(time.Now().Unix(), l.svcCtx.Config.Auth.AccessSecret, payloads, l.svcCtx.Config.Auth.AccessExpire)
	// if tokenErr != nil {
	// 	return nil, tokenErr
	// }

	TokenString, err := tools.CreateToken(uid)
	return &types.UserRegisterHandlerResponse{
		StatusCode: 0,
		StatusMsg:  "注册成功",
		UserID:     int64(uid),
		Token:      TokenString,
	}, err
}
