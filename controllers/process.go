package controllers

import "github.com/kataras/iris/v12/mvc"

type ProcessMVC struct {
	BaseController
}

func UseProcessMVC(app *mvc.Application) {
	app.Handle(new(ProcessMVC))
}

func (p *ProcessMVC) Get() ResponseFmtData {
	return ResponseFmtData{
		Code:    1,
		Message: "成功拉取申请流程",
		Data:    nil,
	}
}

func (p *ProcessMVC) Post() ResponseFmtData {
	return ResponseFmtData{
		Message: "处理成功",
		Code:    1,
		Data:    nil,
	}
}

func (p *ProcessMVC) GetSearch() ResponseFmtData {
	return ResponseFmtData{
		Message: "成功",
		Code:    0,
		Data:    nil,
	}
}
