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

type LoginController struct{}

func LoginRegister(router *gin.RouterGroup) {
	controller := &LoginController{}

	// 头像上传
	router.POST("upload", controller.upload)
	// 注册
	router.POST("register", controller.register)
	// 登陆
	router.POST("login", controller.Login)
	// 登出
	apiRouter := router.Group("")
	apiRouter.Use(middlewares.JWTAuth()).DELETE("logout", controller.Logout)
}

func (controller *LoginController) upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.Error(c, constants.BadRequestCode, err)
		return
	}

	folder := "vista.static.avatar"
	url, err := services.UploadImg(0, folder, file.Filename, file, "image")
	if err != nil {
		response.Error(c, constants.InternalServerErrorCode, err)
		return
	}

	response.Success(c, utils.FulfillImageOSSPrefix(utils.TrimDomainPrefix(url)))
}

func (controller *LoginController) register(c *gin.Context) {
	params := &forms.RegisterForm{}
	if err := utils.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, constants.BadRequestCode, err)
		return
	}

	if params.Nickname == "" {
		params.Nickname = *params.Username
	}

	loginService.Register(c, params)
}

func (controller *LoginController) Login(c *gin.Context) {
	params := &forms.LoginForm{}
	if err := utils.DefaultGetValidParams(c, params); err != nil {
		response.Error(c, 2001, err)
		return
	}

	loginService.Login(c, params)
}

func (controller *LoginController) Logout(c *gin.Context) {
	loginService.Logout(c)
}
