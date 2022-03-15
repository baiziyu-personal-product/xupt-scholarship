package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"xupt-scholarship/global"
	"xupt-scholarship/model"
)

// GetRequestParams 获取请求参数
func GetRequestParams(ctx iris.Context, data interface{}) {
	var params interface{}
	if ctxErr := ctx.ReadJSON(&params); ctxErr != nil {
		panic(ctxErr)
	}
	value, err := json.Marshal(params)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return
	}
	if err := json.Unmarshal(value, data); err != nil {
		panic(err)
	}
}

// HandleControllerRes 处理Controller返回值
func HandleControllerRes(modelData model.BaseModelFmtData, message string) BaseControllerFmtData {
	data := BaseControllerFmtData{
		Message: message,
		Code:    modelData.Code,
		Data:    modelData.Data,
	}
	if modelData.Code == global.SuccessCode {
		data.Message += "成功"
	} else {
		data.Message += "失败"
	}
	return data
}
