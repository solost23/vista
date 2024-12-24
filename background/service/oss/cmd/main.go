package main

import (
	"flag"
	"fmt"
	"os"
	"oss_service/configs"
	"oss_service/internal/server"
	"oss_service/pkg/helper"
	"path"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"go.uber.org/zap"
)

var (
	WebConfigPath = "configs/conf.yml"
	WebLogPath    = "logs"
	version       = "__BUILD_VERSION_"
	execDir       string
	provider      string
	st, v, V      bool
)

func main() {
	flag.StringVar(&execDir, "d", ".", "项目目录")
	flag.StringVar(&provider, "p", "consul", "项目配置提供者")
	flag.BoolVar(&v, "v", false, "查看版本号")
	flag.BoolVar(&V, "V", false, "查看版本号")
	flag.BoolVar(&st, "s", false, "项目状态")
	flag.Parse()
	if v || V {
		fmt.Println(version)
		os.Exit(-1)
	}
	// run
	// 从consul读取配置
	serverConfig := InitConfig()
	sl, err := helper.InitLogger(execDir, WebLogPath, serverConfig.Mode)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	server.NewServer(serverConfig, sl).Run()
}

func InitConfig() *configs.ServerConfig {
	v := viper.New()
	// 通过项目配置读取基本信息
	v.SetConfigFile(path.Join(execDir, WebConfigPath))
	if err := v.ReadInConfig(); err != nil {
		zap.S().Panic("read config failed", zap.Error(err))
		return nil
	}

	serverConfig := &configs.ServerConfig{}
	if err := v.Unmarshal(serverConfig); err != nil {
		zap.S().Panic("unmarshal config failed", zap.Error(err))
		return nil
	}

	// 从配置中心读取配置
	err := v.AddRemoteProvider(provider,
		fmt.Sprintf("%s:%d", serverConfig.ConsulConfig.Host, serverConfig.ConsulConfig.Port),
		serverConfig.ConfigPath)
	if err != nil {
		zap.S().Warn("add remote provider failed")
		return serverConfig
	}

	v.SetConfigType("YAML")

	if err = v.ReadRemoteConfig(); err != nil {
		zap.S().Warn("read remote config failed")
		return serverConfig
	}

	if err = v.Unmarshal(serverConfig); err != nil {
		zap.S().Panic("unmarshal remote config failed")
		return serverConfig
	}
	return serverConfig

}
