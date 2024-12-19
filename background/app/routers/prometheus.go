package routers

import (
	"vista/global"

	"github.com/gin-gonic/gin"
	ginPrometheus "github.com/zsais/go-gin-prometheus"
)

// SetPrometheus sets up prometheus metrics for gin
func SetPrometheus(app *gin.Engine) {
	if !global.ServerConfig.PrometheusEnable {
		return
	}

	ginPrometheus.NewPrometheus("video_server").Use(app)
}
