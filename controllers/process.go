package controllers

import (
	"github.com/kataras/iris/v12/mvc"
	"xupt-scholarship/mvc_struct"
)

type ProcessMVC struct {
	BaseController
}

func UseProcessMVC(app *mvc.Application) {
	app.Handle(new(ProcessMVC))
}

func (p *ProcessMVC) GetBy(processId int) BaseControllerFmtData {
	return BaseControllerFmtData{
		Code:    1,
		Message: "成功拉取申请流程",
		Data:    nil,
	}
}

func (p *ProcessMVC) Post() BaseControllerFmtData {
	var processInfo mvc_struct.ProcessReqData
	GetRequestParams(p.Ctx, &processInfo)
	return BaseControllerFmtData{
		Message: "处理成功",
		Code:    1,
		Data:    nil,
	}
}

func (p *ProcessMVC) PutBy(processId int) BaseControllerFmtData {
	var processInfo mvc_struct.ProcessReqData
	GetRequestParams(p.Ctx, &processInfo)
	return BaseControllerFmtData{
		Message: "处理成功",
		Code:    1,
		Data:    nil,
	}
}
