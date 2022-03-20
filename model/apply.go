package model

import (
	"encoding/json"
	"xupt-scholarship/db"
	"xupt-scholarship/mvc_struct"
	"xupt-scholarship/utils"
)

type ApplyModel struct {
}

func (a *ApplyModel) CreateApplyForm(data mvc_struct.CreateApplyByBaseInfo) BaseModelFmtData {
	jsonForm, _ := json.Marshal(data.Form)
	Application := db.Application{
		Form:    jsonForm,
		History: []byte("{}"),
		Creator: data.StudentId,
		Status:  data.Type,
		Step:    "",
	}
	result := db.Mysql.Create(&Application)
	return HandleDBData(result, Application.ID)
}

func (a *ApplyModel) UpdateApplyForm(data mvc_struct.UpdateApplyBaseInfo) BaseModelFmtData {
	jsonForm, _ := json.Marshal(data.Form)
	apply := db.Application{
		BaseDataModel: db.BaseDataModel{ID: data.Id},
	}
	var updateMap = map[string]interface{}{
		"status": data.Type,
		"form":   string(jsonForm),
		"score":  0,
	}
	result := db.Mysql.Model(&apply).Updates(updateMap)
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
	json.Unmarshal(Application.Form, &applicationData)
	return HandleDBData(result, ApplyFormData{
		ApplyFormBaseData: ApplyFormBaseData{
			CreateAt: utils.FmtTimeByUnix(Application.CreateAt),
			EditAt:   utils.FmtTimeByUnix(Application.UpdateAt),
			Editable: Application.Creator == studentId,
			Status:   Application.Status,
		},
		Form: applicationData,
	})
}
func (a *ApplyModel) GetApplyList(studentId string) BaseModelFmtData {
	var applyList []ApplyFormBaseData
	Application := db.Application{
		Creator: studentId,
	}
	var ApplicationList []db.Application
	result := db.Mysql.Where(&Application).Find(&ApplicationList)
	for _, apply := range ApplicationList {
		var applicationData mvc_struct.ApplicationValue
		json.Unmarshal(Application.Form, &applicationData)
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
