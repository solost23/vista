package scheduler

import (
	"vista/global"

	"github.com/robfig/cron/v3"

	"vista/scheduler/deleteVideo"
)

// 删除数据库中视频数据记录
// 删除本地视频文件
func Run() {
	var task = new(deleteVideo.Task)
	c := cron.New()
	_, _ = c.AddFunc(global.ServerConfig.DeleteCronTime, func() {
		task.DeleteVideo()
	})
	c.Start()
	select {}
}
