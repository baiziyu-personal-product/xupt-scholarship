package db

import "gorm.io/gorm"

var (
	Redis *RedisClient
	Mysql *gorm.DB
)
