package mysqlconnect

import "github.com/zeromicro/go-zero/core/logx"

type CommentUser struct {
	UserId        int64  `gorm:"column:user_id"        json:"id"               form:"user_id"        bson:"user_id"`
	Name          string `gorm:"column:user_nick_name" json:"name"             form:"name"           bson:"name"`
	FollowCount   int64  `gorm:"column:follow_count"   json:"follow_count"     form:"follow_count"   bson:"follow_count"`
	FollowerCount int64  `gorm:"column:follower_count" json:"follower_count"   form:"follower_count" bson:"follower_count"`
	IsFollow      bool   `json:"is_follow"             form:"is_follow"        bson:"is_follow"`
}

func CommentGetUserByUserId(userId int) (CommentUser, error) {
	var user CommentUser
	db := GormDB
	err := db.Table("user_info").Where("user_id=?", userId).First(&user).Error
	if err != nil {
		logx.Errorf("[pkg]mysqlconnect [func]CommentGetUserByUserId [msg]search user_info failed, [err]%v", err)
		return CommentUser{}, err
	}
	return user, err
}