package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"path/filepath"
	"xupt-scholarship/global"
)

const maxSize = 16 * iris.MB

type UploadMvc struct {
	BaseController
}

func UseUploadMVC(app *mvc.Application) {
	app.Handle(new(UploadMvc))
}

func (u *UploadMvc) Post() ResponseFmtData {
	files, _, err := u.Ctx.UploadFormFiles(global.Settings.ImagePath)
	if err != nil {
		u.Ctx.StopWithStatus(iris.StatusInternalServerError)
		return ResponseFmtData{
			Message: "failed",
			Code:    0,
			Data:    nil,
		}
	}
	var res []string
	for _, file := range files {
		p, _ := filepath.Abs(global.Settings.ImagePath + file.Filename)
		fmt.Println(p)
		res = append(res, global.Settings.ImagePath+file.Filename)
	}
	return ResponseFmtData{
		Message: "Success",
		Code:    1,
		Data:    res,
	}
}
