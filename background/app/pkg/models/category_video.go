package models

type CategoryVideo struct {
	CreatorBase
	CategoryID uint `json:"categoryId" gorm:"column:category_id;type:int unsigned;comment:分类id"`
	VideoID    uint `json:"videoId" gorm:"column:video_id;type:int unsigned;comment:视频id"`
}
