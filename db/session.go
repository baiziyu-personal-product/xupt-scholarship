package db

import (
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
	"strconv"
	"xupt-scholarship/global"
)

func InitSessionDB() {
	redisConf := global.Settings.RedisConfig
	db := redis.New(redis.Config{
		Network:   "tcp",
		Addr:      redisConf.Host + ":" + strconv.Itoa(redisConf.Port),
		MaxActive: 100,
		Username:  "",
		Password:  "",
		Database:  "",
		Prefix:    "sr_",
		Driver:    redis.GoRedis(),
	})
	global.UserSession.UseDatabase(db)
}
