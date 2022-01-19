package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/mlogclub/simple"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"xupt-scholarship/db"
	"xupt-scholarship/global"
	"xupt-scholarship/middleware"
)

func Router() {
	app := iris.New()
	app.Logger().SetLevel("warn")
	middleware.UseMiddleWare(app)
	app.AllowMethods(iris.MethodOptions)

	app.OnAnyErrorCode(func(ctx *context.Context) {
		path := ctx.Path()
		var err error
		if strings.Contains(path, "/api/admin") {
			_, err = ctx.JSON(simple.JsonErrorCode(ctx.GetStatusCode(), "Http Error"))
		}

		if err != nil {
			logrus.Error(err)
		}
	})

	app.Any("/", func(c *context.Context) {
		_, _ = c.HTML("<h1>Powered by xupt-scholarshipt@baiziyu-fe - | ©2022</h1>")
	})

	app.Get("/", func(context *context.Context) {
		db.Start()
		context.HTML("<h1>Hello World!!!</h1>")
	})

	server := &http.Server{Addr: ":" + strconv.Itoa(global.Settings.Port)}
	handleSignal(server)
	err := app.Run(iris.Server(server), iris.WithConfiguration(iris.Configuration{
		DisableStartupLog:                 false,
		DisableInterruptHandler:           false,
		DisablePathCorrection:             false,
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		EnableOptimizations:               true,
		TimeFormat:                        "2006-01-02 15:04:05",
		Charset:                           "UTF-8",
	}))

	if err != nil {
		logrus.Error(err)
		os.Exit(-1)
	}

}

func handleSignal(server *http.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		s := <-c
		logrus.Infof("got signal [%s], exiting now!!!", s)
		if err := server.Close(); err != nil {
			logrus.Errorf("server close failed: " + err.Error())
		}
		simple.CloseDB()
		logrus.Infof("Exited")
		os.Exit(0)
	}()
}
