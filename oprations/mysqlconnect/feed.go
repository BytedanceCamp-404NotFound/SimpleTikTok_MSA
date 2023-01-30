package mysqlconnect

import (
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedUserInfo struct {
	UserID        int64  // 用户ID
	UserNickName  string // 用户名称
	FollowCount   int64  // 关注总数
	FollowerCount int64  // 粉丝总数
	IsFollow      bool   // true-已关注
	UpdateTime    time.Time
}

func GetFeedUserInfo(UserId int) (FeedUserInfo, error) {
	db := GormDB
	var tmpFeedUserInfo FeedUserInfo
	err := db.Table("user_info").Where("user_id = ?", UserId).First(&tmpFeedUserInfo).Error
	if err != nil {
		logx.Errorf("[pkg]mysqlconnect [func]FeedGetUserInfo [msg]gorm [err]%v", err)
		return FeedUserInfo{}, err
	}
	return tmpFeedUserInfo, err
}

func GetFeedVideoList(userId int) ([]VideoInfo, error) {
	db := GormDB
	var tmpFeedVideoList []VideoInfo
	err := db.Table("video_info").Where("author_id = ?", userId).Scan(&tmpFeedVideoList).Limit(10).Error
	if err != nil {
		logx.Errorf("[pkg]mysqlconnect [func]GetFeedVideoList [msg]gorm [err]%v", err)
		return []VideoInfo{}, err
	}
	return tmpFeedVideoList, nil
}
