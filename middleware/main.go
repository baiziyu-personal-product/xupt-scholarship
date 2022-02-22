package middleware

import "github.com/kataras/iris/v12"

func UseMiddleWare(app *iris.Application) {
	RecoverServe(app)
	LoggerServe(app)
	//JwtVerify(app)
}
