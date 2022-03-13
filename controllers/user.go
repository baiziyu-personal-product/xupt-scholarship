package controllers

import (
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"time"
	"xupt-scholarship/model"
)

type User struct {
	BaseController
}

var UserModel model.UserModel

var (
	SessionId   = "xupt_session_id"
	UserSession = sessions.New(sessions.Config{
		Cookie:  SessionId,
		Expires: 24 * 2 * time.Hour,
	})
)

func UseUserMVC(app *mvc.Application) {

	app.Register(UserSession.Start).Handle(new(User))
}

func (u *User) Get() ResponseFmtData {
	email := u.Session.GetString(SessionId)
	result := UserModel.GetUser(email)
	if result.Error != nil {
		return ResponseFmtData{
			Message: result.Message,
			Code:    0,
			Data:    result.Error,
		}
	}
	return ResponseFmtData{
		Message: "成功获取用户信息",
		Code:    1,
		Data:    result.Data,
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
