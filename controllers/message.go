package controllers

import "github.com/kataras/iris/v12/mvc"

type MessageMvc struct {
	BaseController
}

func UseMessageMVC(app *mvc.Application) {
	app.Handle(new(MessageMvc))
}
