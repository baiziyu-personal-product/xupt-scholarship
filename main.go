package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"xupt-scholarship/middleware"
)

func main() {
	app := iris.New()
	middleware.UseMiddleWare(app)
	app.Get("/", func(context *context.Context) {
		context.HTML("<h1>Hello World!!!</h1>")
	})
	app.Listen(":8096")
}
