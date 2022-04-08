package model

import (
	"xupt-scholarship/db"
	"xupt-scholarship/mvc_struct"
)

type SignModel struct {
}

type SignModelInterface interface {
	CheckUser(data mvc_struct.SignOfLogin) BaseModelFmtData
	CreateUser(data mvc_struct.SignOfRegister) BaseModelFmtData
}

// >>>>>>>>>>>>>>>>>>> interface <<<<<<<<<<<<<<<<<<<<<<//

// CheckUser 登录
func (s *SignModel) CheckUser(data mvc_struct.SignOfLogin) BaseModelFmtData {
	user := db.User{}
	result := db.Mysql.Where(&db.User{
		Email:    data.Email,
		Password: data.Password,
	}).First(&user)

	return HandleDBData(result, user.UserId)
}

// CreateUser 注册用户
func (s *SignModel) CreateUser(data mvc_struct.SignOfRegister) BaseModelFmtData {
	user := db.User{
		Email:    data.Email,
		Password: data.Password,
		Identity: data.Identity,
	}
	result := db.Mysql.Create(&user)
	return HandleDBData(result, nil)
}
