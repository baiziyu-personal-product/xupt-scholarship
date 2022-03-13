package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"path/filepath"
	"xupt-scholarship/global"
)

const maxSize = 240 * iris.MB

type UploadMVC struct {
	BaseController
}

func UseUploadMVC(app *mvc.Application) {
	app.Handle(new(UploadMVC))
}

func (u *UploadMVC) PostSingle() BaseControllerFmtData {
	u.Ctx.SetMaxRequestBodySize(maxSize)
	_, fileHeader, err := u.Ctx.FormFile("file")
	if err != nil {
		u.Ctx.StopWithError(iris.StatusBadRequest, err)
		return BaseControllerFmtData{
			Message: "failed",
			Code:    0,
			Data:    nil,
		}
	}
	dest, _ := filepath.Abs(filepath.Join(global.Settings.ImagePath, fileHeader.Filename))
	u.Ctx.SaveFormFile(fileHeader, dest)
	return BaseControllerFmtData{
		Message: "success",
		Code:    0,
		Data:    dest,
	}
}

func (u *UploadMVC) Options() BaseControllerFmtData {
	return BaseControllerFmtData{
		Message: "success",
		Code:    0,
		Data:    nil,
	}
}

func (u *UploadMVC) Post() BaseControllerFmtData {
	files, _, err := u.Ctx.UploadFormFiles(global.Settings.ImagePath)
	if err != nil {
		u.Ctx.StopWithStatus(iris.StatusInternalServerError)
		return BaseControllerFmtData{
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
	return BaseControllerFmtData{
		Message: "Success",
		Code:    1,
		Data:    res,
	}
}
