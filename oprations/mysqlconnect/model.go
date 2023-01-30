package mysqlconnect

import (
	"time"
)

// 默认情况下，名为 `ID` 的字段会作为表的主键
type VideoInfo struct {
	VideoID       int64  `gorm:"cloumn:video_id;primaryKey"`
	AuthorID      int64
	PlayUrl       string // 视频播放地址
	CoverUrl      string // 视频封面地址
	FavoriteCount int64  // 视频的点赞总数
	CommentCount  int64  // 视频的评论总数
	IsFavotite    bool   // true-已点赞
	VideoTitle    string // 视频标题
}

type Favorite_list struct {
	Favorite_video_id int64
	Favorite_user_id  int
	Record_time       time.Time
}

type User_login struct {
	UserID       int64     `gorm:"cloumn:user_id;primaryKey"`
	UserName     string    `gorm:"cloumn:user_id;"`
	UserPwd      string    `gorm:"cloumn:user_id;"`
	RegisterDate time.Time `gorm:"cloumn:register_date;"`
}
