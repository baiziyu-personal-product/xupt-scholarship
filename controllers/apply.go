package controllers

import (
	"github.com/kataras/iris/v12/mvc"
	"xupt-scholarship/model"
	"xupt-scholarship/mvc_struct"
)

type ApplyMVC struct {
	BaseController
}

var applyModel model.ApplyModel

func UseApplyMVC(app *mvc.Application) {
	app.Register(userSession.Start).Handle(new(ApplyMVC))
}

type ApplyData struct {
	Data     interface{} `json:"data"`
	EditAble bool        `json:"edit_able"`
}

func (a *ApplyMVC) GetBy(applyId int) BaseControllerFmtData {
	email := a.Session.GetString(sessionId)
	user := UserModel.GetUser(email).Data.(model.LoginUserInfo)
	applyData := applyModel.GetApplyData(applyId, user.UserId)
	return HandleControllerRes(applyData, "获取申请表单")
}

func (a *ApplyMVC) GetFormList() BaseControllerFmtData {
	email := a.Session.GetString(sessionId)
	user := UserModel.GetUser(email).Data.(model.LoginUserInfo)
	applyData := applyModel.GetApplyList(user.UserId)
	return HandleControllerRes(applyData, "获取表单列表")
}

func (a *ApplyMVC) Post() BaseControllerFmtData {
	return BaseControllerFmtData{
		Message: "",
		Code:    1,
		Data:    nil,
	}
}

type handleFormReqParams struct {
	Id    int                         `json:"id" default:"-1"`
	Value mvc_struct.ApplicationValue `json:"value"`
}

func (a *ApplyMVC) PostHandleFormBy(handleType string) BaseControllerFmtData {
	var reqData mvc_struct.ApplicationValue
	GetRequestParams(a.Ctx, &reqData)
	email := a.Session.GetString(sessionId)
	user := UserModel.GetUser(email).Data.(model.LoginUserInfo)
	formModel := applyModel.CreateApplyForm(mvc_struct.CreateApplyByBaseInfo{
		Form:      reqData,
		StudentId: user.UserId,
		Type:      handleType,
	})
	return HandleControllerRes(formModel, "申请信息创建")
}

func (a *ApplyMVC) PutBy(applyId int, applyType string) BaseControllerFmtData {
	var reqData mvc_struct.ApplicationValue
	GetRequestParams(a.Ctx, &reqData)
	email := a.Session.GetString(sessionId)
	user := UserModel.GetUser(email).Data.(model.LoginUserInfo)
	formModel := applyModel.UpdateApplyForm(mvc_struct.UpdateApplyBaseInfo{
		Form:      reqData,
		Id:        applyId,
		StudentId: user.UserId,
		Type:      applyType,
	})
	return HandleControllerRes(formModel, "申请信息处理")
}
