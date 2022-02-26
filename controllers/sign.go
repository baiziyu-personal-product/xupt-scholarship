package controllers

import (
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"time"
	"xupt-scholarship/middleware"
	"xupt-scholarship/model"
	"xupt-scholarship/mvc_struct"
)

type SignMvc struct {
	BaseController
}

var SignModel model.SignModel

func UseSignMvc(app *mvc.Application) {
	userSession := sessions.New(sessions.Config{
		Cookie:  "xupt_session_id",
		Expires: 24 * 2 * time.Hour,
	})
	app.Register(userSession.Start).Handle(new(SignMvc))
}

// PostLogin 登录
func (s *SignMvc) PostLogin() ResponseFmtData {
	var data mvc_struct.LoginForm
	GetRequestParams(s.Ctx, &data)
	s.Session.Set("xupt_session_id", data.Email)
	return ResponseFmtData{
		Message: "登录成功",
		Code:    1,
		Data:    middleware.GenerateToken(data.Email),
	}
}

// PostRegister 注册
func (s *SignMvc) PostRegister() ResponseFmtData {
	var data mvc_struct.RegisterForm
	GetRequestParams(s.Ctx, &data)
	model := SignModel.Register(data)
	res := ResponseFmtData{
		Message: model.Message,
		Code:    0,
		Data:    nil,
	}

	if model.Error == nil {
		s.Session.Set("xupt_session_id", data.Email)
		res.Data = middleware.GenerateToken(data.Email)
		res.Code = 1
	}
	return res
}

type ForgetForm struct {
	Phone     string `json:"phone"`
	StudentId string `json:"student_id"`
	ManagerId string `json:"manager_id"`
	Email     string `json:"email"`
}

// PostForget 忘记密码
func (s *SignMvc) PostForget() ResponseFmtData {
	var data mvc_struct.LoginForm
	GetRequestParams(s.Ctx, &data)
	return ResponseFmtData{
		Message: "登录成功",
		Code:    1,
		Data:    middleware.GenerateToken(data.Email),
	}
}
