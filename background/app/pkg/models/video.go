package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	VideoLanguageChinese = iota
	VideoLanguageEnglish
	VideoLanguageJapan
)

const (
	VideoRegionChina = iota
	VideoRegionAmerica
	VideoRegionJapan
)

const (
	VideoSeasonUpdate = iota
	VideoSeasonFinish
)

type Video struct {
	CreatorBase
	Title     string    `json:"title" gorm:"column:title;type:varchar(100);comment: 视频标题"`
	Introduce string    `json:"introduce" gorm:"column:introduce;type:varchar(500);comment: 视频介绍"`
	Cover     string    `json:"cover" gorm:"column:cover;type:varchar(500);comment: 视频封面"`
	Language  int       `json:"language" gorm:"column:language;type:tinyint unsigned;comment: 语言 0-简体中文 1-英文 2-日文"`
	Master    string    `json:"master" gorm:"column:master;type:varchar(500);comment: 作者"`
	FirstDate time.Time `json:"firstDate" gorm:"column:first_date;type:datetime;comment: 首播时间"`
	Ranking   int       `json:"ranking" gorm:"column:ranking;type:int;comment: 排名"`
	Score     float64   `json:"score" gorm:"column:score;type:float;comment: 评分"`
	Region    int       `json:"region" gorm:"column:region;type:tinyint unsigned;comment: 地区 0-中国 1-美国 2-日本"`
	Season    int       `json:"season" gorm:"column:season;type:tinyint unsigned;comment: 动漫状态 0-更新 1-完结"`
	Actors    string    `json:"actors" gorm:"column:actors;type:varchar(100);comment: 声优名称列表"`
}

func (v *Video) TableName() string {
	return "videos"
}

func (t *Video) Insert(db *gorm.DB) error {
	return db.Model(&t).Create(&t).Error
}

func (t *Video) Delete(db *gorm.DB, conditions interface{}, args ...interface{}) error {
	return db.Model(&t).Where(conditions, args...).Error
}

func (t *Video) Updates(db *gorm.DB, value interface{}, conditions interface{}, args ...interface{}) error {
	return db.Model(&t).Where(conditions, args...).Updates(value).Error
}

func (t *Video) WhereOne(db *gorm.DB, query interface{}, args ...interface{}) (video *Video, err error) {
	err = db.Model(&t).Where(query, args...).First(&video).Error
	if err != nil {
		return nil, err
	}
	return video, nil
}

func (t *Video) WhereAll(db *gorm.DB, query interface{}, args ...interface{}) (videos []*Video, err error) {
	err = db.Model(&t).Where(query, args...).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (t *Video) PageListOrder(db *gorm.DB, order string, params *ListPageInput, conditions interface{}, args ...interface{}) (videos []*Video, total int64, err error) {
	if order == "" {
		order = "created_at DESC"
	}
	offset := (params.Page - 1) * params.Size

	err = db.Model(&t).Where(conditions, args...).Offset(offset).Limit(params.Size).Order(order).Find(&videos).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Model(&t).Where(conditions, args...).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	return videos, total, nil
}
