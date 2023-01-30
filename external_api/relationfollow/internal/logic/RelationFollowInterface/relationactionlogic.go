package RelationFollowInterface

import (
	"context"
	"time"

	"SimpleTikTok/external_api/relationfollow/internal/svc"
	"SimpleTikTok/external_api/relationfollow/internal/types"
	"SimpleTikTok/oprations/commonerror"
	"SimpleTikTok/oprations/mysqlconnect"
	tools "SimpleTikTok/tools/token"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationActionLogic {
	return &RelationActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type Follow_and_follower_list struct {
	UserID     int
	Token      string
	RecordTime time.Time
}

func (l *RelationActionLogic) RelationAction(req *types.RelationActionHandlerRequest) (resp *types.RelationActionHandlerResponse, err error) {
	ok, id, err := tools.CheckToke(req.Token)
	resultJson := types.RelationActionHandlerResponse{}

	if !ok {
		logx.Infof("[pkg]logic [func]PublishList [msg]req.Token is wrong ")
		return &types.RelationActionHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_PARSE_TOKEN_ERROR),
			StatusMsg:  "登录过期，请重新登陆",
		}, nil
	}
	if err != nil {
		logx.Errorf("[pkg]logic [func]PublishListr [msg]func CheckToken [err]%v", err)
		return &types.RelationActionHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_INTERNAL_ERROR),
			StatusMsg:  "Token校验出错",
		}, nil
	}

	// ！！！胡海龙：我先将sql代码抽离出来，有没有bug暂时没测试，过两天考试驾驶证再调试！！！

	ok, err2 := mysqlconnect.RelationAction(int64(id), req.To_user_id, int8(req.Sction_type))
	if ok {
		resultJson.StatusCode = 0
		resultJson.StatusMsg = "success"
	} else {
		resultJson.StatusCode = 500
		resultJson.StatusMsg = err.Error()
	}
	return &resultJson, err2
}
