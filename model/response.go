package model

type ResponseBase struct {
	Code    int
	Message string
	Data    interface{}
}

func ResponseResult(data interface{}, code int, message string) ResponseBase {
	return ResponseBase{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func ResponseError(code int, message string) ResponseBase {
	return ResponseBase{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}
