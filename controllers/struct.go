package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"xupt-scholarship/model"
)

// BaseControllerFmtData 返回数据格式
type BaseControllerFmtData struct {
	Message string      `json:"message" default:"成功"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

// BaseController 基本配置
type BaseController struct {
	Ctx     iris.Context
	Session *sessions.Session
}

type AccessController struct {
	User model.LoginUserInfo
}
