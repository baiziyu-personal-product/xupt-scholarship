package controllers

import "github.com/kataras/iris/v12/mvc"

type ApplyMVC struct {
	BaseController
}

func UseApplyMVC(app *mvc.Application) {
	app.Handle(new(ApplyMVC))
}

func (a *ApplyMVC) GetBy(applyId int64) ResponseFmtData {
	return ResponseFmtData{
		Message: "",
		Code:    1,
		Data:    applyId,
	}
}

func (a *ApplyMVC) Post() ResponseFmtData {
	return ResponseFmtData{
		Message: "",
		Code:    1,
		Data:    nil,
	}
}

func (a *ApplyMVC) PutBy(applyId int64) ResponseFmtData {
	return ResponseFmtData{
		Message: "",
		Code:    0,
		Data:    applyId,
	}
}
