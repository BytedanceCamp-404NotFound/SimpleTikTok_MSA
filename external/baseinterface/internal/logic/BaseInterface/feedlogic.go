package BaseInterface

import (
	"context"
	"time"

	"SimpleTikTok/external/baseinterface/internal/svc"
	"SimpleTikTok/external/baseinterface/internal/types"
	"SimpleTikTok/oprations/commonerror"
	"SimpleTikTok/oprations/mysqlconnect"
	tools "SimpleTikTok/tools/token"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedLogic) Feed(req *types.FeedHandlerRequest) (resp *types.FeedHandlerResponse, err error) {
	ok, userId, err := tools.CheckToke(req.Token)
	if err != nil {
		return &types.FeedHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_INTERNAL_ERROR),
			StatusMsg:  "Token校验出错",
			VideoList:  []types.Video{},
			NextTime:   time.Now().Unix(), // 暂时返回当前时间
		}, nil
	}
	if !ok {
		logx.Infof("[pkg]logic [func]Feed [msg]feedUserInfo.Name is nuil ")
		return &types.FeedHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_PARAMETER_FAILED),
			StatusMsg:  "登录过期，请重新登陆",
			VideoList:  []types.Video{},
			NextTime:   time.Now().Unix(), // 暂时返回当前时间
		}, nil
	}

	feedUserInfo, err := mysqlconnect.GetFeedUserInfo(userId)
	if err != nil {
		logx.Errorf("[pkg]logic [func]Feed [msg]gorm GetFeedUserInfo [err]%v", err)
		return &types.FeedHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_INTERNAL_ERROR),
			StatusMsg:  "获取用户信息失败",
			VideoList:  []types.Video{},
			NextTime:   time.Now().Unix(), // 暂时返回当前时间
		}, nil
	}
	if feedUserInfo.UserNickName == "" {
		logx.Infof("[pkg]logic [func]Feed [msg]feedUserInfo.Name is nuil ")
		return &types.FeedHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_PARAMETER_FAILED),
			StatusMsg:  "用户信息为空",
			VideoList:  []types.Video{},
			NextTime:   time.Now().Unix(), // 暂时返回当前时间
		}, nil
	}
	var respFeedUserInfo types.User
	respFeedUserInfo.UserId = feedUserInfo.UserID
	respFeedUserInfo.Name = feedUserInfo.UserNickName
	respFeedUserInfo.FollowCount = feedUserInfo.FollowCount
	respFeedUserInfo.FollowerCount = feedUserInfo.FollowerCount
	respFeedUserInfo.IsFollow = feedUserInfo.IsFollow

	var feedVideLists []mysqlconnect.VideoInfo
	feedVideLists, err = mysqlconnect.GetFeedVideoList(userId)
	if err != nil {
		logx.Errorf("[pkg]logic [func]Feed [msg]gorm GetFeedVideoList [err]%v", err)
		return &types.FeedHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_INTERNAL_ERROR),
			StatusMsg:  "获取视频信息失败",
			VideoList:  []types.Video{},
			NextTime:   time.Now().Unix(), // 暂时返回当前时间
		}, nil
	}
	if feedVideLists == nil {
		logx.Infof("[pkg]logic [func]Feed [msg]feedVideLists is nil", err)
		return &types.FeedHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_INTERNAL_ERROR),
			StatusMsg:  "此用户没有视频信息",
			VideoList:  []types.Video{},
			NextTime:   time.Now().Unix(), // 暂时返回当前时间
		}, nil
	}

	var respFeedVideoList = make([]types.Video, len(feedVideLists))
	for index, val := range feedVideLists {
		respFeedVideoList[index].Id = val.VideoID
		respFeedVideoList[index].Author = respFeedUserInfo
		respFeedVideoList[index].PlayUrl = val.PlayUrl
		respFeedVideoList[index].CoverUrl = val.CoverUrl
		respFeedVideoList[index].FavoriteCount = val.FavoriteCount
		respFeedVideoList[index].CommentCount = val.CommentCount
		respFeedVideoList[index].IsFavotite = val.IsFavotite
		respFeedVideoList[index].VideoTitle = val.VideoTitle
	}

	return &types.FeedHandlerResponse{
		StatusCode: 200,
		StatusMsg:  "feed video success",
		VideoList:  respFeedVideoList,
		NextTime:   time.Now().Unix(), // 暂时返回当前时间
	}, nil
}
