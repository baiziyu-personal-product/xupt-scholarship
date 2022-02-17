package db

import (
	"gorm.io/gorm"
	"xupt-scholarship/common"
	"xupt-scholarship/global"
)

// UseMysql 连接Mysql
func UseMysql() *gorm.DB {

	return NewMysql()
}

// UseRedis 连接Redis
func UseRedis() *RedisClient {

	redisConfig := global.Settings.RedisConfig
	redisOpt := common.RedisConnOpt{
		Enable: true,
		Host:   redisConfig.Host,
		Port:   redisConfig.Port,
		TTL:    240,
	}
	redisCli := NewRedis(redisOpt)

	return redisCli
}
