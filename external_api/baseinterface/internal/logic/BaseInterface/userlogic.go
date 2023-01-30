package BaseInterface

import (
	"context"

	"SimpleTikTok/external_api/baseinterface/internal/svc"
	"SimpleTikTok/external_api/baseinterface/internal/types"
	"SimpleTikTok/oprations/mysqlconnect"
	tools "SimpleTikTok/tools/token"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.UserHandlerRequest) (resp *types.UserHandlerResponse, err error) {
	ok, id, err := tools.CheckToke(req.Token)
	if !ok {
		return &types.UserHandlerResponse{
			StatusCode: -1,
			StatusMsg:  "请登录！",
			User:       types.User{},
		}, err
	}
	ui, ok := mysqlconnect.CheckUserInf(int(req.UserID), id)
	if !ok {
		return &types.UserHandlerResponse{
			StatusCode: -1,
			StatusMsg:  "查询的用户不存在！",
			User:       types.User{},
		}, err
	}
	return &types.UserHandlerResponse{
		StatusCode: 0,
		StatusMsg:  "查询成功！",
		User: types.User{
			UserId:        ui.User.UserID,
			Name:          ui.User.UserNickName,
			FollowCount:   ui.User.FollowCount,
			FollowerCount: ui.User.FollowerCount,
			IsFollow:      ui.IsFollow,
		},
	}, err
}
