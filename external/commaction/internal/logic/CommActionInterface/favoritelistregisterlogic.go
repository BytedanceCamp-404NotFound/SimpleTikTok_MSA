package CommActionInterface

import (
	"context"

	"SimpleTikTok/external/commaction/internal/svc"
	"SimpleTikTok/external/commaction/internal/types"
	"SimpleTikTok/oprations/commonerror"
	"SimpleTikTok/oprations/mysqlconnect"
	tools "SimpleTikTok/tools/token"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteListRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListRegisterLogic {
	return &FavoriteListRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteListRegisterLogic) FavoriteListRegister(req *types.FavoriteListRegisterHandlerRequest) (resp *types.FavoriteListRegisterHandlerResponse, err error) {
	ok, id, err := tools.CheckToke(req.Token)
	if !ok {
		logx.Infof("[pkg]logic [func]FavoriteListRegister [msg]req.Token is wrong ")
		return &types.FavoriteListRegisterHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_PARSE_TOKEN_ERROR),
			StatusMsg:  "登录过期，请重新登陆",
			VideoList:  []types.Video{},
		}, nil
	}
	if err != nil {
		logx.Errorf("[pkg]logic [func]FavoriteListRegister [msg]func CheckToken [err]%v", err)
		return &types.FavoriteListRegisterHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_INTERNAL_ERROR),
			StatusMsg:  "Token校验出错",
			VideoList:  []types.Video{},
		}, nil
	}

	n, err := mysqlconnect.FavoriteVideoNum(req.UserID)
	if err != nil {
		logx.Errorf("[pkg]logic [func]FavoriteListRegister [msg]func FavoriteVideoNum [err]%v", err)
		return &types.FavoriteListRegisterHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_DB_ERROR),
			StatusMsg:  "操作失败",
			VideoList:  []types.Video{},
		}, nil
	}
	if n == -1 {
		logx.Infof("[pkg]logic [func]FavoriteListRegister [msg]User does not exit ")
		return &types.FavoriteListRegisterHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_DB_ERROR),
			StatusMsg:  "获取喜欢列表的用户信息失败",
			VideoList:  []types.Video{},
		}, nil
	}

	v, err := mysqlconnect.GetFavoriteVideoList(req.UserID)
	if err != nil {
		logx.Errorf("[pkg]logic [func]FavoriteListRegister [msg]func GetFavoriteVideoList [err]%v", err)
		return &types.FavoriteListRegisterHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_DB_ERROR),
			StatusMsg:  "操作失败",
			VideoList:  []types.Video{},
		}, nil
	}

	videolist := make([]types.Video, n)
	for i := 0; i < int(n); i++ {
		user, ok := mysqlconnect.CheckUserInf(int(v[i].AuthorID), id)
		if !ok {
			logx.Infof("[pkg]logic [func]FavoriteListRegister [msg]User does not exist")
			return &types.FavoriteListRegisterHandlerResponse{
				StatusCode: int32(commonerror.CommonErr_INTERNAL_ERROR),
				StatusMsg:  "获取视频对应的用户信息失败",
				VideoList:  []types.Video{},
			}, nil
		}
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
	return &types.FavoriteListRegisterHandlerResponse{
		StatusCode: 0,
		StatusMsg:  "查询喜欢列表成功",
		VideoList:  videolist,
	}, err
}
