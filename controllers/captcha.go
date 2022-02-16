package controllers

import (
	"bytes"
	"github.com/dchest/captcha"
	"github.com/kataras/iris/v12/mvc"
	"net/http"
	"strings"
	"text/template"
	"time"
)

type Captcha struct {
	BaseController
	code string
}

func CaptchaMVC(app *mvc.Application) {
	app.Handle(new(Captcha))
}

func (C *Captcha) GetBy(code string) {
	w, r := C.Ctx.ResponseWriter(), C.Ctx.Request()
	C.code = code
	point := strings.Index(code, ".")
	var content bytes.Buffer
	w.Header().Set("Content-Type", "image/png")
	captcha.WriteImage(&content, code[:point], captcha.StdWidth, captcha.StdWidth)
	http.ServeContent(w, r, code, time.Time{}, bytes.NewReader(content.Bytes()))
}

func (C *Captcha) Get() http.Handler {
	w := C.Ctx.ResponseWriter()
	//if r.URL.Path != "/" {
	//	http.NotFound(w, r)
	//	return nil
	//}
	d := struct {
		CaptchaId string
	}{
		captcha.New(),
	}
	if err := formTemplate.Execute(w, &d); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return captcha.Server(captcha.StdWidth, captcha.StdHeight)
}

var formTemplate = template.Must(template.New("example").Parse(formTemplateSrc))

const formTemplateSrc = `<!doctype html>
<head><title>Captcha Example</title></head>
<body>
<script>
function setSrcQuery(e, q) {
    var src  = e.src;
    var p = src.indexOf('?');
    if (p >= 0) {
        src = src.substr(0, p);
    }
    e.src = src + "?" + q
}
function reload() {
    setSrcQuery(document.getElementById('image'), "reload=" + (new Date()).getTime());
    setSrcQuery(document.getElementById('audio'), (new Date()).getTime());
    return false;
}
</script>
<form action="/process" method=post>
<p>Type the numbers you see in the picture below:</p>
<p><img id=image src="/captcha/{{.CaptchaId}}.png" alt="Captcha image"></p>
<a href="#" onclick="reload()">Reload</a> 
<input type=hidden name=captchaId value="{{.CaptchaId}}"><br>
<input name=captchaSolution>
<input type=submit value=Submit>
</form>
`
