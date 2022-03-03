package controllers

import (
	"github.com/kataras/iris/v12/mvc"
)

type User struct {
	BaseController
}

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
