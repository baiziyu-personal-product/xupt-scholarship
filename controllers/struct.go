package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

// ResponseFmtData 返回数据格式
type ResponseFmtData struct {
	Message string      `json:"message" default:"成功"`
	Code    int         `json:"code" default:"1"`
	Data    interface{} `json:"data"`
}

type RequestFmtData struct {
	Data interface{} `json:"data"`
}

// BaseController 基本配置
type BaseController struct {
	Ctx     iris.Context
	Session *sessions.Session
}
