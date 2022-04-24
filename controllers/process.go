package controllers

import (
	"github.com/kataras/iris/v12/mvc"
	"xupt-scholarship/global"
	"xupt-scholarship/model"
	"xupt-scholarship/mvc_struct"
)

type ProcessMVC struct {
	BaseController
}

type ProcessController interface {
	Get() BaseControllerFmtData
	GetBy(processId int) BaseControllerFmtData
	GetStepBy(processId int) BaseControllerFmtData
	Post() BaseControllerFmtData
	PutBy(processId int) BaseControllerFmtData
}

var ProcessModel model.ProcessModel

func UseProcessMVC(app *mvc.Application) {
	app.Register(userSession.Start).Handle(new(ProcessMVC))
}

func (p *ProcessMVC) Get() BaseControllerFmtData {
	user := GetUserData(p.Session)
	m := ProcessModel.GetCurrentYearProcess(user.UserId, user.Identity)
	return HandleControllerRes(m, "获取评定信息")
}

func (p *ProcessMVC) GetBy(processId int) BaseControllerFmtData {
	m := ProcessModel.GetProcessFormData(processId)
	return HandleControllerRes(m, "获取评定流程")
}

func (p *ProcessMVC) GetStepBy(processId int) BaseControllerFmtData {
	m := ProcessModel.GetProcessStep(processId)
	return HandleControllerRes(m, "评定流程状态")
}

func (p *ProcessMVC) Post() BaseControllerFmtData {
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

	if m.Error == nil {
		model.DispatchProcessNoticeEvent(processInfo, m.Data.(int))
	}
	return HandleControllerRes(m, "创建流程")
}

func (p *ProcessMVC) PutBy(processId int) BaseControllerFmtData {
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
