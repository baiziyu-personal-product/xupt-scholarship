package controllers

import (
	"bytes"
	"github.com/dchest/captcha"
	"github.com/kataras/iris/v12/mvc"
	"net/http"
	"strings"
	"time"
	"xupt-scholarship/db"
)

type CaptchaMVC struct {
	BaseController
}

type CaptchaStore struct {
}

func (CS *CaptchaStore) Set(id string, digits []byte) {
	redis := db.Redis
	redis.SET(id, string(digits))
	defer redis.Stop()
}

func (CS *CaptchaStore) Get(id string, clear bool) (digits []byte) {
	redis := db.Redis
	code, err := redis.GET(id)
	if clear {
		redis.DEL(id)
	}
	if err != nil {
		code = ""
	}
	defer redis.Stop()
	digits = []byte(code)
	return digits
}

func UseCaptchaMVC(app *mvc.Application) {
	app.Handle(new(CaptchaMVC))
}

func (C *CaptchaMVC) GetBy(imgPath string) {
	w, r := C.Ctx.ResponseWriter(), C.Ctx.Request()
	point := strings.Index(imgPath, ".")
	var content bytes.Buffer
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
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
func (C *CaptchaMVC) PostVerify() ResponseFmtData {
	var data VerifyData
	GetRequestParams(C.Ctx, &data)
	return ResponseFmtData{
		Message: "verifyed",
		Code:    0,
		Data:    captcha.VerifyString(data.Code, data.InputCode),
	}
}

func (C *CaptchaMVC) Get() ResponseFmtData {
	captcha.SetCustomStore(new(CaptchaStore))
	captchaId := captcha.New()
	return ResponseFmtData{
		Message: "success",
		Code:    0,
		Data:    captchaId,
	}
}

func (C *CaptchaMVC) GetReload() ResponseFmtData {
	code := C.Ctx.URLParam("code")
	return ResponseFmtData{
		Message: "success",
		Code:    0,
		Data:    captcha.Reload(code),
	}
}
