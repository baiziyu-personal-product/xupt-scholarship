package controllers

import (
	"github.com/kataras/iris/v12/mvc"
	"xupt-scholarship/model"
	"xupt-scholarship/mvc_struct"
)

type processMVC struct {
	BaseController
}

var ProcessModel model.ProcessModel

func UseProcessMVC(app *mvc.Application) {
	app.Register(userSession.Start).Handle(new(processMVC))
}

func (p *processMVC) GetBy(processId int) BaseControllerFmtData {
	return BaseControllerFmtData{
		Code:    1,
		Message: "成功拉取申请流程",
		Data:    nil,
	}
}

func (p *processMVC) Post() BaseControllerFmtData {
	var processInfo mvc_struct.ProcessFormData
	GetRequestParams(p.Ctx, &processInfo)
	email := p.Session.GetString(sessionId)
	user := UserModel.GetUser(email).Data.(model.LoginUserInfo)
	ProcessModel.CreateProcess(processInfo, user.UserId)
	return BaseControllerFmtData{
		Message: "处理成功",
		Code:    1,
		Data:    nil,
	}
}

func (p *processMVC) PutBy(processId int) BaseControllerFmtData {
	var processInfo mvc_struct.ProcessFormData
	GetRequestParams(p.Ctx, &processInfo)
	return BaseControllerFmtData{
		Message: "处理成功",
		Code:    1,
		Data:    nil,
	}
}
