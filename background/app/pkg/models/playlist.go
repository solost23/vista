package models

type Playlist struct {
	CreatorBase
	VideoID uint   `json:"videoId" gorm:"column:video_id;type:bigint unsigned;comment: 视频 ID"`
	Title   string `json:"title" gorm:"column:title;type:varchar(100);comment: 播放列表标题"`
	Link    string `json:"link" gorm:"column:link;type:varchar(500);comment: 播放列表链接"`
}
