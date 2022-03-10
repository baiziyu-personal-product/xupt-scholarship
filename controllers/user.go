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
	app.Handle(new(User))
}

func (u *User) Get() ResponseFmtData {
	return ResponseFmtData{
		Message: "成功获取用户信息",
		Code:    1,
		Data:    nil,
	}
}

func (u *User) GetList() ResponseFmtData {
	res := UserModel.GetUserList()
	return ResponseFmtData{
		Message: res.Message,
		Code:    1,
		Data:    res.Data,
	}
}
