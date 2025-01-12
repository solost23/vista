package services

import (
	"errors"
	"math/rand"
	"mime/multipart"
	"strings"
	"time"
	"vista/forms"
	"vista/global"
	"vista/pkg/constants"
	"vista/pkg/models"
	"vista/pkg/response"
	"vista/pkg/utils"
	"vista/services/servants"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/now"
	"gorm.io/gorm"
)

// func (s *Service) VideoList(c *gin.Context, params *forms.VideoListForm) (response *forms.VideoListResponse, err error) {
// 	db := global.DB

// 	categoryIds := make([]uint, 0)
// 	if params.CategoryId > 0 {
// 		categoryIds = append(categoryIds, params.CategoryId)
// 	}
// 	if params.CategoryName != "" {
// 		query := []string{"title LIKE ?"}
// 		args := []interface{}{models.LikeFilter(params.CategoryName)}
// 		categories, err := (&models.Category{}).WhereAll(db, strings.Join(query, " AND "), args...)
// 		if err != nil {
// 			return nil, err
// 		}
// 		for _, category := range categories {
// 			categoryIds = append(categoryIds, category.ID)
// 		}
// 	}
// 	userIds := make([]uint, 0)
// 	if params.UserName != "" {
// 		query := []string{"user_name LIKE ?"}
// 		args := []interface{}{models.LikeFilter(params.UserName)}
// 		users, err := (&models.User{}).WhereAll(db, strings.Join(query, " AND "), args...)
// 		if err != nil {
// 			return nil, err
// 		}
// 		for _, user := range users {
// 			userIds = append(userIds, user.ID)
// 		}
// 	}

// 	query := make([]string, 0, 4)
// 	args := make([]interface{}, 0, 4)
// 	if len(categoryIds) > 0 {
// 		query = append(query, "category_id IN ?")
// 		args = append(args, categoryIds)
// 	}
// 	if len(userIds) > 0 {
// 		query = append(query, "creator_id IN ?")
// 		args = append(args, userIds)
// 	}
// 	if params.VideoTitle != "" {
// 		query = append(query, "title LIKE ?")
// 		args = append(args, models.LikeFilter(params.VideoTitle))
// 	}
// 	if params.Introduce != "" {
// 		query = append(query, "introduce LIKE ?")
// 		args = append(args, models.LikeFilter(params.Introduce))
// 	}
// 	videos, total, err := (&models.Video{}).PageListOrder(db, "", &models.ListPageInput{Page: params.Page, Size: params.Size}, strings.Join(query, " AND "), args...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	videoIds := make([]uint, 0, len(videos))
// 	categoryIds = make([]uint, 0, len(videos))
// 	userIds = make([]uint, 0, len(videos))
// 	for _, video := range videos {
// 		videoIds = append(videoIds, video.ID)
// 		categoryIds = append(categoryIds, video.CategoryId)
// 		userIds = append(userIds, video.UserId)
// 	}
// 	// 求出所有所属人和所属类别
// 	userNameMaps := make(map[uint]string)
// 	categoryNameMaps := make(map[uint]string)
// 	query = []string{"id IN ?"}
// 	args = []interface{}{userIds}
// 	users, err := (&models.User{}).WhereAll(db, strings.Join(query, " AND "), args...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	query = []string{"id IN ?"}
// 	args = []interface{}{categoryIds}
// 	categories, err := (&models.Category{}).WhereAll(db, strings.Join(query, " AND "), args...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if err != nil {
// 		return nil, err
// 	}
// 	for _, user := range users {
// 		userNameMaps[user.ID] = user.Username
// 	}
// 	for _, category := range categories {
// 		categoryNameMaps[category.ID] = category.Title
// 	}
// 	// 求出视频点赞数和评论数
// 	thumbCountMap := make(map[uint]uint)
// 	commentCountMap := make(map[uint]uint)
// 	query = []string{"video_id IN ?", "is_thumb = ?"}
// 	args = []interface{}{videoIds, "ISTHUMB"}
// 	thumbs, err := (&models.Comment{}).WhereCountGroup(db, "", strings.Join(query, " AND "))
// 	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
// 		return nil, err
// 	}
// 	query = []string{"video_id IN ?", "is_thumb = ?"}
// 	args = []interface{}{videoIds, "ISCOMMENT"}
// 	comments, err := (&models.Comment{}).WhereCountGroup(db, "", strings.Join(query, " AND "))
// 	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
// 		return nil, err
// 	}
// 	for _, thumb := range thumbs {
// 		thumbCountMap[thumb.VideoId] = thumb.CommentCount
// 	}
// 	for _, comment := range comments {
// 		commentCountMap[comment.VideoId] = comment.CommentCount
// 	}
// 	// 封装数据返回
// 	records := make([]forms.VideoListRecord, 0, len(videos))
// 	for _, video := range videos {
// 		records = append(records, forms.VideoListRecord{
// 			ID:           video.ID,
// 			UserID:       video.UserId,
// 			UserName:     userNameMaps[video.UserId],
// 			CategoryID:   video.CategoryId,
// 			CategoryName: categoryNameMaps[video.CategoryId],
// 			Title:        video.Title,
// 			Introduce:    video.Introduce,
// 			ImageUrl:     utils.FulfillImageOSSPrefix(video.ImageUrl),
// 			VideoUrl:     utils.FulfillImageOSSPrefix(video.VideoUrl),
// 			ThumbCount:   int64(thumbCountMap[video.ID]),
// 			CommentCount: int64(commentCountMap[video.ID]),
// 			CreatedAt:    video.CreatedAt.Format(constants.DateTime),
// 			UpdatedAt:    video.UpdatedAt.Format(constants.DateTime),
// 		})
// 	}
// 	response = &forms.VideoListResponse{
// 		Records: records,
// 		PageList: &utils.PageList{
// 			Size:    params.Size,
// 			Pages:   int64(math.Ceil(float64(total) / float64(params.Size))),
// 			Total:   total,
// 			Current: params.Page,
// 		},
// 	}
// 	return response, nil
// }

// func (s *Service) SearchVideo(c *gin.Context, params *forms.SearchForm) (*forms.VideoListResponse, error) {
// 	db := global.DB

// 	keyword := "*"
// 	if params.Keyword != nil && *params.Keyword != "" {
// 		keyword = *params.Keyword
// 	}
// 	z := NewZinc()
// 	from := int32((params.Page - 1) * params.Size)
// 	size := from + int32(params.Size) - 1
// 	searchResults, total, err := z.SearchDocument(c, constants.ZINCINDEXVIDEO, keyword, from, size)
// 	if err != nil {
// 		return nil, err
// 	}
// 	userIds := make([]uint, 0, len(searchResults))
// 	videoIds := make([]uint, 0, len(searchResults))
// 	categoryIds := make([]uint, 0, len(searchResults))
// 	for _, searchResult := range searchResults {
// 		userId := uint(searchResult.Source["user_id"].(float64))
// 		videoId, _ := strconv.Atoi(*searchResult.Id)
// 		categoryId := uint(searchResult.Source["category_id"].(float64))
// 		userIds = append(userIds, userId)
// 		videoIds = append(videoIds, uint(videoId))
// 		categoryIds = append(categoryIds, categoryId)
// 	}
// 	users, err := (&models.User{}).WhereAll(db, "id IN ?", userIds)
// 	if err != nil {
// 		return nil, err
// 	}
// 	userIdToUsernameMaps := make(map[uint]string, len(users))
// 	for _, user := range users {
// 		userIdToUsernameMaps[user.ID] = user.Username
// 	}
// 	categories, err := (&models.Category{}).WhereAll(db, "id IN ?", categoryIds)
// 	if err != nil {
// 		return nil, err
// 	}
// 	categoryIdToNameMaps := make(map[uint]string, len(categories))
// 	for _, category := range categories {
// 		categoryIdToNameMaps[category.ID] = category.Title
// 	}
// 	videos, err := (&models.Video{}).WhereAll(db, "id IN ?", videoIds)
// 	if err != nil {
// 		return nil, err
// 	}
// 	videoIdToVideoInfoMaps := make(map[uint]struct {
// 		ImageUrl     string
// 		VideoUrl     string
// 		ThumbCount   int64
// 		CommentCount int64
// 		CreatedAt    time.Time
// 		UpdatedAt    time.Time
// 	}, len(videos))
// 	for _, video := range videos {
// 		videoIdToVideoInfoMaps[video.ID] = struct {
// 			ImageUrl     string
// 			VideoUrl     string
// 			ThumbCount   int64
// 			CommentCount int64
// 			CreatedAt    time.Time
// 			UpdatedAt    time.Time
// 		}{ImageUrl: video.ImageUrl, VideoUrl: video.VideoUrl, ThumbCount: video.ThumbCount, CommentCount: video.CommentCount, CreatedAt: video.CreatedAt, UpdatedAt: video.UpdatedAt}
// 	}

// 	records := make([]forms.VideoListRecord, 0, len(searchResults))
// 	for _, searchResult := range searchResults {
// 		videoId, _ := strconv.Atoi(*searchResult.Id)
// 		userId := uint(searchResult.Source["user_id"].(float64))
// 		categoryId := uint(searchResult.Source["category_id"].(float64))
// 		records = append(records, forms.VideoListRecord{
// 			ID:           uint(videoId),
// 			UserID:       userId,
// 			UserName:     userIdToUsernameMaps[userId],
// 			CategoryID:   categoryId,
// 			CategoryName: categoryIdToNameMaps[categoryId],
// 			Title:        searchResult.Source["title"].(string),
// 			Introduce:    searchResult.Source["introduce"].(string),
// 			ImageUrl:     utils.FulfillImageOSSPrefix(videoIdToVideoInfoMaps[uint(videoId)].ImageUrl),
// 			VideoUrl:     utils.FulfillImageOSSPrefix(videoIdToVideoInfoMaps[uint(videoId)].VideoUrl),
// 			ThumbCount:   videoIdToVideoInfoMaps[uint(videoId)].ThumbCount,
// 			CommentCount: videoIdToVideoInfoMaps[uint(videoId)].CommentCount,
// 			CreatedAt:    videoIdToVideoInfoMaps[uint(videoId)].CreatedAt.Format(constants.DateTime),
// 			UpdatedAt:    videoIdToVideoInfoMaps[uint(videoId)].UpdatedAt.Format(constants.DateTime),
// 		})
// 	}
// 	result := &forms.VideoListResponse{
// 		Records: records,
// 		PageList: &utils.PageList{
// 			Size:    params.Size,
// 			Pages:   int64(math.Ceil(float64(total) / float64(params.Size))),
// 			Total:   total,
// 			Current: params.Page,
// 		},
// 	}
// 	return result, nil
// }

// func (s *Service) VideoDetail(c *gin.Context, id uint) (response *forms.VideoListRecord, err error) {
// 	db := global.DB

// 	query := []string{"id = ?"}
// 	args := []interface{}{id}
// 	video, err := (&models.Video{}).WhereOne(db, strings.Join(query, " AND "), args...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// 统计点赞数和评论数
// 	query = []string{"video_id = ?", "is_thumb = ?"}
// 	args = []interface{}{video.ID, "ISTHUMB"}
// 	thumbs, err := (&models.Comment{}).WhereCountGroup(db, "", strings.Join(query, " AND "))
// 	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
// 		return nil, err
// 	}
// 	query = []string{"video_id = ?", "is_thumb = ?"}
// 	args = []interface{}{video.ID, "ISCOMMENT"}
// 	comments, err := (&models.Comment{}).WhereCountGroup(db, "", strings.Join(query, " AND "))
// 	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
// 		return nil, err
// 	}
// 	var thumbCount, commentCount int64
// 	if len(thumbs) > 0 {
// 		thumbCount = int64(thumbs[0].CommentCount)
// 	}
// 	if len(comments) > 0 {
// 		commentCount = int64(comments[0].CommentCount)
// 	}
// 	// 封装数据返回
// 	response = &forms.VideoListRecord{
// 		ID:           video.ID,
// 		UserID:       video.UserId,
// 		CategoryID:   video.CategoryId,
// 		Title:        video.Title,
// 		Introduce:    video.Introduce,
// 		ImageUrl:     utils.FulfillImageOSSPrefix(video.ImageUrl),
// 		VideoUrl:     utils.FulfillImageOSSPrefix(video.VideoUrl),
// 		ThumbCount:   thumbCount,
// 		CommentCount: commentCount,
// 		CreatedAt:    video.CreatedAt.Format(constants.DateTime),
// 		UpdatedAt:    video.UpdatedAt.Format(constants.DateTime),
// 	}
// 	return response, nil
// }

// func (s *Service) VideoDelete(c *gin.Context, id uint) (err error) {
// 	// base logic: 删视频，删评论
// 	db := global.DB
// 	tx := db.Begin()

// 	query := []string{"id = ?"}
// 	args := []interface{}{id}
// 	_, err = (&models.Video{}).WhereOne(db, strings.Join(query, " AND "), args...)
// 	if err != nil {
// 		return err
// 	}
// 	err = (&models.Video{}).Delete(tx, strings.Join(query, " AND "), args...)
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}
// 	query = []string{"video_id = ?"}
// 	err = (&models.Comment{}).Delete(tx, strings.Join(query, " AND "), args...)
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}
// 	_ = tx.Commit().Error
// 	z := NewZinc()
// 	err = z.DeleteDocument(c, constants.ZINCINDEXVIDEO, strconv.Itoa(int(id)))
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (s *Service) VideoInsert(c *gin.Context, params *forms.VideoInsertForm) (id uint, err error) {
// 	db := global.DB
// 	user := utils.GetUser(c)

// 	// base logic: 查看分类是否存在，若存在，则创建视频
// 	query := []string{"id = ?"}
// 	args := []interface{}{params.CategoryId}
// 	_, err = (&models.Category{}).WhereOne(db, strings.Join(query, " AND "), args...)
// 	if err != nil {
// 		return 0, err
// 	}
// 	video := &models.Video{
// 		UserId:       user.ID,
// 		CategoryId:   params.CategoryId,
// 		Title:        params.Title,
// 		Introduce:    params.Introduce,
// 		ImageUrl:     utils.TrimDomainPrefix(params.ImageUrl),
// 		VideoUrl:     utils.TrimDomainPrefix(params.VideoUrl),
// 		ThumbCount:   0,
// 		CommentCount: 0,
// 	}
// 	err = video.Insert(db)
// 	if err != nil {
// 		return 0, err
// 	}
// 	z := NewZinc()
// 	err = z.InsertDocument(c, constants.ZINCINDEXVIDEO, strconv.Itoa(int(video.ID)), map[string]interface{}{
// 		"user_id":     video.UserId,
// 		"category_id": video.CategoryId,
// 		"title":       video.Title,
// 		"introduce":   video.Introduce,
// 	})
// 	if err != nil {
// 		return 0, err
// 	}
// 	return video.ID, nil
// }

func (s *Service) VideoUploadImg(c *gin.Context, file *multipart.FileHeader) (result string, err error) {
	user := utils.GetUser(c)

	folder := "video.server.videos.img"

	url, err := UploadImg(user.ID, folder, file.Filename, file, "image")
	if err != nil {
		return "", err
	}

	return utils.FulfillImageOSSPrefix(utils.TrimDomainPrefix(url)), nil
}

func (s *Service) VideoUploadVid(c *gin.Context, file *multipart.FileHeader) (result string, err error) {
	user := utils.GetUser(c)

	folder := "video.server.videos.vid"

	url, err := UploadVid(user.ID, folder, file.Filename, file, "video")
	if err != nil {
		return "", err
	}
	return utils.FulfillImageOSSPrefix(utils.TrimDomainPrefix(url)), nil
}

type VideoService struct{}

func (*VideoService) Search(c *gin.Context, params *forms.VideoSearchForm) {
	query := []string{"1 = ?"}
	args := []any{1}
	if params.Name != "" {
		query = append(query, "title LIKE ?")
		args = append(args, "%"+params.Name+"%")
	}
	db := global.DB
	sqlVideos, total, pages, err := models.GPaginate[models.Video](db, &models.ListPageInput{
		Page: params.Page,
		Size: params.Size,
	}, strings.Join(query, " AND "), args...)
	if err != nil {
		response.Error(c, constants.InternalServerErrorCode, err)
		return
	}

	videoIds := make([]uint, 0, len(sqlVideos))
	for i := 0; i != len(sqlVideos); i++ {
		videoIds = append(videoIds, sqlVideos[i].ID)
	}
	sqlRelation, err := models.GWhereAllSelect[models.CategoryVideo](db, "category_id,video_id", "video_id IN (?)", videoIds)
	if err != nil {
		response.Error(c, constants.InternalServerErrorCode, err)
		return
	}
	categoryIds := make([]uint, 0, len(sqlRelation))
	videoIdMap := make(map[uint][]uint, len(sqlRelation))
	for i := 0; i != len(sqlRelation); i++ {
		categoryIds = append(categoryIds, sqlRelation[i].CategoryID)
		videoIdMap[sqlRelation[i].VideoID] = append(videoIdMap[sqlRelation[i].VideoID], sqlRelation[i].CategoryID)
	}
	sqlCategories, err := models.GWhereAllSelect[models.Category](db, "id,title", "id IN (?)", categoryIds)
	if err != nil {
		response.Error(c, constants.InternalServerErrorCode, err)
		return
	}
	categoryIdMap := make(map[uint]string, len(sqlCategories))
	for i := 0; i != len(sqlCategories); i++ {
		categoryIdMap[sqlCategories[i].ID] = sqlCategories[i].Title
	}

	records := make([]forms.VideoSearchRecord, 0, len(sqlVideos))
	for i := 0; i != len(sqlVideos); i++ {
		categories := make([]string, 0, len(videoIdMap[sqlVideos[i].ID]))
		for j := 0; j != len(videoIdMap[sqlVideos[i].ID]); j++ {
			categories = append(categories, categoryIdMap[videoIdMap[sqlVideos[i].ID][j]])
		}
		season := "更新"
		if sqlVideos[i].Season == models.VideoSeasonFinish {
			season = "完结"
		}
		records = append(records, forms.VideoSearchRecord{
			ID:          sqlVideos[i].ID,
			Title:       sqlVideos[i].Title,
			Season:      season,
			Date:        sqlVideos[i].FirstDate.Format(constants.DateTime),
			Categories:  categories,
			Description: sqlVideos[i].Introduce,
			Cover:       utils.FulfillImageOSSPrefix(sqlVideos[i].Cover),
		})
	}

	response.Success(c, forms.VideoSearch{
		PageList: utils.PageList{
			Size:    params.Size,
			Pages:   pages,
			Total:   total,
			Current: params.Page,
		},
		Records: records,
	})
}

func (*VideoService) Detail(c *gin.Context, videoId uint) {
	db := global.DB
	sqlVideo, err := models.GWhereFirstSelect[models.Video](db, "*", "id = ?", videoId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		response.Error(c, constants.InternalServerErrorCode, err)
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.Error(c, constants.BadRequestCode, errors.New("视频不存在"))
		return
	}
	// 查找标签列表
	sqlRelation, err := models.GWhereAllSelect[models.CategoryVideo](db, "category_id,video_id", "video_id = ?", videoId)
	if err != nil {
		response.Error(c, constants.InternalServerErrorCode, err)
		return
	}
	categoryIds := make([]uint, 0, len(sqlRelation))
	for i := 0; i != len(sqlRelation); i++ {
		categoryIds = append(categoryIds, sqlRelation[i].CategoryID)
	}
	sqlCategories, err := models.GWhereAllSelect[models.Category](db, "id,title", "id IN (?)", categoryIds)
	if err != nil {
		response.Error(c, constants.InternalServerErrorCode, err)
		return
	}
	categories := make([]string, 0, len(sqlCategories))
	for i := 0; i != len(sqlCategories); i++ {
		categories = append(categories, sqlCategories[i].Title)
	}
	// 查找播放列表
	sqlPlaylists, err := models.GWhereAllSelectOrder[models.Playlist](db, "*", "sort ASC", "video_id = ?", videoId)
	if err != nil {
		response.Error(c, constants.InternalServerErrorCode, err)
		return
	}
	playlist := make([]forms.Playlist, 0, len(sqlPlaylists))
	for i := 0; i != len(sqlPlaylists); i++ {
		playlist = append(playlist, forms.Playlist{
			ID:    sqlPlaylists[i].ID,
			Title: sqlPlaylists[i].Title,
			Link:  []string{utils.FulfillImageOSSPrefix(sqlPlaylists[i].Link)},
			Sort:  sqlPlaylists[i].Sort,
		})
	}

	// 拼接参数返回
	lang := ""
	switch sqlVideo.Language {
	case models.VideoLanguageChinese:
		lang = "简体中文"
	case models.VideoLanguageEnglish:
		lang = "英文"
	case models.VideoLanguageJapan:
		lang = "日文"
	default:
		response.Error(c, constants.BadRequestCode, errors.New("语言类型错误"))
		return
	}

	region := ""
	switch sqlVideo.Region {
	case models.VideoRegionChina:
		region = "中国"
	case models.VideoRegionJapan:
		region = "日本"
	case models.VideoRegionAmerica:
		region = "美国"
	default:
		response.Error(c, constants.BadRequestCode, errors.New("地区类型错误"))
		return
	}

	season := "更新"
	if sqlVideo.Season == models.VideoSeasonFinish {
		season = "完结"
	}

	playlistMap := map[string][]forms.Playlist{}
	playlistMap["0"] = playlist
	response.Success(c, forms.VideoDetail{
		Actors:     strings.Split(sqlVideo.Actors, ","),
		Categories: categories,
		Cover:      utils.FulfillImageOSSPrefix(sqlVideo.Cover),
		Date:       sqlVideo.FirstDate.Format(constants.DateTime),
		Lang:       lang,
		Master:     sqlVideo.Master,
		Playlist:   playlistMap,
		Score:      sqlVideo.Score,
		Rank:       sqlVideo.Ranking,
		Region:     region,
		Season:     season,
		Title:      sqlVideo.Title,
	})
}

func (*VideoService) Filter(c *gin.Context, params *forms.VideoFilterForm) {
	query := []string{"1 = ?"}
	args := []any{1}
	order := "created_at desc"

	db := global.DB
	if params.Type != 0 {
		// 暂时忽略
	}
	if params.CategoryID != 0 {
		sqlRelation, err := models.GWhereAllSelect[models.CategoryVideo](db, "video_id", "category_id = ?", params.CategoryID)
		if err != nil {
			response.Error(c, constants.InternalServerErrorCode, err)
			return
		}
		videoIds := make([]uint, 0, len(sqlRelation))
		for i := 0; i != len(sqlRelation); i++ {
			videoIds = append(videoIds, sqlRelation[i].VideoID)
		}

		query = append(query, "id IN (?)")
		args = append(args, videoIds)
	}
	if params.Order != "" {
		switch params.Order {
		case "time":
			order = "updated_at DESC"
		case "score":
			order = "score DESC"
		case "hits":
			order = "ranking DESC"
		default:
			response.Error(c, constants.BadRequestCode, errors.New("order 参数错误"))
			return
		}
	}
	if params.Letter != "" {
		query = append(query, "title LIKE ?")
		args = append(args, params.Letter+"%")
	}
	if params.Year != 0 {
		paramTime := time.Date(params.Year, 1, 1, 0, 0, 0, 0, time.Local)
		begin := now.With(paramTime).BeginningOfYear()
		end := now.With(paramTime).EndOfYear()
		query = append(query, "first_date BETWEEN ? AND ?")
		args = append(args, begin, end)
	}

	sqlVideos, total, pages, err := models.GPaginateOrder[models.Video](db, &models.ListPageInput{
		Page: params.Page,
		Size: params.Size,
	}, order, strings.Join(query, " AND "), args...)
	if err != nil {
		response.Error(c, constants.InternalServerErrorCode, err)
		return
	}

	records := make([]forms.VideoFilterRecord, 0, len(sqlVideos))
	for i := 0; i != len(sqlVideos); i++ {
		season := "更新"
		if sqlVideos[i].Season == models.VideoSeasonFinish {
			season = "完结"
		}
		records = append(records, forms.VideoFilterRecord{
			Cover:  utils.FulfillImageOSSPrefix(sqlVideos[i].Cover),
			ID:     sqlVideos[i].ID,
			Season: season,
			Title:  sqlVideos[i].Title,
		})
	}

	response.Success(c, forms.VideoFilter{
		PageList: utils.PageList{
			Size:    params.Size,
			Pages:   pages,
			Total:   total,
			Current: params.Page,
		},
		Records: records,
	})
}

func (*VideoService) Playlist(c *gin.Context, videoId uint) {
	db := global.DB
	sqlPlaylist, err := models.GWhereAllSelectOrder[models.Playlist](db, "*", "sort ASC", "video_id = ?", videoId)
	if err != nil {
		response.Error(c, constants.InternalServerErrorCode, err)
		return
	}

	// records := make([]forms.Playlist, 0, len(sqlPlaylist))
	// for i := 0; i != len(sqlPlaylist); i++ {
	// 	records = append(records, forms.Playlist{
	// 		ID:    sqlPlaylist[i].ID,
	// 		Link:  []string{utils.FulfillImageOSSPrefix(sqlPlaylist[i].Link)},
	// 		Sort:  sqlPlaylist[i].Sort,
	// 		Title: sqlPlaylist[i].Title,
	// 	})
	// }

	records := make([]string, 0, len(sqlPlaylist))
	for i := 0; i != len(sqlPlaylist); i++ {
		records = append(records, utils.FulfillImageOSSPrefix(sqlPlaylist[i].Link))
	}

	response.Success(c, map[int][]string{0: records})
}

func (*VideoService) Config(c *gin.Context) {
	db := global.DB
	sqlCategories, err := models.GWhereAllSelectOrder[models.Category](db, "*", "id DESC", "1 = ?", 1)
	if err != nil {
		response.Error(c, constants.InternalServerErrorCode, err)
		return
	}

	records := make([]forms.Category, 0, len(sqlCategories))
	for i := 0; i != len(sqlCategories); i++ {
		records = append(records, forms.Category{
			ID:   sqlCategories[i].ID,
			Name: sqlCategories[i].Title,
		})
	}

	response.Success(c, forms.VideoConfig{
		FiltersConfig: []forms.FilterConfig{
			{Name: "全部", ID: 0, Categories: records},
		},
	})
}

func (*VideoService) Index(c *gin.Context) {
	db := global.DB
	sqlCategories, err := models.GWhereAllSelect[models.Category](db, "id,title", "title IN (?)", servants.IndexPlate)
	if err != nil {
		response.Error(c, constants.InternalServerErrorCode, err)
		return
	}
	categoryIds := make([]uint, 0, len(sqlCategories))
	categoryNameMap := make(map[string]uint, len(sqlCategories))
	for i := 0; i != len(sqlCategories); i++ {
		categoryIds = append(categoryIds, sqlCategories[i].ID)
		categoryNameMap[sqlCategories[i].Title] = sqlCategories[i].ID
	}
	sqlRelation, err := models.GWhereAllSelect[models.CategoryVideo](db, "category_id,video_id", "category_id IN (?)", categoryIds)
	if err != nil {
		response.Error(c, constants.InternalServerErrorCode, err)
		return
	}
	videoIds := make([]uint, 0, len(sqlRelation))
	categoryIdMap := make(map[uint][]uint, len(sqlCategories))
	for i := 0; i != len(sqlRelation); i++ {
		videoIds = append(videoIds, sqlRelation[i].VideoID)
		categoryIdMap[sqlRelation[i].CategoryID] = append(categoryIdMap[sqlRelation[i].CategoryID], sqlRelation[i].VideoID)
	}

	sqlVideos, err := models.GWhereAllSelectOrder[models.Video](db, "*", "created_at DESC", "id IN (?)", videoIds)
	if err != nil {
		response.Error(c, constants.InternalServerErrorCode, err)
		return
	}
	videoIdMap := make(map[uint]models.Video, len(sqlVideos))
	for i := 0; i != len(sqlVideos); i++ {
		videoIdMap[sqlVideos[i].ID] = sqlVideos[i]
	}

	// 收集数据返回
	banners := make([]forms.VideoRecord, 0, len(sqlVideos)/7+1)
	chineseComics := make([]forms.VideoRecord, 0, len(sqlVideos)/7+1)
	hots := make([]forms.VideoRecord, 0, len(sqlVideos)/7+1)
	japancomic := make([]forms.VideoRecord, 0, len(sqlVideos)/7+1)
	latest := make([]forms.VideoRecord, 0, len(sqlVideos)/7+1)
	perweek := make(map[int][]forms.VideoRecord, len(sqlVideos)/7+1)
	theatreComic := make([]forms.VideoRecord, 0, len(sqlVideos)/7+1)

	if id, exist := categoryNameMap[servants.IndexPlate[0]]; exist {
		for i := 0; i != len(categoryIdMap[id]); i++ {
			sqlVideo := videoIdMap[categoryIdMap[id][i]]
			season := "更新"
			if sqlVideos[i].Season == models.VideoSeasonFinish {
				season = "完结"
			}
			banners = append(banners, forms.VideoRecord{
				ID:          sqlVideo.ID,
				Cover:       utils.FulfillImageOSSPrefix(sqlVideo.Cover),
				Title:       sqlVideo.Title,
				Season:      season,
				Date:        sqlVideo.FirstDate.Format(constants.DateTime),
				Description: sqlVideo.Introduce,
			})
		}
	}
	if id, exist := categoryNameMap[servants.IndexPlate[1]]; exist {
		for i := 0; i != len(categoryIdMap[id]); i++ {
			sqlVideo := videoIdMap[categoryIdMap[id][i]]
			season := "更新"
			if sqlVideos[i].Season == models.VideoSeasonFinish {
				season = "完结"
			}
			chineseComics = append(chineseComics, forms.VideoRecord{
				ID:          sqlVideo.ID,
				Cover:       utils.FulfillImageOSSPrefix(sqlVideo.Cover),
				Title:       sqlVideo.Title,
				Season:      season,
				Date:        sqlVideo.FirstDate.Format(constants.DateTime),
				Description: sqlVideo.Introduce,
			})
		}
	}
	if id, exist := categoryNameMap[servants.IndexPlate[2]]; exist {
		for i := 0; i != len(categoryIdMap[id]); i++ {
			sqlVideo := videoIdMap[categoryIdMap[id][i]]
			season := "更新"
			if sqlVideos[i].Season == models.VideoSeasonFinish {
				season = "完结"
			}
			hots = append(hots, forms.VideoRecord{
				ID:          sqlVideo.ID,
				Cover:       utils.FulfillImageOSSPrefix(sqlVideo.Cover),
				Title:       sqlVideo.Title,
				Season:      season,
				Date:        sqlVideo.FirstDate.Format(constants.DateTime),
				Description: sqlVideo.Introduce,
			})
		}
	}
	if id, exist := categoryNameMap[servants.IndexPlate[3]]; exist {
		for i := 0; i != len(categoryIdMap[id]); i++ {
			sqlVideo := videoIdMap[categoryIdMap[id][i]]
			season := "更新"
			if sqlVideos[i].Season == models.VideoSeasonFinish {
				season = "完结"
			}
			japancomic = append(japancomic, forms.VideoRecord{
				ID:          sqlVideo.ID,
				Cover:       utils.FulfillImageOSSPrefix(sqlVideo.Cover),
				Title:       sqlVideo.Title,
				Season:      season,
				Date:        sqlVideo.FirstDate.Format(constants.DateTime),
				Description: sqlVideo.Introduce,
			})
		}
	}
	if id, exist := categoryNameMap[servants.IndexPlate[4]]; exist {
		for i := 0; i != len(categoryIdMap[id]); i++ {
			sqlVideo := videoIdMap[categoryIdMap[id][i]]
			season := "更新"
			if sqlVideos[i].Season == models.VideoSeasonFinish {
				season = "完结"
			}
			latest = append(latest, forms.VideoRecord{
				ID:          sqlVideo.ID,
				Cover:       utils.FulfillImageOSSPrefix(sqlVideo.Cover),
				Title:       sqlVideo.Title,
				Season:      season,
				Date:        sqlVideo.FirstDate.Format(constants.DateTime),
				Description: sqlVideo.Introduce,
			})
		}
	}
	if id, exist := categoryNameMap[servants.IndexPlate[5]]; exist {
		for i := 0; i != len(categoryIdMap[id]); i++ {
			sqlVideo := videoIdMap[categoryIdMap[id][i]]
			season := "更新"
			if sqlVideos[i].Season == models.VideoSeasonFinish {
				season = "完结"
			}

			// 先随机模拟一个更新日期
			rand.Seed(time.Now().UnixNano())
			key := rand.Intn(7)

			perweek[key] = append(perweek[key], forms.VideoRecord{
				ID:          sqlVideo.ID,
				Cover:       utils.FulfillImageOSSPrefix(sqlVideo.Cover),
				Title:       sqlVideo.Title,
				Season:      season,
				Date:        sqlVideo.FirstDate.Format(constants.DateTime),
				Description: sqlVideo.Introduce,
			})
		}
	}
	if id, exist := categoryNameMap[servants.IndexPlate[6]]; exist {
		for i := 0; i != len(categoryIdMap[id]); i++ {
			sqlVideo := videoIdMap[categoryIdMap[id][i]]
			season := "更新"
			if sqlVideo.Season == models.VideoSeasonFinish {
				season = "完结"
			}
			theatreComic = append(theatreComic, forms.VideoRecord{
				ID:          sqlVideo.ID,
				Cover:       utils.FulfillImageOSSPrefix(sqlVideo.Cover),
				Title:       sqlVideo.Title,
				Season:      season,
				Date:        sqlVideo.FirstDate.Format(constants.DateTime),
				Description: sqlVideo.Introduce,
			})
		}
	}

	response.Success(c, forms.VideoIndex{
		Banners:       banners,
		ChineseComics: chineseComics,
		Hots:          map[string][]forms.VideoRecord{"results": hots},
		Japancomic:    japancomic,
		Latest:        latest,
		Perweek:       perweek,
		TheatreComic:  theatreComic,
	})
}

func (*VideoService) Delete(c *gin.Context, videoID uint) {
	db := global.DB
	sqlVideo, err := models.GWhereFirstSelect[models.Video](db, "creator_id", "id = ?", videoID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		response.Error(c, constants.InternalServerErrorCode, err)
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.Error(c, constants.BadRequestCode, errors.New("视频不存在"))
		return
	}
	if sqlVideo.CreatorId != uint(c.Value("userId").(int)) {
		response.Error(c, constants.BadRequestCode, errors.New("没有权限"))
		return
	}

	tx := db.Begin()
	commit := false
	defer func() {
		if commit {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	// 删除视频本身
	if err = models.GDelete[models.Video](tx, "id = ?", videoID); err != nil {
		response.Error(c, constants.InternalServerErrorCode, err)
		return
	}
	// 删除视频下属播放列表
	if err = models.GDelete[models.Playlist](tx, "video_id = ?", videoID); err != nil {
		response.Error(c, constants.InternalServerErrorCode, err)
		return
	}
	// 删除视频与模块关系
	if err = models.GDelete[models.CategoryVideo](tx, "video_id = ?", videoID); err != nil {
		response.Error(c, constants.InternalServerErrorCode, err)
		return
	}

	commit = true
	response.Success(c, "success")
}
