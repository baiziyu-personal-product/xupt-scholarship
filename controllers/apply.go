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

func (a *ApplyMVC) GetBy(applyId int) BaseControllerFmtData {
	return BaseControllerFmtData{
		Message: "",
		Code:    1,
		Data:    applyId,
	}
}

func (a *ApplyMVC) Post() BaseControllerFmtData {
	return BaseControllerFmtData{
		Message: "",
		Code:    1,
		Data:    nil,
	}
}

func (a *ApplyMVC) PostHandleFormBy(handleType string) BaseControllerFmtData {
	var reqData mvc_struct.ApplicationValue
	GetRequestParams(a.Ctx, &reqData)
	email := a.Session.GetString(sessionId)
	user := UserModel.GetUser(email).Data.(model.LoginUserInfo)
	formModel := applyModel.CreateApplyForm(mvc_struct.BaseApply{
		Form:      reqData,
		StudentId: user.StudentId,
		Type:      handleType,
	})
	return HandleControllerRes(formModel, "申请信息保存")
}

func (a *ApplyMVC) PostFormSubmit() {

}

func (a *ApplyMVC) PutBy(applyId int64) BaseControllerFmtData {
	return BaseControllerFmtData{
		Message: "",
		Code:    0,
		Data:    applyId,
	}
}
