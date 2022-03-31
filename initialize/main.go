package initialize

import (
	"xupt-scholarship/config"
	"xupt-scholarship/db"
	"xupt-scholarship/global"
)

// Init 初始化配置信息
func Init() {
	global.Settings = config.GetProdGlobalConfig()
	db.Mysql = db.NewMysql()
	db.Redis = db.NewRedis()
	db.InitSessionDB()
}
