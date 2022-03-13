package model

import (
	"time"
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

	return HandleDBData(result, user.StudentId+user.ManageId)
}

// Register postRegister
func (s *SignModel) Register(data mvc_struct.SignOfRegister) BaseModelFmtData {
	user := db.User{
		Email:     data.Email,
		Phone:     data.Phone,
		Password:  data.Password,
		Avatar:    data.Avatar,
		Identity:  "",
		CreateAt:  time.Now().Unix(),
		ManageId:  data.ManageId,
		StudentId: data.StudentId,
	}
	result := db.Mysql.Create(&user)
	return HandleDBData(result, nil)
}
