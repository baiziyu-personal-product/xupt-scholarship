package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12/mvc"
	"xupt-scholarship/model"
	"xupt-scholarship/mvc_struct"
)

type ProcessMVC struct {
	BaseController
}

func UseProcessMVC(app *mvc.Application) {
	app.Register(userSession.Start).Handle(new(ProcessMVC))
}

func (p *ProcessMVC) GetBy(processId int) BaseControllerFmtData {
	return BaseControllerFmtData{
		Code:    1,
		Message: "成功拉取申请流程",
		Data:    nil,
	}
}

func (p *ProcessMVC) Post() BaseControllerFmtData {
	var processInfo mvc_struct.ProcessFormData
	GetRequestParams(p.Ctx, &processInfo)
	email := p.Session.GetString(sessionId)
	user := UserModel.GetUser(email).Data.(model.LoginUserInfo)
	fmt.Println(user)
	fmt.Println("email" + email)
	return BaseControllerFmtData{
		Message: "处理成功",
		Code:    1,
		Data:    nil,
	}
}

func (p *ProcessMVC) PutBy(processId int) BaseControllerFmtData {
	var processInfo mvc_struct.ProcessFormData
	GetRequestParams(p.Ctx, &processInfo)
	return BaseControllerFmtData{
		Message: "处理成功",
		Code:    1,
		Data:    nil,
	}
}
