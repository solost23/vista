package global

import (
	"vista/configs"

	"github.com/go-redis/redis/v8"
	"github.com/minio/minio-go"
	"github.com/solost23/go_interface/gen_go/oss"
	"gorm.io/gorm"
)

var (
	ServerConfig = &configs.ServerConfig{}
	DB           *gorm.DB
	Minio        *minio.Client
	RedisMapPool = make(map[int]*redis.Client, 15)
	OSSSrvClient oss.OssClient
)
