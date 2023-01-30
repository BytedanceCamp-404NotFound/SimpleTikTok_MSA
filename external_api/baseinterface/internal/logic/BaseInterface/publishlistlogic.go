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

type PublishListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishListLogic) PublishList(req *types.PublishListHandlerRequest) (resp *types.PublishListHandlerResponse, err error) {
	ok, id, err := tools.CheckToke(req.Token)
	if !ok {
		logx.Infof("[pkg]logic [func]PublishList [msg]req.Token is wrong ")
		return &types.PublishListHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_PARSE_TOKEN_ERROR),
			StatusMsg:  "登录过期，请重新登陆",
			VideoList:  []types.Video{},
		}, nil
	}
	if err != nil {
		logx.Errorf("[pkg]logic [func]PublishListr [msg]func CheckToken [err]%v", err)
		return &types.PublishListHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_INTERNAL_ERROR),
			StatusMsg:  "Token校验出错",
			VideoList:  []types.Video{},
		}, nil
	}

	user, ok := mysqlconnect.CheckUserInf(int(req.UserID), id)
	if !ok {
		logx.Infof("[pkg]logic [func]PublishList [msg]User does not exist")
		return &types.PublishListHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_INTERNAL_ERROR),
			StatusMsg:  "获取用户信息失败",
			VideoList:  []types.Video{},
		}, nil
	}

	n, err := mysqlconnect.VideoNum(req.UserID)
	if err != nil {
		logx.Errorf("[pkg]logic [func]PublishList [msg]func VideoNum [err]%v", err)
		return &types.PublishListHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_DB_ERROR),
			StatusMsg:  "操作失败",
			VideoList:  []types.Video{},
		}, nil
	}

	v, err := mysqlconnect.GetVideoList(req.UserID, int64(id))
	if err != nil {
		logx.Errorf("[pkg]logic [func]PublishList [msg]func GetVideoList [err]%v", err)
		return &types.PublishListHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_DB_ERROR),
			StatusMsg:  "操作失败",
			VideoList:  []types.Video{},
		}, nil
	}

	videolist := make([]types.Video, n)
	for i := 0; i < int(n); i++ {
		videolist[i] = types.Video{
			Id: v[i].VideoID,
			Author: types.User{
				UserId:        user.User.UserID,
				Name:          user.User.UserNickName,
				FollowCount:   user.User.FollowCount,
				FollowerCount: user.User.FollowerCount,
				IsFollow:      user.IsFollow,
			},
			PlayUrl:       v[i].PlayUrl,
			CoverUrl:      v[i].CoverUrl,
			FavoriteCount: v[i].FavoriteCount,
			CommentCount:  v[i].CommentCount,
			IsFavotite:    v[i].IsFavotite,
			VideoTitle:    v[i].VideoTitle,
		}
	}
	return &types.PublishListHandlerResponse{
		StatusCode: 0,
		StatusMsg:  "查询发布列表成功",
		VideoList:  videolist,
	}, err
}
