package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"
	"vista/global"
	"vista/global/initialize"
	"vista/pkg/models"
	"vista/pkg/utils"
	"vista/routers"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	WebConfigPath = "configs/config.yml"
	WebLogPath    = "logs"
	version       = "__BUILD_VERSION_"
	execDir       string
	st, v, V      bool
)

func main() {
	flag.StringVar(&execDir, "d", ".", "项目目录")
	flag.BoolVar(&v, "v", false, "查看版本号")
	flag.BoolVar(&V, "V", false, "查看版本号")
	flag.BoolVar(&st, "s", false, "项目状态")
	flag.Parse()
	if v || V {
		fmt.Println(version)
		os.Exit(-1)
	}

	initialize.Initialize(path.Join(execDir, WebConfigPath))
	if err := global.DB.AutoMigrate(&models.User{}, &models.Category{}, &models.Video{}, &models.CategoryVideo{}, &models.Playlist{}); err != nil {
		zap.S().Panic(err)
	}

	// 表记录初始化
	{
		db := global.DB
		// 创建管理员
		sqlCnt, err := models.GWhereCount[models.User](db, "1 = ?", 1)
		if err != nil {
			zap.S().Panic(err)
		}
		userId := uint(0)
		if sqlCnt == 0 {
			sqlUser := &models.User{
				Username:      "admin",
				Password:      utils.NewMd5("admin", global.ServerConfig.Md5Config.Secret),
				Nickname:      "admin",
				Role:          models.UserRoleAdmin,
				Introduce:     "system init user",
				FansCount:     0,
				CommentCount:  0,
				LastLoginTime: time.Now(),
			}
			if err = models.GInsert(db, sqlUser); err != nil {
				zap.S().Panic(err)
			}
			userId = sqlUser.ID
		}
		// 创建视频分类
		sqlCnt, err = models.GWhereCount[models.Category](db, "1 = ?", 1)
		if err != nil {
			zap.S().Panic(err)
		}
		if sqlCnt == 0 {
			sqlCategories := []models.Category{
				{CreatorBase: models.CreatorBase{CreatorId: userId}, Title: "轮播", Introduce: "轮播"},
				{CreatorBase: models.CreatorBase{CreatorId: userId}, Title: "国漫", Introduce: "国漫"},
				{CreatorBase: models.CreatorBase{CreatorId: userId}, Title: "热门动漫", Introduce: "热门动漫"},
				{CreatorBase: models.CreatorBase{CreatorId: userId}, Title: "完结动漫", Introduce: "完结动漫"},
				{CreatorBase: models.CreatorBase{CreatorId: userId}, Title: "最新更新", Introduce: "最新更新"},
				{CreatorBase: models.CreatorBase{CreatorId: userId}, Title: "周更新列表", Introduce: "周更新列表"},
				{CreatorBase: models.CreatorBase{CreatorId: userId}, Title: "推荐动漫番外", Introduce: "推荐动漫番外"},
			}
			if err = models.GBatchInsert[models.Category](db, sqlCategories); err != nil {
				zap.S().Panic(err)
			}
		}
	}

	// HTTP init
	app := gin.New()
	routers.Setup(app)

	server := &http.Server{
		Addr:         global.ServerConfig.Addr,
		Handler:      app,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	// 启动定时任务，每天晚上三点删除视频文件，降低用户删除请求的io操作过多
	// go func() {
	// 	fmt.Println("start scheduler")
	// 	scheduler.Run()
	// }()
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			if err == http.ErrServerClosed {
				log.Panicf("Server closed")
			} else {
				log.Panicf("Server closed unexpect %s", err.Error())
			}
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-c
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			_ = server.Close()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
