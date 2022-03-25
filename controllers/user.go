package controllers

import (
	"github.com/kataras/iris/v12/mvc"
	"xupt-scholarship/model"
	"xupt-scholarship/mvc_struct"
)

type User struct {
	BaseController
}

var UserModel model.UserModel

func UseUserMVC(app *mvc.Application) {

	app.Register(userSession.Start).Handle(new(User))
}

func (u *User) Get() BaseControllerFmtData {
	email := u.Session.GetString(sessionId)
	result := UserModel.GetUser(email)
	return HandleControllerRes(result, "获取用户信息")
}

func (u *User) GetList() BaseControllerFmtData {
	res := UserModel.GetUserList()
	return HandleControllerRes(res, "获取用户列表")
}

func (u *User) PostUpdateUserInfo() BaseControllerFmtData {
	var info mvc_struct.UpdateUserInfo
	email := u.Session.GetString(sessionId)
	GetRequestParams(u.Ctx, &info)
	res := UserModel.UpdateUser(email, info)
	return HandleControllerRes(res, "更新用户信息")
}
