package initialize

import (
	"fmt"
	"vista/global"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"go.uber.org/zap"
)

const (
	provider = "consul"
)

func InitConfig(filePath string) {
	v := viper.New()
	v.SetConfigFile(filePath)
	if err := v.ReadInConfig(); err != nil {
		zap.S().Panic("read config failed", zap.Error(err))
	}
	if err := v.Unmarshal(global.ServerConfig); err != nil {
		zap.S().Panic("unmarshal config failed", zap.Error(err))
	}

	// 从配置中心读取配置
	err := v.AddRemoteProvider(provider,
		fmt.Sprintf("%s:%d", global.ServerConfig.ConsulConfig.Host, global.ServerConfig.ConsulConfig.Port),
		global.ServerConfig.ConfigPath)
	if err != nil {
		zap.S().Warn("add remote provider failed")
		return
	}

	v.SetConfigType("YAML")

	if err = v.ReadRemoteConfig(); err != nil {
		zap.S().Warn("read remote config failed")
		return
	}

	if err = v.Unmarshal(global.ServerConfig); err != nil {
		zap.S().Panic("unmarshal remote config failed")
	}
}
