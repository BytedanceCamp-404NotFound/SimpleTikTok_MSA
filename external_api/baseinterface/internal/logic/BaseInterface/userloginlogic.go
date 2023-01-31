package BaseInterface

import (
	"context"

	"SimpleTikTok/external_api/baseinterface/internal/svc"
	"SimpleTikTok/external_api/baseinterface/internal/types"
	"SimpleTikTok/internal_proto/microservices/mysqlmanage/types/mysqlmanageserver"
	"SimpleTikTok/oprations/commonerror"
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
	uid, err := l.svcCtx.MySQLManageRpc.UserLogin(l.ctx, &mysqlmanageserver.UserLoginRequest{
		Username: req.UserName,
		Password: req.PassWord,
	})
	if err != nil {
		logx.Error("Check user rpc err: %v", err)
		return &types.UserloginHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_INTERNAL_ERROR),
			StatusMsg:  "服务器出错，等待修复",
		}, nil
	}

	logx.Infof("UserloginLogic CheckUser,uid:%v", uid)
	if uid.UserId == -1 {
		return &types.UserloginHandlerResponse{
			StatusCode: -1,
			StatusMsg:  "用户名或密码错误，请重试",
			UserID:     -1,
			Token:      "",
		}, err
	}
	TokenString, err := tools.CreateToken(int(uid.UserId))
	return &types.UserloginHandlerResponse{
		StatusCode: 0,
		StatusMsg:  "登录成功",
		UserID:     int64(uid.UserId),
		Token:      TokenString,
	}, err
}
