package routers

import (
	"vista/global"

	"github.com/gin-gonic/gin"
)

func Setup(app *gin.Engine) {
	gin.SetMode(global.ServerConfig.Mode)
	app.Use(gin.Logger(), gin.Recovery())

	// Debug for gin
	if gin.Mode() == gin.DebugMode {
		SetPProf(app)
	}
	SetPrometheus(app) // Set up Prometheus.
	SetRouters(app)    // Set up all API routers.
}
