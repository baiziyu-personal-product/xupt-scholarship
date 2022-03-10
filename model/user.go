package model

import (
	"xupt-scholarship/db"
)

var users []db.User
var user db.User

type UserModel struct {
}

type UserListItem struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
}

func (u *UserModel) GetUserList() DataBaseFmtData {
	result := db.Mysql.Find(&users)
	var userList []UserListItem
	for _, user := range users {
		if user.Email != "" {
			userList = append(userList, UserListItem{
				Name:   user.Name,
				Email:  user.Email,
				Avatar: user.Avatar,
			})
		}
	}
	return DataBaseFmtData{
		Message: "success",
		Data:    userList,
		Error:   result.Error,
	}
}
