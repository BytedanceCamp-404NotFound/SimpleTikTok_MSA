package BaseInterface

import (
	"context"

	"SimpleTikTok/external_api/baseinterface/internal/svc"
	"SimpleTikTok/external_api/baseinterface/internal/types"
	"SimpleTikTok/oprations/commonerror"
	"SimpleTikTok/oprations/mysqlconnect"
	tools "SimpleTikTok/tools/token"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserloginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserloginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserloginLogic {
	return &UserloginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserloginLogic) Userlogin(req *types.UserloginHandlerRequest) (resp *types.UserloginHandlerResponse, err error) {
	uid, err := mysqlconnect.CheckUser(req.UserName, req.PassWord)
	if err != nil {
		logx.Error("Check user err: %v", err)
		return &types.UserloginHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_INTERNAL_ERROR),
			StatusMsg:  "服务器出错，等待修复",
		}, nil
	}
	logx.Infof("UserloginLogic CheckUser,uid:%v", uid)
	if uid == -1 {
		return &types.UserloginHandlerResponse{
			StatusCode: -1,
			StatusMsg:  "用户名或密码错误，请重试",
			UserID:     -1,
			Token:      "",
		}, err
	}
	TokenString, err := tools.CreateToken(uid)
	return &types.UserloginHandlerResponse{
		StatusCode: 0,
		StatusMsg:  "登录成功",
		UserID:     int64(uid),
		Token:      TokenString,
	}, err
}
