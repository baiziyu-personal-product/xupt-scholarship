package controllers

import (
	"github.com/kataras/iris/v12/mvc"
	"xupt-scholarship/model"
	"xupt-scholarship/mvc_struct"
)

type UserMvc struct {
	BaseController
}

type UserController interface {
	Get() BaseControllerFmtData
	GetList() BaseControllerFmtData
	PostUpdateUserInfo() BaseControllerFmtData
}

var UserModel model.UserModel

func UseUserMVC(app *mvc.Application) {

	app.Register(userSession.Start).Handle(new(UserMvc))
}

func (u *UserMvc) Get() BaseControllerFmtData {
	email := u.Session.GetString(sessionId)
	result := UserModel.GetUser(email)
	return HandleControllerRes(result, "获取用户信息")
}

func (u *UserMvc) GetList() BaseControllerFmtData {
	res := UserModel.GetUserList()
	return HandleControllerRes(res, "获取用户列表")
}

func (u *UserMvc) PostUpdateUserInfo() BaseControllerFmtData {
	var info mvc_struct.UpdateUserInfo
	email := u.Session.GetString(sessionId)
	GetRequestParams(u.Ctx, &info)
	res := UserModel.UpdateUser(email, info)
	return HandleControllerRes(res, "更新用户信息")
}
