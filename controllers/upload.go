package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"path/filepath"
	"xupt-scholarship/global"
	"xupt-scholarship/mvc_struct"
)

const maxSize = 240 * iris.MB

type UploadMVC struct {
	BaseController
}

type UploadController interface {
	OptionsSingle() BaseControllerFmtData
	PostSingleBy(upType string) BaseControllerFmtData
	Options() BaseControllerFmtData
	Post() BaseControllerFmtData
	PostStudentList() BaseControllerFmtData
}

func UseUploadMVC(app *mvc.Application) {
	app.Handle(new(UploadMVC))
}

func (u *UploadMVC) OptionsSingle() BaseControllerFmtData {
	return BaseControllerFmtData{
		Message: "success",
		Code:    global.SuccessCode,
		Data:    nil,
	}
}

func (u *UploadMVC) PostSingleBy(upType string) BaseControllerFmtData {
	u.Ctx.SetMaxRequestBodySize(maxSize)
	_, fileHeader, err := u.Ctx.FormFile("file")
	if err != nil {
		u.Ctx.StopWithError(iris.StatusBadRequest, err)
		return BaseControllerFmtData{
			Message: "failed",
			Code:    global.ErrorCode,
			Data:    nil,
		}
	}
	var filePath = global.Settings.FilePath
	if upType == "avatar" {
		filePath = global.Settings.AvatarPath
	} else if upType == "image" {
		filePath = global.Settings.ImagePath
	}
	dest, _ := filepath.Abs(filepath.Join(filePath, fileHeader.Filename))
	u.Ctx.SaveFormFile(fileHeader, dest)
	return BaseControllerFmtData{
		Message: "success",
		Code:    global.SuccessCode,
		Data:    filePath + fileHeader.Filename,
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
			Code:    global.ErrorCode,
			Data:    nil,
		}
	}
	var res []string
	for _, file := range files {
		filepath.Abs(filepath.Join(global.Settings.FilePath, file.Filename))
		res = append(res, global.Settings.FilePath+file.Filename)
	}
	return BaseControllerFmtData{
		Message: "Success",
		Code:    global.SuccessCode,
		Data:    res,
	}
}

func (u *UploadMVC) PostStudentList() BaseControllerFmtData {
	var userList []mvc_struct.StudentItem
	GetRequestParams(u.Ctx, &userList)

	m := UserModel.CreateStudentByList(userList)
	return HandleControllerRes(m, "上传学生名单")
}
