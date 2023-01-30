package mysqlconnect

import "github.com/zeromicro/go-zero/core/logx"

type PublishActionVideoInfo struct {
	Video_id       int32 
	Author_id      int64
	Play_url       string // 视频播放地址
	Cover_url      string // 视频封面地址
	Favorite_count int64  // 视频的点赞总数
	Comment_count  int64  // 视频的评论总数
	Video_title    string // 视频标题
}

func CreatePublishActionViedeInfo(tmpvideoInfo *PublishActionVideoInfo) error {
	db := GormDB
	err := db.Table("video_info").Create(&tmpvideoInfo).Error
	if err != nil {
		logx.Errorf("[pkg]mysqlconnect [func]CreatePublishActionViedeInfo [msg]gorm [err]%v", err)
		return err
	}
	return nil
}
