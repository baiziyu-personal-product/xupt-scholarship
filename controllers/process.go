package controllers

import "github.com/kataras/iris/v12/mvc"

type Process struct {
	BaseController
}

func ProcessMvc(app *mvc.Application) {
	app.Handle(new(Process))
}

func (p *Process) Get() ResponseFmtData {
	return ResponseFmtData{
		Code:    1,
		Message: "成功拉取申请流程",
		Data:    nil,
	}
}

func (p *Process) Post() ResponseFmtData {
	return ResponseFmtData{
		Message: "处理成功",
		Code:    1,
		Data:    nil,
	}
}

func (p *Process) GetSearch() ResponseFmtData {
	return ResponseFmtData{
		Message: "成功",
		Code:    0,
		Data:    nil,
	}
}
