package controllers

// ResponseFmtData 返回数据格式
type ResponseFmtData struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}
