package mysqlconnect

import "github.com/zeromicro/go-zero/core/logx"

func VideoNum(AuthorID int64) (n int64, err error) {
	db := GormDB
	db.Table("video_info").Where("author_id = ?", AuthorID).Count(&n)
	return n,err
}

//获得用户的发布列表
func GetVideoList(AuthorID int64, UserID int64) (list []VideoInfo, err error) {
	db := GormDB
	n, err:= VideoNum(AuthorID)
	if err != nil {
		logx.Errorf("[pkg]logic [func]GetVideoList [msg]func VideoNum [err]%v", err)
		return []VideoInfo{},nil
	}
	if n == 0 {
		logx.Infof("[pkg]logic [func]GetVideoList [msg]VideoList is empty")
		return []VideoInfo{},nil
	}

	err = db.Table("video_info").Where("author_id = ?", AuthorID).Find(&list).Error
	if err != nil {
		logx.Errorf("[pkg]logic [func]GetVideoList [msg]gorm Find [err]%v", err)
	}
	for i := 0; i < int(n); i++ {
		list[i].IsFavotite, err = IsFavotite(int64(UserID), list[i].VideoID)
	}
	return list,err
}