package initialize

import (
	"github.com/fatih/color"
	"github.com/spf13/viper"
	"xupt-scholarship/config"
	"xupt-scholarship/global"
)

func InitServeConfig() {
	v := viper.New()
	v.SetConfigFile("./dev.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	serverConfig := config.ServerConfig{}

	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	global.Settings = serverConfig

	color.Blue("[[Serve has been start🎉]]", global.Settings.LogsAddr)
}
