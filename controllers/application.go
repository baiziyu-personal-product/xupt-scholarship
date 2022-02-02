package controllers

import "github.com/kataras/iris/v12/mvc"

type Application struct {
	BaseController
}

func ApplicationMVC(app *mvc.Application) {
	app.Handle(new(Application))
}

func (a *Application) Get() ResponseFmtData {
	return ResponseFmtData{
		Message: "",
		Code:    1,
		Data:    nil,
	}
}

func (a *Application) Post() ResponseFmtData {
	return ResponseFmtData{
		Message: "",
		Code:    1,
		Data:    nil,
	}
}
