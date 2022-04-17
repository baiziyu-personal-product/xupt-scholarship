package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12/mvc"
	"xupt-scholarship/global"
	"xupt-scholarship/model"
	"xupt-scholarship/mvc_struct"
)

type ApplyMVC struct {
	BaseController
}

type ApplyController interface {
	Get() BaseControllerFmtData
	GetBy(applyId int) BaseControllerFmtData
	GetFormList() BaseControllerFmtData
	GetApplyHistory(id int) BaseControllerFmtData
	PostHandleFormBy(handleType string) BaseControllerFmtData
	PostScoreBy(id int) BaseControllerFmtData
	PutBy(applyId int, applyType string) BaseControllerFmtData
}

var applyModel model.ApplyModel

func UseApplyMVC(app *mvc.Application) {
	app.Register(userSession.Start).Handle(new(ApplyMVC))
}

//>>>>>>>>>>>>>>>>>>>>>> struct <<<<<<<<<<<<<<<<<<<<<<<<<<<//

type ApplyData struct {
	Data     interface{} `json:"data"`
	EditAble bool        `json:"edit_able"`
}

//>>>>>>>>>>>>>>>>>>>>> controllers <<<<<<<<<<<<<<<<<<<<//

// Get 获取用户是否创建本年度流程的申请表单
func (a *ApplyMVC) Get() BaseControllerFmtData {
	email := a.Session.GetString(sessionId)
	user := UserModel.GetUser(email).Data.(model.LoginUserInfo)
	applyInfo := applyModel.CheckIsExistThisYear(user.UserId)
	return HandleControllerRes(applyInfo, "获取用户申请状态")
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

func (a *ApplyMVC) GetApplyHistory(id int) BaseControllerFmtData {
	m := applyModel.GetApplyHistory(id)
	return HandleControllerRes(m, "获取评定记录")
}

// PostHandleFormBy 创建奖学金申请
func (a *ApplyMVC) PostHandleFormBy(handleType string) BaseControllerFmtData {
	var reqData mvc_struct.ApplicationRequest
	GetRequestParams(a.Ctx, &reqData)
	user := GetUserData(a.Session)
	applyInfo := applyModel.CheckIsExistThisYear(user.UserId)
	// 当前年度内无法重复创建
	if applyInfo.Error == nil {
		return BaseControllerFmtData{
			Message: "已存在当前年度申请奖学金表单，无法重复申请",
			Code:    global.ErrorCode,
			Data:    applyInfo.Data,
		}
	}
	fmt.Println(reqData)
	formModel := applyModel.CreateApplyForm(mvc_struct.CreateApplyByBaseInfo{
		Form: mvc_struct.ApplicationValue{
			Moral:    reqData.Moral,
			Practice: reqData.Practice,
			Academic: reqData.Academic,
		},
		StudentId: user.UserId,
		Type:      handleType,
		ScoreInfo: reqData.ScoreInfo,
	})
	return HandleControllerRes(formModel, "申请信息创建")
}

func (a *ApplyMVC) PostScoreBy(id int) BaseControllerFmtData {
	var reqData mvc_struct.ApplyScoreInfo
	GetRequestParams(a.Ctx, &reqData)
	user := GetUserData(a.Session)
	if user.Identity == "student" {
		return BaseControllerFmtData{
			Message: "没有权限参与当前评定",
			Code:    global.ErrorCode,
			Data:    nil,
		}
	}
	comment, _ := json.Marshal(reqData)
	res := applyModel.UpdateApplyScore(id, user.UserId, user.Identity, reqData, string(comment))
	return HandleControllerRes(res, "评定申请流程")
}

func (a *ApplyMVC) PutBy(applyId int, applyType string) BaseControllerFmtData {
	var reqData mvc_struct.ApplicationRequest
	GetRequestParams(a.Ctx, &reqData)
	user := GetUserData(a.Session)
	formModel := applyModel.UpdateApplyForm(user.UserId, mvc_struct.UpdateApplyBaseInfo{
		Form: mvc_struct.ApplicationValue{
			Moral:    reqData.Moral,
			Practice: reqData.Practice,
			Academic: reqData.Academic,
		},
		Id:        applyId,
		StudentId: user.UserId,
		Type:      applyType,
		ScoreInfo: reqData.ScoreInfo,
	})
	return HandleControllerRes(formModel, "申请信息处理")
}
