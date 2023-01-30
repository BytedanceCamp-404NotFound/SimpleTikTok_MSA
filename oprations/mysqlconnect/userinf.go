package mysqlconnect

import "github.com/zeromicro/go-zero/core/logx"

type User struct {
	UserID        int64  `gorm:"cloumn:user_id;primaryKey"`
	UserNickName  string `gorm:"cloumn:user_nick_name"`
	FollowCount   int64  `gorm:"cloumn:follow_count"`
	FollowerCount int64  `gorm:"cloumn:follower_count"`
}

type User_inf struct {
	User     User
	IsFollow bool
}

func CheckUserInf(UserID int, FollowerID int) (u User_inf, exist bool) {
	db := GormDB
	result := db.Table("user_info").Where("user_id = ?", UserID).Find(&u.User)
	if result.RowsAffected == 0 {
		logx.Infof("[pkg]mysqlconnect [func]CheckUserInf [msg]User does not exit")
		return u, false
	}
	var err error
	u.IsFollow, err = CheckIsFollow(UserID, FollowerID)
	if err != nil {
		logx.Errorf("[pkg]mysqlconnect [func]CheckUserInf [msg]func CheckIsFollow")
		return 
	}
	return u, true
}

func CheckIsFollow(UserID int, FollowerID int) (ok bool, err error) {
	var num int64
	db := GormDB
	db.Table("follow_and_follower_list").Where("user_id = ? and follower_id = ?", UserID, FollowerID).Count(&num)

	return num > 0, err
}

func CreateInfo(UserName string, uid int64) error {
	info := User{UserID: int64(uid), UserNickName: UserName, FollowCount: 0, FollowerCount: 0}
	db := GormDB
	err := db.Table("user_info").Create(&info).Error

	return err
}
