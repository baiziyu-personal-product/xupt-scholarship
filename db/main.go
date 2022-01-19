package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"xupt-scholarship/global"
)

func Start() {
	db := connectDB()
	add(db)
}

func add(db *gorm.DB) {
	student := Students{
		Name:       "1ssssdf",
		Email:      "1234sfd56",
		Phone:      "123456sd",
		Password:   "123456",
		AccessType: 0,
		Access:     "1,3,4",
	}
	fmt.Printf("%v", student)
	result := db.Table("students").Create(&student)

	if result.Error != nil {
		panic(result.Error)
	}
}

func connectDB() *gorm.DB {
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
