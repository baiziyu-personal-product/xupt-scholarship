package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"
	"github.com/mlogclub/simple"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
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

	// MVC
	mvc.Configure(app.Party("/sign"), UseSignMvc)
	// 配置验证码 相关路由
	mvc.Configure(app.Party("/captcha"), UseCaptchaMVC)

	// 需要权限验证
	// 上传文件
	mvc.Configure(app.Party("/upload", middleware.JwtVerify()), UseUploadMVC)
	// 配置 用户 下相关路由
	mvc.Configure(app.Party("/user", middleware.JwtVerify()), UseUserMVC)
	// 配置 进度管理 下相关路由
	mvc.Configure(app.Party("/process", middleware.JwtVerify()), UseProcessMVC)
	// 配置 申请 相关路由
	mvc.Configure(app.Party("/apply", middleware.JwtVerify()), UseApplyMVC)

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
