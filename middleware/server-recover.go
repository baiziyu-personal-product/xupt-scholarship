package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
)

func RecoverServe(app *iris.Application) {
	app.Use(recover.New())
}
