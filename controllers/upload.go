package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"path/filepath"
	"xupt-scholarship/global"
)

const maxSize = 240 * iris.MB

type UploadMvc struct {
	BaseController
}

func UseUploadMVC(app *mvc.Application) {
	app.Handle(new(UploadMvc))
}

func (u *UploadMvc) PostSingle() ResponseFmtData {
	u.Ctx.SetMaxRequestBodySize(maxSize)
	_, fileHeader, err := u.Ctx.FormFile("file")
	if err != nil {
		u.Ctx.StopWithError(iris.StatusBadRequest, err)
		return ResponseFmtData{
			Message: "failed",
			Code:    0,
			Data:    nil,
		}
	}
	dest, _ := filepath.Abs(filepath.Join(global.Settings.ImagePath, fileHeader.Filename))
	u.Ctx.SaveFormFile(fileHeader, dest)
	return ResponseFmtData{
		Message: "success",
		Code:    0,
		Data:    dest,
	}
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
		filePath, _ := filepath.Abs(filepath.Join(global.Settings.ImagePath, file.Filename))
		res = append(res, filePath)
	}
	return ResponseFmtData{
		Message: "Success",
		Code:    1,
		Data:    res,
	}
}
