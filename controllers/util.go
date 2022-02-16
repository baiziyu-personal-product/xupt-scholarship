package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
)

// GetRequestParams 获取请求参数
func GetRequestParams(ctx iris.Context, data interface{}) {
	var params RequestFmtData
	ctx.ReadJSON(&params)
	value, err := json.Marshal(params.Data)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return
	}
	json.Unmarshal(value, data)
}
