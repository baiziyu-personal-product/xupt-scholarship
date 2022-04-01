package controllers

import (
	"github.com/kataras/iris/v12/mvc"
	"xupt-scholarship/global"
	"xupt-scholarship/model"
	"xupt-scholarship/mvc_struct"
)

type processMVC struct {
	BaseController
}

var ProcessModel model.ProcessModel

func UseProcessMVC(app *mvc.Application) {
	app.Register(userSession.Start).Handle(new(processMVC))
}

func (p *processMVC) Get() BaseControllerFmtData {
	user := GetUserData(p.Session)
	m := ProcessModel.GetCurrentYearProcess(user.UserId, user.Identity)
	return HandleControllerRes(m, "获取评定信息")
}

func (p *processMVC) GetBy(processId int) BaseControllerFmtData {
	m := ProcessModel.GetProcessFormData(processId)
	return HandleControllerRes(m, "获取评定流程")
}

func (p *processMVC) Post() BaseControllerFmtData {
	var processInfo mvc_struct.ProcessFormData
	GetRequestParams(p.Ctx, &processInfo)
	user := GetUserData(p.Session)
	exist := ProcessModel.GetCurrentYearProcess(user.UserId, user.Identity)
	if exist.Code == global.SuccessCode {
		processId := exist.Data.(model.ProcessStatusRes).ProcessId
		if processId > 0 {
			return HandleControllerRes(model.BaseModelFmtData{
				Message: "Has exist",
				Data:    processId,
				Error:   nil,
				Code:    global.ErrorCode,
			}, "已存在当前学期的流程，无法二次创建")
		}
	}
	m := ProcessModel.CreateProcess(processInfo, user.UserId)
	return HandleControllerRes(m, "创建流程")
}

func (p *processMVC) PutBy(processId int) BaseControllerFmtData {
	var processInfo mvc_struct.ProcessFormData
	GetRequestParams(p.Ctx, &processInfo)
	user := GetUserData(p.Session)
	exist := ProcessModel.GetCurrentYearProcess(user.UserId, user.Identity)
	editable := exist.Data.(model.ProcessStatusRes).Editable
	if editable == false {
		return HandleControllerRes(model.BaseModelFmtData{
			Message: "不可修改",
			Data:    processId,
			Error:   nil,
			Code:    global.ErrorCode,
		}, "当前阶段无法进行修改")
	}
	m := ProcessModel.UpdateProcessFormData(processId, processInfo)
	return HandleControllerRes(m, "修改流程")
}
