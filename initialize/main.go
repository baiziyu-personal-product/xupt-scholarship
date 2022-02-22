package initialize

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"xupt-scholarship/config"
	"xupt-scholarship/global"
	"xupt-scholarship/utils"
)

// 最终方案-全兼容
func getCurrentAbPath() string {
	dir := getCurrentAbPathByExecutable()
	tmpDir, _ := filepath.EvalSymlinks(os.TempDir())
	if strings.Contains(dir, tmpDir) {
		return getCurrentAbPathByCaller()
	}
	return dir
}

// 获取当前执行文件绝对路径
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

func InitServeConfig() {
	v := viper.New()

	execpath := getCurrentAbPath()

	configfile := filepath.Join(execpath, "../dev.yaml")
	v.SetConfigFile(configfile)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	serverConfig := config.ServerConfig{}
	serveErr := v.Unmarshal(&serverConfig)
	mysqlErr := v.UnmarshalKey("mysql", &serverConfig.MysqlConfig)
	redisErr := v.UnmarshalKey("redis", &serverConfig.RedisConfig)
	utils.DealErrors(serveErr, mysqlErr, redisErr)
	global.Settings = serverConfig
}
