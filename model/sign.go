package model

import (
	"xupt-scholarship/db"
	"xupt-scholarship/mvc_struct"
)

type SignModel struct {
}

// Login postLogin
func (s *SignModel) Login(data mvc_struct.SignOfLogin) BaseModelFmtData {
	user := db.User{}
	result := db.Mysql.Where(&db.User{
		Email:    data.Email,
		Password: data.Password,
	}).First(&user)

	return HandleDBData(result, user.UserId)
}

// Register postRegister
func (s *SignModel) Register(data mvc_struct.SignOfRegister) BaseModelFmtData {
	user := db.User{
		Email:    data.Email,
		Password: data.Password,
		Identity: data.Identity,
	}
	result := db.Mysql.Create(&user)
	return HandleDBData(result, nil)
}
