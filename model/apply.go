package model

import (
	"encoding/json"
	"gorm.io/gorm"
	"xupt-scholarship/db"
	"xupt-scholarship/mvc_struct"
	"xupt-scholarship/utils"
)

type ApplyModel struct {
}

func (a *ApplyModel) CreateApplyForm(data mvc_struct.CreateApplyByBaseInfo) BaseModelFmtData {
	jsonForm, _ := json.Marshal(data.Form)
	Application := db.Application{
		Info:    jsonForm,
		History: []byte("{}"),
		UserId:  data.StudentId,
		Status:  data.Type,
		Step:    "",
	}
	result := db.Mysql.Create(&Application)
	return HandleDBData(result, Application.ID)
}

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

func (a *ApplyModel) GetApplyList(
	userId string,
	pageCount int,
	pageIndex int,
	isCheck string,
	lastDate string,
) BaseModelFmtData {
	var applyList []ApplyFormBaseData
	var application db.Application
	var ApplicationList []db.Application
	var startDate string
	if lastDate == "" {
		startDate = lastDate
	} else {
		// 获取最近一次发布奖学金申请流程
		// 并且获取其中运行学生申请奖学金的时间
		startDate = GetProcessFormData(-1).Data.(ProcedureModelFormData).Form.Form.IndividualApplicationStage.Date[0]
	}
	yearTime := GetCurrentYear(startDate)
	// 分页
	offset, limit := GetPageLimit(pageCount, pageIndex)
	var result *gorm.DB
	if isCheck == "manager" {
		result = db.Mysql.Limit(limit).Offset(offset).Where("create_at > ?", yearTime).Find(&ApplicationList)
	} else {
		result = db.Mysql.Limit(limit).Offset(offset).Where("user_id = ? AND create_at > ?", userId, yearTime).Find(&ApplicationList)
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
