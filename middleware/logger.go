package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
)

func LoggerServe(app *iris.Application) {

	requestLogger := logger.New(
		logger.Config{
			Status:             true,
			IP:                 true,
			Method:             true,
			Path:               true,
			Query:              false,
			MessageContextKeys: []string{"[Logger_MESSAGE]"},
			MessageHeaderKeys:  []string{"[User-Agent]"},
		},
	)
	app.Use(requestLogger)
}
