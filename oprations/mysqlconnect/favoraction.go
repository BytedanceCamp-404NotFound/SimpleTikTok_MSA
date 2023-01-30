package mysqlconnect

import (
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// 获得传入用户id的喜欢列表的视频数量
func FavoriteVideoNum(UserID int64) (n int64, err error) {
	db := GormDB
	var exist int64
	err = db.Table("user_info").Where("user_id = ?", UserID).Count(&exist).Error
	if err != nil {
		logx.Errorf("[pkg]mysqlconnect [func]FavoriteVideoNum [msg]gorm UserID.Count [err]%v", err)
		return n, nil
	}
	if exist == 0 {
		logx.Infof("[pkg]mysqlconnect [func]FavoriteVideoNum [msg]User does not exit")
		return -1, err
	}

	err = db.Table("favorite_list").Where("favorite_user_id = ?", UserID).Count(&n).Error
	if err != nil {
		logx.Errorf("[pkg]mysqlconnect [func]FavoriteVideoNum [msg]gorm avorite_user_id.Count [err]%v", err)
		return n, nil
	}
	return n, err
}

// 获得传入用户的喜欢列表，以数组返回
func GetFavoriteVideoList(UserID int64) (list []VideoInfo, err error) {
	db := GormDB
	n, _ := FavoriteVideoNum(UserID)
	if n == 0 {
		logx.Infof("[pkg]mysqlconnect [func]GetVideoList [msg]FavoriteVideoNum is 0")
		return []VideoInfo{}, err
	}

	var VideoIdList []int
	err = db.Table("favorite_list").Select("favorite_video_id").Where("favorite_user_id = ?", UserID).Find(&VideoIdList).Error
	if err != nil {
		logx.Errorf("[pkg]mysqlconnect [func]GetVideoList [msg]gorm VideoID.Find  [err]%v", err)
		return []VideoInfo{}, nil
	}
	for i := 0; i < len(VideoIdList); i++ {
		var vl VideoInfo
		err = db.Table("video_info").Where("video_id = ?", VideoIdList[i]).Take(&vl).Error
		if err != nil {
			logx.Errorf("[pkg]mysqlconnect [func]GetVideoList [msg]gorm VideoInfo.Find [err]%v", err)
			return []VideoInfo{}, nil
		}
		vl.IsFavotite, err = IsFavotite(UserID, vl.VideoID)
		if err != nil {
			logx.Errorf("[pkg]mysqlconnect [func]GetVideoList [msg]func IsFavotite [err]%v", err)
			return []VideoInfo{}, nil
		}

		list = append(list, vl)
	}
	return list, err
}

// 判断用户对某视频是否点赞
func IsFavotite(UserID int64, VideoID int64) (ok bool, err error) {
	db := GormDB
	var n int64
	err = db.Table("favorite_list").Where("favorite_video_id = ? and favorite_user_id = ?", VideoID, UserID).Count(&n).Error
	if err != nil {
		logx.Errorf("[pkg]mysqlconnect [func]IsFavotite [msg]gorm VideoID = UserID  [err]%v", err)
		return false, nil
	}
	return n > 0, err
}

// 对视频点赞
func AddVideoFavorite(UserID int64, VideoID int64) (ok int ,err error) {
	db := GormDB
	var n int64
	err = db.Table("favorite_list").Where("favorite_video_id = ? AND favorite_user_id = ?",VideoID,UserID).Count(&n).Error
	if err != nil {
		logx.Errorf("[pkg]mysqlconnect [func]AddVideoFavorite [msg]gorm favorite_list.Count [err]%v", err)
		return 0,nil
	}
	if int(n) > 0 {
		logx.Infof("[pkg]mysqlconnect [func]AddVideoFavorite [msg]favorite already")
		return -2,nil
	}

	err = db.Table("video_info").Where("video_id = ?", VideoID).Count(&n).Error
	if err != nil {
		logx.Errorf("[pkg]mysqlconnect [func]AddVideoFavorite [msg]gorm favorite_list.Count [err]%v", err)
		return 0,nil
	}
	if int(n) == 0 {
		logx.Infof("[pkg]mysqlconnect [func]AddVideoFavorite [msg]Video does no exist")
		return -1,nil
	}

	err = db.Table("video_info").Where("video_id = ?", VideoID).Update("favorite_count", gorm.Expr("favorite_count + 1")).Error
	if err != nil {
		logx.Errorf("[pkg]mysqlconnect [func]AddVideoFavorite [msg]gorm favorite_count.UPDATE [err]%v", err)
		return 0,nil
	}

	fl := Favorite_list{
		Favorite_video_id: VideoID,
		Favorite_user_id: int(UserID),
		Record_time: time.Now(),
	}

	err = db.Table("favorite_list").Omit("record_id").Create(&fl).Error
	if err != nil {
		logx.Errorf("[pkg]mysqlconnect [func]AddVideoFavorite [msg]gorm favorite_list.Create [err]%v", err)
		return 0,nil
	}
	return 1,err
}

//取消视频点赞
func SubVideoFavorite(UserID int64, VideoID int64) (ok int, err error) {
	db := GormDB
	var n int64
	err = db.Table("video_info").Where("video_id = ?", VideoID).Count(&n).Error
	if err != nil {
		logx.Errorf("[pkg]mysqlconnect [func]SubVideoFavorite [msg]VideoID.Find [err]%v", err)
		return 0,nil
	}
	if int(n) == 0 {
		logx.Infof("[pkg]mysqlconnect [func]AddVideoFavorite [msg]Video does no exist")
		return -2,nil
	}

	err = db.Table("favorite_list").Where("favorite_video_id = ? AND favorite_user_id = ?", VideoID, UserID).Count(&n).Error
	if err != nil {
		logx.Errorf("[pkg]mysqlconnect [func]SubVideoFavorite [msg]favorite_list.Count [err]%v", err)
		return 0,nil
	}
	if int(n) == 0 {
		logx.Infof("[pkg]mysqlconnect [func]]SubVideoFavorite [msg]remove favorite already")
		return -1,nil
	}

	err = db.Table("video_info").Where("video_id = ?", VideoID).Update("favorite_count", gorm.Expr("favorite_count - 1")).Error
	if err != nil {
		logx.Errorf("[pkg]mysqlconnect [func]SubVideoFavorite [msg]gorm favorite_count.UPDATE [err]%v", err)
		return 0,nil
	}

	err = db.Table("favorite_list").Where("favorite_video_id = ? AND favorite_user_id = ?", VideoID, UserID).Delete(Favorite_list{}).Error
	if err != nil {
		logx.Errorf("[pkg]mysqlconnect [func]SubVideoFavorite [msg]gorm favorite_list.delete [err]%v", err)
		return 0,nil
	}
	return 1,err
}
