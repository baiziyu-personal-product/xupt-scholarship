package middleware

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

func CorsServer(app *iris.Application) {
	corsOpt := cors.Options{
		AllowedOrigins:   []string{"*"},
		MaxAge:           600,
		AllowCredentials: true,
	}
	app.Use(cors.New(corsOpt))
}
