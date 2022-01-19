package initialize

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/viper"
	"xupt-scholarship/config"
	"xupt-scholarship/global"
	"xupt-scholarship/utils"
)

func InitServeConfig() {
	v := viper.New()
	v.SetConfigFile("./dev.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	serverConfig := config.ServerConfig{}
	fmt.Println(v.Get("mysql"))
	serveErr := v.Unmarshal(&serverConfig)
	mysqlErr := v.UnmarshalKey("mysql", &serverConfig.MysqlConfig)
	redisErr := v.UnmarshalKey("redis", &serverConfig.RedisConfig)
	utils.DealErrors(serveErr, mysqlErr, redisErr)
	global.Settings = serverConfig
	fmt.Println(serverConfig)
	color.Blue("[[Serve has been start🎉]]", global.Settings.LogsAddr)
}
