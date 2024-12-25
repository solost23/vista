package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
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

			// 初始化三个视频信息，方便调试，后面删除
			{
				sqlVideos := []models.Video{
					{CreatorBase: models.CreatorBase{CreatorId: userId}, Title: "蜡笔小新", Introduce: "蜡笔小新", Language: models.VideoLanguageJapan, Master: "臼井仪人", FirstDate: time.Now(), Ranking: 1, Score: 90.1, Region: models.VideoRegionJapan, Season: models.VideoSeasonUpdate, Actors: "小红,小明"},
					{CreatorBase: models.CreatorBase{CreatorId: userId}, Title: "龙猫", Introduce: "龙猫", Language: models.VideoLanguageJapan, Master: "宫崎骏", FirstDate: time.Now(), Ranking: 2, Score: 90.2, Region: models.VideoRegionJapan, Season: models.VideoSeasonFinish, Actors: "小明,小红"},
					{CreatorBase: models.CreatorBase{CreatorId: userId}, Title: "天空之城", Introduce: "天空之城", Language: models.VideoLanguageJapan, Master: "宫崎骏", FirstDate: time.Now(), Ranking: 3, Score: 90.0, Region: models.VideoRegionJapan, Season: models.VideoSeasonFinish, Actors: "小明,小李"},
				}
				if err = models.GBatchInsert[models.Video](db, sqlVideos); err != nil {
					zap.S().Panic(err)
				}
				// 插入关系表
				sqlCategoryVideos := make([]models.CategoryVideo, 0, 3*len(sqlVideos))
				rand.Seed(time.Now().UnixNano())
				for i := 0; i != len(sqlVideos); i++ {
					for j := 0; j != 3; j++ {
						sqlCategoryVideos = append(sqlCategoryVideos, models.CategoryVideo{
							CreatorBase: models.CreatorBase{CreatorId: userId},
							VideoID:     sqlVideos[i].ID,
							CategoryID:  sqlCategories[rand.Intn(len(sqlCategories))].ID,
						})
					}
				}
				if err = models.GBatchInsert[models.CategoryVideo](db, sqlCategoryVideos); err != nil {
					zap.S().Panic(err)
				}
				// 创建播放列表
				sqlPlaylist := make([]models.Playlist, 0, 3*len(sqlVideos))
				for i := 0; i != len(sqlVideos); i++ {
					for j := 0; j != 3; j++ {
						sqlPlaylist = append(sqlPlaylist, models.Playlist{
							CreatorBase: models.CreatorBase{CreatorId: userId},
							VideoID:     sqlVideos[i].ID,
							Title:       fmt.Sprintf("第%d集", j+1),
							Sort:        j + 1,
						})
					}
				}
				if err = models.GBatchInsert[models.Playlist](db, sqlPlaylist); err != nil {
					zap.S().Panic(err)
				}
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
