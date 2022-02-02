package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"xupt-scholarship/global"
)

// ConnectDB 连接数据库
func ConnectDB() *gorm.DB {
	mysqlConfig := global.Settings.MysqlConfig
	dsn := mysqlConfig.Name + ":" + mysqlConfig.Password + "@tcp(" + mysqlConfig.Host + ":" + strconv.Itoa(mysqlConfig.Port) + ")/" + mysqlConfig.DBName + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}
