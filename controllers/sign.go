package controllers

import (
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"time"
	"xupt-scholarship/middleware"
)

type Sign struct {
	BaseController
}

func SignMvc(app *mvc.Application) {
	userSession := sessions.New(sessions.Config{
		Cookie:  "xupt_session_id",
		Expires: 24 * 2 * time.Hour,
	})
	app.Register(userSession.Start).Handle(new(Sign))
}

type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Remember bool   `json:"remember"`
}

// PostLogin 登录
func (s *Sign) PostLogin() ResponseFmtData {
	var data LoginForm
	GetRequestParams(s.Ctx, &data)
	s.Session.Set("xupt_session_id", data.Email)
	return ResponseFmtData{
		Message: "登录成功",
		Code:    1,
		Data:    middleware.GenerateToken(data.Email),
	}
}

func (s *Sign) PostForget() ResponseFmtData {
	var data LoginForm
	GetRequestParams(s.Ctx, &data)
	s.Session.Set("xupt_session_id", data.Email)
	return ResponseFmtData{
		Message: "登录成功",
		Code:    1,
		Data:    middleware.GenerateToken(data.Email),
	}
}
