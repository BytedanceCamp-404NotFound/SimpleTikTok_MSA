package RelationFollowInterface

import (
	"context"

	"SimpleTikTok/external_api/relationfollow/internal/svc"
	"SimpleTikTok/external_api/relationfollow/internal/types"
	"SimpleTikTok/oprations/commonerror"
	"SimpleTikTok/oprations/mysqlconnect"
	tools "SimpleTikTok/tools/token"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationFollowerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFollowerListLogic {
	return &RelationFollowerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationFollowerListLogic) RelationFollowerList(req *types.RelationFollowerListHandlerRequest) (resp *types.RelationFollowerListHandlerResponse, err error) {
	ok, id, err := tools.CheckToke(req.Token)
	if !ok {
		logx.Infof("[pkg]logic [func]PublishList [msg]req.Token is wrong ")
		return &types.RelationFollowerListHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_PARSE_TOKEN_ERROR),
			StatusMsg:  "登录过期，请重新登陆",
		}, nil
	}
	if err != nil {
		logx.Errorf("[pkg]logic [func]PublishListr [msg]func CheckToken [err]%v", err)
		return &types.RelationFollowerListHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_INTERNAL_ERROR),
			StatusMsg:  "Token校验出错",
		}, nil
	}
	resultJson := &types.RelationFollowerListHandlerResponse{}

	// ！！！胡海龙：我先将sql代码抽离出来，有没有bug暂时没测试，过两天考试驾驶证再调试！！！

	rflhr, err2 := mysqlconnect.GetFollowerList(int64(id), req.UserId)
	Userlist := make([]types.RelationUser, 0)
	for i := 0; i < len(rflhr); i++ {
		Userlist[i] = types.RelationUser{
			Id:            rflhr[i].Id,
			Name:          rflhr[i].Name,
			FollowCount:   rflhr[i].FollowCount,
			FollowerCount: rflhr[i].FollowerCount,
			IsFollow:      rflhr[i].IsFollow,
		}
		resultJson.StatusCode = 0
		resultJson.StatusMsg = "success"
	}
	resultJson.UserList = Userlist
	return resultJson, err2
}
