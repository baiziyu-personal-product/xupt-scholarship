package model

import (
	"encoding/json"
	"gorm.io/gorm"
	"xupt-scholarship/db"
	"xupt-scholarship/mvc_struct"
	"xupt-scholarship/utils"
)

var processModel ProcessModel

type ApplyModel struct {
}

type ApplyModelInterface interface {
	CheckIsExistThisYear(userId string) BaseModelFmtData
	CreateApplyForm(data mvc_struct.CreateApplyByBaseInfo) BaseModelFmtData
	UpdateApplyForm(data mvc_struct.UpdateApplyBaseInfo) BaseModelFmtData
	GetApplyData(applyId int, studentId string) BaseModelFmtData
	GetApplyList(filter mvc_struct.ApplyListFilterParams) BaseModelFmtData
}

// >>>>>>>>>>>>>> struct <<<<<<<<<<<<<<<<<<//

type ApplyFormBaseData struct {
	Id       int    `json:"id"`
	EditAt   string `json:"edit_at"`
	CreateAt string `json:"create_at"`
	Editable bool   `json:"editable"`
	Status   string `json:"status"`
}

type ApplyFormData struct {
	ApplyFormBaseData
	Form mvc_struct.ApplicationValue `json:"form"`
}

// >>>>>>>>>>>>>> interface <<<<<<<<<<<<<<<//

// CheckIsExistThisYear 检查是否存在当前学年的奖学金申请
func (a *ApplyModel) CheckIsExistThisYear(userId string) BaseModelFmtData {
	var application db.Application
	procedureId := processModel.GetProcessFormData(-1).Data.(ProcedureModelFormData).Id
	result := db.Mysql.Where("procedure_id = ? AND user_id = ?", procedureId, userId).Find(&application)
	return HandleDBData(result, application.ID)
}

// CreateApplyForm 创建奖学金申请
func (a *ApplyModel) CreateApplyForm(data mvc_struct.CreateApplyByBaseInfo) BaseModelFmtData {
	jsonForm, _ := json.Marshal(data.Form)
	Application := db.Application{
		Info:        jsonForm,
		History:     []byte("{}"),
		UserId:      data.StudentId,
		Status:      data.Type,
		Step:        "",
		ProcedureId: processModel.GetProcessFormData(-1).Data.(ProcedureModelFormData).Id,
	}
	result := db.Mysql.Create(&Application)
	return HandleDBData(result, Application.ID)
}

// UpdateApplyForm 更新奖学金信息
func (a *ApplyModel) UpdateApplyForm(data mvc_struct.UpdateApplyBaseInfo) BaseModelFmtData {
	jsonForm, _ := json.Marshal(data.Form)
	var apply db.Application
	var updateMap = map[string]interface{}{
		"status": data.Type,
		"info":   string(jsonForm),
		"score":  0,
	}
	result := db.Mysql.Model(&apply).Where("id = ?", data.Id).Updates(updateMap)
	return HandleDBData(result, apply.ID)
}

// GetApplyData 申请评定信息
func (a *ApplyModel) GetApplyData(applyId int, studentId string) BaseModelFmtData {
	var Application db.Application
	result := db.Mysql.First(&Application, applyId)
	var applicationData mvc_struct.ApplicationValue
	json.Unmarshal(Application.Info, &applicationData)
	return HandleDBData(result, ApplyFormData{
		ApplyFormBaseData: ApplyFormBaseData{
			CreateAt: utils.FmtTimeByUnix(Application.CreateAt),
			EditAt:   utils.FmtTimeByUnix(Application.UpdateAt),
			Editable: Application.UserId == studentId,
			Status:   Application.Status,
		},
		Form: applicationData,
	})
}

// GetApplyList 获取申请列表，用于评审
func (a *ApplyModel) GetApplyList(filter mvc_struct.ApplyListFilterParams) BaseModelFmtData {
	var applyList []ApplyFormBaseData
	var application db.Application
	var ApplicationList []db.Application
	processId := filter.ProcedureId
	if processId == -1 {
		processId = processModel.GetProcessFormData(-1).Data.(ProcedureModelFormData).Id
	}
	// 分页
	offset, limit := GetPageLimit(filter.PageCount, filter.PageIndex)
	var result *gorm.DB
	if filter.IsCheck == "manager" {
		result = db.Mysql.Limit(limit).Offset(offset).Where("procedure_id = ?", filter.ProcedureId).Find(&ApplicationList)
	} else {
		result = db.Mysql.Limit(limit).Offset(offset).Where("user_id = ? AND procedure_id = ?", filter.UserId, processId).Find(&ApplicationList)
	}
	for _, apply := range ApplicationList {
		var applicationData mvc_struct.ApplicationValue
		json.Unmarshal(application.Info, &applicationData)
		applyList = append(applyList, ApplyFormBaseData{
			Id:       apply.ID,
			CreateAt: utils.FmtTimeByUnix(apply.CreateAt),
			EditAt:   utils.FmtTimeByUnix(apply.UpdateAt),
			Editable: true,
			Status:   apply.Status,
		})
	}
	return HandleDBData(result, applyList)
}
