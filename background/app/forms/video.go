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
	Type     int    `form:"type"`
	Category string `form:"category"`
	Order    string `form:"order"`
	Letter   string `form:"letter"`
	Year     int    `form:"year"`
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
	Actors     []string   `json:"actors"`
	Categories []string   `json:"categories"`
	Cover      string     `json:"cover"`
	Date       string     `json:"date"`
	Lang       string     `json:"lang"`
	Master     string     `json:"master"`
	Playlist   []Playlist `json:"playlist"`
	Rank       string     `json:"rank"`
	Region     string     `json:"region"`
	Season     string     `json:"season"`
	Title      string     `json:"title"`
}

type Playlist struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Link  string `json:"link"`
}

type VideoConfig struct {
	FiltersConfig []FilterConfig `json:"filtersConfig"`
}

type FilterConfig struct {
	Name       string      `json:"name"`
	ID         uint        `json:"id"`
	Categories []Categorie `json:"categories"`
}

type Categorie struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type VideoIndex struct {
	Banners       []VideoBanner       `json:"banner"`
	ChineseComics []VideoChineseComic `json:"chineseComic"`
	Hots          []VideoHot          `json:"hots"`
	Japancomic    []VideoJapanComic   `json:"japanComic"`
	Latest        []VideoLatest       `json:"latest"`
	Perweek       []VideoPerweek      `json:"perweek"`
	TheatreComic  []VideoTheatreComic `json:"theatreComic"`
}

type VideoBanner struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Cover string `json:"cover"`
}

type VideoChineseComic struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Cover  string `json:"cover"`
	Season string `json:"season"`
}

type VideoHot struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Cover       string `json:"cover"`
	Season      string `json:"season"`
	Date        string `json:"date"`
	Description string `json:"description"`
}

type VideoJapanComic struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Cover  string `json:"cover"`
	Season string `json:"season"`
}

type VideoLatest struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Cover  string `json:"cover"`
	Season string `json:"season"`
}

type VideoPerweek struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Season string `json:"season"`
}

type VideoTheatreComic struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Cover  string `json:"cover"`
	Season string `json:"season"`
}
