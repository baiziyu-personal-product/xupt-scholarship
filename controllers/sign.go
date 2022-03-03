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
	signModel := SignModel.Login(data)
	res := ResponseFmtData{
		Message: "登录失败",
		Code:    1,
		Data:    nil,
	}
	if signModel.Error == nil {
		res.Message = "登录成功"
		res.Code = 1
		res.Data = middleware.GenerateToken(data.Email)
		s.Session.Set("xupt_session_id", data.Email)
	}
	return res
}

// PostRegister 注册
func (s *SignMvc) PostRegister() ResponseFmtData {
	var data mvc_struct.RegisterForm
	GetRequestParams(s.Ctx, &data)
	signModel := SignModel.Register(data)
	res := ResponseFmtData{
		Message: signModel.Message,
		Code:    0,
		Data:    nil,
	}

	if signModel.Error == nil {
		s.Session.Set("xupt_session_id", data.Email)
		res.Data = middleware.GenerateToken(data.Email)
		res.Code = 1
	}
	return res
}

// PostForget 忘记密码
func (s *SignMvc) PostForget() ResponseFmtData {
	var data mvc_struct.ForgetForm
	GetRequestParams(s.Ctx, &data)
	return ResponseFmtData{
		Message: "登录成功",
		Code:    1,
		Data:    middleware.GenerateToken(data.Email),
	}
}
