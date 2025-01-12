package controllers

import (
	"vista/forms"
	"vista/pkg/constants"
	"vista/pkg/middlewares"
	"vista/pkg/response"
	"vista/pkg/utils"
	"vista/services"

	"github.com/gin-gonic/gin"
)

type VideoController struct{}

func VideoRegister(router *gin.RouterGroup) {
	controller := &VideoController{}

	// 上传视频封面
	router.POST("upload/cover", controller.uploadCover)
	// 上传视频
	router.POST("upload", controller.upload)
	// 根据名称获取视频列表
	router.GET("search", controller.search)
	// 获取动漫详情
	router.GET(":id", controller.detail)
	// 动漫筛选
	router.GET("filter", controller.filter)

	// 获取视频地址集
	router.GET(":id/playlist", controller.playlist)
	// 获取动漫配置
	router.GET("config", controller.config)
	// 获取混合列表
	router.GET("index", controller.index)

	// 删除视频
	apiRouter := router.Group("")
	apiRouter.Use(middlewares.JWTAuth()).DELETE(":id", controller.delete)
}

func (*VideoController) uploadCover(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.Error(c, constants.BadRequestCode, err)
		return
	}

	folder := "vista.static.cover"
	url, err := services.UploadImg(0, folder, file.Filename, file, "image")
	if err != nil {
		response.Error(c, constants.InternalServerErrorCode, err)
		return
	}

	response.Success(c, utils.FulfillImageOSSPrefix(utils.TrimDomainPrefix(url)))
}

func (*VideoController) upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.Error(c, constants.BadRequestCode, err)
		return
	}

	folder := "vista.dynamic.video"
	url, err := services.UploadImg(0, folder, file.Filename, file, "video")
	if err != nil {
		response.Error(c, constants.InternalServerErrorCode, err)
		return
	}

	response.Success(c, utils.FulfillImageOSSPrefix(utils.TrimDomainPrefix(url)))
}

func (*VideoController) search(c *gin.Context) {
	params := &forms.VideoSearchForm{}
	if err := utils.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, constants.BadRequestCode, err)
		return
	}

	if params.Page == 0 {
		params.Page = 1
	}
	if params.Size == 0 {
		params.Size = 20
	}

	videoService.Search(c, params)
}

func (*VideoController) detail(c *gin.Context) {
	UID := &utils.UIdForm{}
	if err := utils.GetValidUriParams(c, UID); err != nil {
		response.Error(c, constants.BadRequestCode, err)
		return
	}

	videoService.Detail(c, UID.Id)
}

func (*VideoController) filter(c *gin.Context) {
	params := &forms.VideoFilterForm{}
	if err := utils.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, constants.BadRequestCode, err)
		return
	}

	if params.Page == 0 {
		params.Page = 1
	}
	if params.Size == 0 {
		params.Size = 20
	}

	videoService.Filter(c, params)
}

func (*VideoController) playlist(c *gin.Context) {
	UID := &utils.UIdForm{}
	if err := utils.GetValidUriParams(c, UID); err != nil {
		response.Error(c, constants.BadRequestCode, err)
		return
	}

	videoService.Playlist(c, UID.Id)
}

func (*VideoController) config(c *gin.Context) {
	videoService.Config(c)
}

func (*VideoController) index(c *gin.Context) {
	videoService.Index(c)
}

func (*VideoController) delete(c *gin.Context) {
	UID := &utils.UIdForm{}
	if err := utils.GetValidUriParams(c, UID); err != nil {
		response.Error(c, constants.BadRequestCode, err)
		return
	}

	videoService.Delete(c, UID.Id)
}
