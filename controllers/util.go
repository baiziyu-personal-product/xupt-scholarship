package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
)

func DeepCopy(value interface{}) interface{} {
	if valueMap, ok := value.(map[string]interface{}); ok {
		newMap := make(map[string]interface{})
		for k, v := range valueMap {
			newMap[k] = DeepCopy(v)
		}

		return newMap
	} else if valueSlice, ok := value.([]interface{}); ok {
		newSlice := make([]interface{}, len(valueSlice))
		for k, v := range valueSlice {
			newSlice[k] = DeepCopy(v)
		}

		return newSlice
	}

	return value
}

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
