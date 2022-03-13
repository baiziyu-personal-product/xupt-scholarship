package controllers

import (
	"github.com/kataras/iris/v12/mvc"
	"xupt-scholarship/model"
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
