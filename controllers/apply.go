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

// Get 获取用户是否创建本年度流程的申请表单
func (a *ApplyMVC) Get() BaseControllerFmtData {
	email := a.Session.GetString(sessionId)
	user := UserModel.GetUser(email).Data.(model.LoginUserInfo)
	applyId := applyModel.CheckIsExistThisYear(user.UserId)
	return HandleControllerRes(applyId, "获取用户申请状态")
}

// GetBy 获取对应ID的申请表单
func (a *ApplyMVC) GetBy(applyId int) BaseControllerFmtData {
	email := a.Session.GetString(sessionId)
	user := UserModel.GetUser(email).Data.(model.LoginUserInfo)
	applyData := applyModel.GetApplyData(applyId, user.UserId)
	return HandleControllerRes(applyData, "获取申请表单")
}

// GetFormList 获取对应筛选条件下的申请表单，支持分页，默认返回最近学年
func (a *ApplyMVC) GetFormList() BaseControllerFmtData {
	pageCount := a.Ctx.URLParamIntDefault("page_count", 10)
	pageIndex := a.Ctx.URLParamIntDefault("page_index", 1)
	isCheck := a.Ctx.URLParamDefault("is_check", "manager")
	procedureId := a.Ctx.URLParamIntDefault("procedure_id", -1)
	email := a.Session.GetString(sessionId)
	user := UserModel.GetUser(email).Data.(model.LoginUserInfo)
	applyData := applyModel.GetApplyList(mvc_struct.ApplyListFilterParams{
		UserId:      user.UserId,
		PageCount:   pageCount,
		PageIndex:   pageIndex,
		IsCheck:     isCheck,
		ProcedureId: procedureId,
	})
	return HandleControllerRes(applyData, "获取表单列表")
}

func (a *ApplyMVC) Post() BaseControllerFmtData {
	return BaseControllerFmtData{
		Message: "",
		Code:    1,
		Data:    nil,
	}
}

// PostHandleFormBy 创建奖学金申请
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
