package model

import (
	"encoding/json"
	"xupt-scholarship/db"
	"xupt-scholarship/mvc_struct"
)

type UserModel struct {
}

type UserModelInterface interface {
	GetUserList() BaseModelFmtData
	GetUser(email string) BaseModelFmtData
	UpdateUser(email string, userIdentity string, info mvc_struct.UpdateUserInfo) BaseModelFmtData
	CreateStudentByList(list []mvc_struct.StudentItem) BaseModelFmtData
	GetStudent(userId string) BaseModelFmtData
	GetUserBaseInfo(userId string) BaseModelFmtData
}

// >>>>>>>>>>>>>>> struct <<<<<<<<<<<<<<<<//

type LoginUserInfo struct {
	UserBaseInfo
	Phone    string                 `json:"phone"`
	Identity string                 `json:"identity"`
	UserId   string                 `json:"user_id"`
	Student  mvc_struct.StudentInfo `json:"student"`
	Manager  mvc_struct.ManagerInfo `json:"manager"`
}

// >>>>>>>>>>>>>>> interface <<<<<<<<<<<<<//

type UserBaseInfo struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
}

func (u *UserModel) GetUserList() BaseModelFmtData {
	var users []db.User
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

func (u *UserModel) GetUserBaseInfo(userId string) BaseModelFmtData {
	var user db.User
	result := db.Mysql.Where("user_id = ?", userId).First(&user)
	return HandleDBData(result, UserBaseInfo{
		Name:   user.Name,
		Email:  user.Email,
		Avatar: user.Avatar,
	})
}

func (u *UserModel) GetStudent(userId string) BaseModelFmtData {
	var user db.User
	result := db.Mysql.Where("user_id = ?", userId).First(&user)
	var info mvc_struct.StudentInfo
	json.Unmarshal(user.Info, &info)
	student := mvc_struct.Student{
		Info:   info,
		Course: user.Course,
		Name:   user.Name,
		Avatar: user.Avatar,
		Phone:  user.Phone,
		UserId: user.UserId,
	}
	return HandleDBData(result, student)
}

func (u *UserModel) GetUser(email string) BaseModelFmtData {
	var user db.User
	result := db.Mysql.Where("email = ?", email).First(&user)

	var userInfo LoginUserInfo
	userInfo = LoginUserInfo{
		UserBaseInfo: UserBaseInfo{
			Name:   user.Name,
			Email:  user.Email,
			Avatar: user.Avatar,
		},
		Phone:    user.Phone,
		Identity: user.Identity,
		UserId:   user.UserId,
	}

	if user.Identity == "manager" {
		var manager mvc_struct.ManagerInfo
		json.Unmarshal(user.Info, &manager)
		userInfo.Manager = manager
	} else {
		var student mvc_struct.StudentInfo
		json.Unmarshal(user.Info, &student)
		userInfo.Student = student
	}
	return HandleDBData(result, userInfo)
}

func (u *UserModel) UpdateUser(email string, userIdentity string, info mvc_struct.UpdateUserInfo) BaseModelFmtData {
	var user db.User
	var jsonInfo []byte
	if userIdentity == "manager" {
		jsonInfo, _ = json.Marshal(info.Manager)
	} else {
		jsonInfo, _ = json.Marshal(info.Student)
	}
	infoMap := map[string]interface{}{
		"avatar":  info.Avatar,
		"phone":   info.Phone,
		"user_id": info.UserId,
		"name":    info.Name,
		"info":    jsonInfo,
	}
	result := db.Mysql.Model(&user).Where("email = ?", email).Updates(infoMap)
	return HandleDBData(result, user.ID)
}

func (u *UserModel) CreateStudentByList(list []mvc_struct.StudentItem) BaseModelFmtData {
	var users []db.User
	for _, temp := range list {
		info := mvc_struct.StudentInfo{
			Professional: temp.Professional,
			Grade:        temp.Grade,
			Class:        temp.Class,
			Type:         temp.Type,
		}
		jsonInfo, _ := json.Marshal(info)
		student := db.User{
			Name:     temp.Name,
			Email:    temp.Email,
			Phone:    temp.Phone,
			Password: temp.Password,
			Identity: "student",
			UserId:   temp.StudentId,
			Info:     jsonInfo,
			Course:   temp.CourseCredit,
		}
		users = append(users, student)
	}

	result := db.Mysql.Create(&users)
	var res []string
	for _, user := range users {
		res = append(res, user.UserId)
	}
	return HandleDBData(result, res)
}
