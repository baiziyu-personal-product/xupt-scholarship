package controllers

import (
	"github.com/kataras/iris/v12/mvc"
	"xupt-scholarship/global"
	"xupt-scholarship/middleware"
	"xupt-scholarship/model"
	"xupt-scholarship/mvc_struct"
)

type SignMvc struct {
	BaseController
}

var SignModel model.SignModel

func UseSignMvc(app *mvc.Application) {
	app.Register(userSession.Start).Handle(new(SignMvc))
}

type SignController interface {
	PostLogin() BaseControllerFmtData
	PostRegister() BaseControllerFmtData
	GetOut() BaseControllerFmtData
	PostForget() BaseControllerFmtData
}

// PostLogin 登录
func (s *SignMvc) PostLogin() BaseControllerFmtData {
	var data mvc_struct.SignOfLogin
	GetRequestParams(s.Ctx, &data)
	signModel := SignModel.CheckUser(data)
	if signModel.Code == global.SuccessCode && signModel.Error == nil {
		signModel.Data = middleware.GenerateToken(data.Email)
		s.Session.Set(sessionId, data.Email)
	}
	return HandleControllerRes(signModel, "登录")
}

// PostRegister 注册
func (s *SignMvc) PostRegister() BaseControllerFmtData {
	var data mvc_struct.SignOfRegister
	GetRequestParams(s.Ctx, &data)
	signModel := SignModel.CreateUser(data)
	if signModel.Code == global.SuccessCode {
		signModel.Data = middleware.GenerateToken(data.Email)
		s.Session.Set(sessionId, data.Email)
	}
	return HandleControllerRes(signModel, "注册")
}

// GetOut 注销登录
func (s *SignMvc) GetOut() BaseControllerFmtData {
	s.Session.Delete(s.Session.ID())
	return HandleControllerRes(model.BaseModelFmtData{
		Message: "success",
		Data:    nil,
		Error:   nil,
		Code:    global.SuccessCode,
	}, "注销")
}

// PostForget 忘记密码
func (s *SignMvc) PostForget() BaseControllerFmtData {
	var data mvc_struct.SignOfForget
	GetRequestParams(s.Ctx, &data)
	return BaseControllerFmtData{
		Message: "登录成功",
		Code:    1,
		Data:    middleware.GenerateToken(data.Email),
	}
}
