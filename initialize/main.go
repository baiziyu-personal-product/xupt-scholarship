package initialize

import (
	"xupt-scholarship/config"
	"xupt-scholarship/global"
)

// Init 初始化配置信息
func Init() {
	global.Settings = config.GetGlobalConfig()
}
