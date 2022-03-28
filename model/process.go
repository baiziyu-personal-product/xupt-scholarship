package model

import (
	"encoding/json"
	"gorm.io/gorm"
	"xupt-scholarship/db"
	"xupt-scholarship/mvc_struct"
	"xupt-scholarship/utils"
)

type ProcessModel struct {
}

func (p *ProcessModel) CreateProcess(data mvc_struct.ProcessFormData, userId string) BaseModelFmtData {
	info, _ := json.Marshal(data)
	process := db.Procedure{
		CurrentStep: []byte("{}"),
		UserId:      userId,
		Info:        info,
		History:     []byte("{}"),
	}
	result := db.Mysql.Create(&process)
	return HandleDBData(result, process.ID)
}

type ProcedureModelFormData struct {
	Id       int                        `json:"id"`
	Form     mvc_struct.ProcessFormData `json:"form"`
	UserId   string                     `json:"user_id"`
	CreateAt string                     `json:"create_at"`
	EditAt   string                     `json:"edit_at"`
}

func (p *ProcessModel) GetProcessFormData(id int) BaseModelFmtData {
	var processData mvc_struct.ProcessFormData
	var processInfo db.Procedure
	var result *gorm.DB
	if id == -1 {
		result = db.Mysql.Last(&processInfo)
	} else {
		result = db.Mysql.First(&processInfo, id)
	}
	json.Unmarshal(processInfo.Info, &processData)
	return HandleDBData(result, ProcedureModelFormData{
		Id:       id,
		Form:     processData,
		UserId:   processInfo.UserId,
		CreateAt: utils.FmtTimeByUnix(processInfo.CreateAt),
		EditAt:   utils.FmtTimeByUnix(processInfo.UpdateAt),
	})
}

func (p *ProcessModel) UpdateProcessFormData(id int, info mvc_struct.ProcessFormData) BaseModelFmtData {
	jsonInfo, _ := json.Marshal(info)
	var procedure db.Procedure
	result := db.Mysql.Model(procedure).Where("id = ?", id).Update("info", jsonInfo)
	return HandleDBData(result, id)
}

type stepData struct {
	History []mvc_struct.ProcessHistoryItem `json:"history"`
	Current mvc_struct.ProcessHistoryItem   `json:"current"`
}

func (p *ProcessModel) GetProcessStep(id int) BaseModelFmtData {
	var procedure db.Procedure
	var stepHistory []mvc_struct.ProcessHistoryItem
	var currentStep mvc_struct.ProcessHistoryItem
	result := db.Mysql.First(&procedure, id)
	json.Unmarshal(procedure.CurrentStep, &currentStep)
	json.Unmarshal(procedure.History, &stepHistory)
	return HandleDBData(result, stepData{
		History: stepHistory,
		Current: currentStep,
	})
}
