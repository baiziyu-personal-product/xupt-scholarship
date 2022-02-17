package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"xupt-scholarship/global"
)

func NewMysql() *gorm.DB {
	mysqlConfig := global.Settings.MysqlConfig
	dsn := mysqlConfig.Name + ":" + mysqlConfig.Password + "@tcp(" + mysqlConfig.Host + ":" + strconv.Itoa(mysqlConfig.Port) + ")/" + mysqlConfig.DBName + "?charset=utf8&parseTime=True&loc=Local"
	mysqlCli, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return mysqlCli
}
