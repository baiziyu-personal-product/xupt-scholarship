package controllers

import (
	"bytes"
	"github.com/dchest/captcha"
	"github.com/kataras/iris/v12/mvc"
	"net/http"
	"strings"
	"time"
)

type Captcha struct {
	BaseController
}

func CaptchaMVC(app *mvc.Application) {
	app.Handle(new(Captcha))
}

func (C *Captcha) GetBy(imgPath string) {
	w, r := C.Ctx.ResponseWriter(), C.Ctx.Request()
	point := strings.Index(imgPath, ".")
	var content bytes.Buffer
	w.Header().Set("Content-Type", "image/png")
	code := imgPath[:point]
	captcha.WriteImage(&content, code, captcha.StdWidth, captcha.StdHeight)
	http.ServeContent(w, r, imgPath, time.Time{}, bytes.NewReader(content.Bytes()))
}

type VerifyData struct {
	Code      string `json:"code"`
	InputCode string `json:"input_code"`
}

// PostVerify 判断验证码是否正确
func (C *Captcha) PostVerify() ResponseFmtData {
	var data VerifyData
	GetRequestParams(C.Ctx, &data)
	return ResponseFmtData{
		Message: "verifyed",
		Code:    0,
		Data:    captcha.VerifyString(data.Code, data.InputCode),
	}
}

func (C *Captcha) Get() ResponseFmtData {
	captchaId := captcha.New()
	return ResponseFmtData{
		Message: "success",
		Code:    0,
		Data:    captchaId,
	}
}

func (C *Captcha) GetReload() ResponseFmtData {
	code := C.Ctx.URLParam("code")
	return ResponseFmtData{
		Message: "success",
		Code:    0,
		Data:    captcha.Reload(code),
	}
}
