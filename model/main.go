package model

import (
	"gorm.io/gorm"
	"xupt-scholarship/global"
)

type BaseModelFmtData struct {
	Message string
	Data    interface{}
	Error   error
	Code    int
}

var (
	successData = BaseModelFmtData{
		Message: "[✅]Database processing complete~",
		Data:    nil,
		Error:   nil,
		Code:    global.SuccessCode,
	}
	errorData = BaseModelFmtData{
		Message: "[❎]Database processing failed~",
		Data:    nil,
		Error:   nil,
		Code:    global.ErrorCode,
	}
)

// HandleDBData 处理数据库错误
func HandleDBData(result *gorm.DB, data interface{}) BaseModelFmtData {
	if result.Error != nil {
		errorData.Error = result.Error
		return errorData
	}
	successData.Data = data
	return successData
}
