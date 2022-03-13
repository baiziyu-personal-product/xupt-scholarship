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
	ctx.ReadJSON(&params)
	value, err := json.Marshal(params)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return
	}
	json.Unmarshal(value, data)
}

// HandleControllerRes 处理Controller返回值
func HandleControllerRes(modelData model.BaseModelFmtData, message string) BaseControllerFmtData {
	data := BaseControllerFmtData{
		Message: message,
		Code:    modelData.Code,
		Data:    modelData.Error,
	}
	if modelData.Code == global.SuccessCode {
		data.Message += "成功"
	} else {
		data.Message += "失败"
	}
	return data
}
