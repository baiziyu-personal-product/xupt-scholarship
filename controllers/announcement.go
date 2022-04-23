package controllers

import (
	"github.com/kataras/iris/v12/mvc"
	"xupt-scholarship/model"
)

type AnnouncementMVC struct {
	BaseController
}

type AnnouncementController interface {
	GetBy(processId int) BaseControllerFmtData
}

var announcementModel model.AnnouncementModel

func UseAnnouncementMVC(app *mvc.Application) {
	app.Register(userSession.Start).Handle(new(AnnouncementMVC))
}

func (a *AnnouncementMVC) GetBy(procedureId int) BaseControllerFmtData {
	result := announcementModel.GetAnnouncementDataByProcedureId(procedureId)
	return HandleControllerRes(result, "获取公示信息")
}
