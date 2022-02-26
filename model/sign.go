package model

import (
	"time"
	"xupt-scholarship/db"
	"xupt-scholarship/mvc_struct"
)

type SignModel struct {
}

// Login postLogin
func (s *SignModel) Login(data mvc_struct.LoginForm) {

}

// Register postRegister
func (s *SignModel) Register(data mvc_struct.RegisterForm) DataBaseFmtData {
	user := db.User{
		Email:     data.Email,
		Phone:     data.Phone,
		Password:  data.Password,
		Avatar:    data.Avatar,
		Identity:  "",
		CreateAt:  time.Now().Unix(),
		ManageId:  data.ManagerId,
		StudentId: data.StudentId,
	}
	result := db.Mysql.Create(&user)
	res := DataBaseFmtData{
		Message: "注册成功",
		Data:    nil,
		Error:   result.Error,
	}
	if result.Error != nil {
		res.Message = "注册失败"
	}
	return res
}
