package common

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

// BaseResponse 基本Controller数据返回格式
type BaseResponse struct {
	Message string      `json:"message" default:"成功"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

// BaseMVC 基本Controller(包含用户Session)
type BaseMVC struct {
	Ctx     iris.Context
	Session *sessions.Session
}
