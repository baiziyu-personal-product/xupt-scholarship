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

type CaptchaController interface {
	GetBy(imgPath string)
	PostVerify() BaseControllerFmtData
	Get() BaseControllerFmtData
	GetReload() BaseControllerFmtData
}

//>>>>>>>>>>>>>>>>>>struct <<<<<<<<<<<<<<<//

type CaptchaStore struct {
}

func (CS *CaptchaStore) Set(id string, digits []byte) {
	redis := db.Redis
	redis.SETEX(id, int(time.Minute*5), string(digits))
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
	digits = []byte(code)
	return digits
}

func UseCaptchaMVC(app *mvc.Application) {
	app.Handle(new(CaptchaMVC))
}

//>>>>>>>>>>>>>>>>> controllers <<<<<<<<<<<<<<<<//

func (C *CaptchaMVC) GetBy(imgPath string) {
	w, r := C.Ctx.ResponseWriter(), C.Ctx.Request()
	var content bytes.Buffer
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Status", "image/png")
	code := imgPath[:strings.Index(imgPath, "-")]
	captcha.WriteImage(&content, code, 180, 60)
	http.ServeContent(w, r, imgPath, time.Time{}, bytes.NewReader(content.Bytes()))
}

type VerifyData struct {
	Code      string `json:"code"`
	InputCode string `json:"input_code"`
}

// PostVerify 判断验证码是否正确
func (C *CaptchaMVC) PostVerify() BaseControllerFmtData {
	var data VerifyData
	GetRequestParams(C.Ctx, &data)
	return BaseControllerFmtData{
		Message: "verified",
		Code:    0,
		Data:    captcha.VerifyString(data.Code, data.InputCode),
	}
}

func (C *CaptchaMVC) Get() BaseControllerFmtData {
	captcha.SetCustomStore(new(CaptchaStore))
	captchaId := captcha.New()
	return BaseControllerFmtData{
		Message: "success",
		Code:    0,
		Data:    captchaId,
	}
}

func (C *CaptchaMVC) GetReload() BaseControllerFmtData {
	code := C.Ctx.URLParam("code")
	return BaseControllerFmtData{
		Message: "success",
		Code:    0,
		Data:    captcha.Reload(code),
	}
}
