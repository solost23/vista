package forms

import "vista/pkg/utils"

type VideoInsertForm struct {
	CategoryId uint   `form:"categoryId" binding:"required"`
	Title      string `form:"title" binding:"required,max=20"`
	Introduce  string `form:"introduce" binding:"required,max=20"`
	ImageUrl   string `form:"imageUrl" binding:"required" comment:"oss上传之后生成的url,前端再传过来"`
	VideoUrl   string `form:"videoUrl" binding:"required" comment:"oss上传之后生成的url,前端再传过来"`
}

type VideoListForm struct {
	utils.PageForm
	CategoryId   uint   `form:"categoryId"`
	CategoryName string `form:"categoryName"`
	UserName     string `form:"userName"`
	VideoTitle   string `form:"videoTitle"`
	Introduce    string `form:"introduce"`
}

type VideoListRecord struct {
	ID           uint   `json:"id"`
	UserID       uint   `json:"userId"`
	UserName     string `json:"userName"`
	CategoryID   uint   `json:"categoryId"`
	CategoryName string `json:"categoryName"`
	Title        string `json:"title"`
	Introduce    string `json:"introduce"`
	ImageUrl     string `json:"imageUrl"`
	VideoUrl     string `json:"videoUrl"`
	ThumbCount   int64  `json:"thumbCount"`
	CommentCount int64  `json:"commentCount"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

type VideoListResponse struct {
	Records  []VideoListRecord `json:"records"`
	PageList *utils.PageList
}

type VideoSearchForm struct {
	Name string `form:"name" comment:"视频名称"`
	utils.PageForm
}

type VideoSearch struct {
	utils.PageList
	Records []VideoSearchRecord `json:"results"`
}

type VideoSearchRecord struct {
	Categories  []string `json:"categories"`
	Cover       string   `json:"cover"`
	Date        string   `json:"date"`
	Description string   `json:"description"`
	ID          uint     `json:"id"`
	Season      string   `json:"season"`
	Title       string   `json:"title"`
}

type VideoFilterForm struct {
	utils.PageForm
	Type       int    `form:"type"`
	CategoryID uint   `form:"category"`
	Order      string `form:"order"`
	Letter     string `form:"letter"`
	Year       int    `form:"year"`
}

type VideoFilter struct {
	utils.PageList
	Records []VideoFilterRecord `json:"results"`
}

type VideoFilterRecord struct {
	Cover  string `json:"cover"`
	ID     uint   `json:"id"`
	Season string `json:"season"`
	Title  string `json:"title"`
}

type VideoDetail struct {
	Actors     []string              `json:"actors"`
	Categories []string              `json:"categories"`
	Cover      string                `json:"cover"`
	Date       string                `json:"date"`
	Lang       string                `json:"lang"`
	Master     string                `json:"master"`
	Playlist   map[string][]Playlist `json:"playlist"`
	Score      float64               `json:"score"`
	Rank       int                   `json:"rank"`
	Region     string                `json:"region"`
	Season     string                `json:"season"`
	Title      string                `json:"title"`
}

type Playlist struct {
	ID    uint     `json:"id"`
	Title string   `json:"key"`
	Link  []string `json:"value"`
	Sort  int      `json:"sort"`
}

type VideoConfig struct {
	FiltersConfig []FilterConfig `json:"filtersConfig"`
}

type FilterConfig struct {
	Name       string     `json:"name"`
	ID         uint       `json:"id"`
	Categories []Category `json:"categories"`
}

type Category struct {
	ID   uint   `json:"classid"`
	Name string `json:"classname"`
}

type VideoIndex struct {
	Banners       []VideoRecord `json:"banner"`
	ChineseComics []VideoRecord `json:"chineseComic"`
	Hots          []VideoRecord `json:"hots"`
	Japancomic    []VideoRecord `json:"japanComic"`
	Latest        []VideoRecord `json:"latest"`
	Perweek       []VideoRecord `json:"perweek"`
	TheatreComic  []VideoRecord `json:"theatreComic"`
}

type VideoRecord struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Cover       string `json:"cover"`
	Season      string `json:"season"`
	Date        string `json:"date"`
	Description string `json:"description"`
}
