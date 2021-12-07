package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()
	app.Get("/", index)
	app.Listen(":8096")
}

func index(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"Message": "Hello, Welcome to Iris Application!!!",
	})
}
