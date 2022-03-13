package model

import (
	"xupt-scholarship/db"
)

var users []db.User
var user db.User

type UserModel struct {
}

type UserBaseInfo struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
}

func (u *UserModel) GetUserList() BaseModelFmtData {
	result := db.Mysql.Find(&users)
	var userList []UserBaseInfo
	for _, user := range users {
		if user.Email != "" {
			userList = append(userList, UserBaseInfo{
				Name:   user.Name,
				Email:  user.Email,
				Avatar: user.Avatar,
			})
		}
	}
	return HandleDBData(result, userList)
}

type LoginUserInfo struct {
	UserBaseInfo
	Phone     string `json:"phone"`
	Identity  string `json:"identity"`
	StudentId string `json:"student_id"`
	ManagerId string `json:"manager_id"`
}

func (u *UserModel) GetUser(email string) BaseModelFmtData {
	result := db.Mysql.Where(&db.User{Email: email}).First(&user)
	var userInfo LoginUserInfo
	userInfo = LoginUserInfo{
		UserBaseInfo: UserBaseInfo{
			Name:   user.Name,
			Email:  user.Email,
			Avatar: user.Avatar,
		},
		Phone:     user.Phone,
		Identity:  user.Identity,
		StudentId: user.StudentId,
		ManagerId: user.ManageId,
	}
	return HandleDBData(result, userInfo)
}
