package controllers

import (
	"vista/forms"
	"vista/pkg/constants"
	"vista/pkg/response"
	"vista/pkg/utils"

	"github.com/gin-gonic/gin"
)

type VideoController struct{}

func VideoRegister(router *gin.RouterGroup) {
	controller := &VideoController{}

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

}

func (*VideoController) filter(c *gin.Context) {

}

func (*VideoController) playlist(c *gin.Context) {

}

func (*VideoController) config(c *gin.Context) {

}

func (*VideoController) index(c *gin.Context) {

}
