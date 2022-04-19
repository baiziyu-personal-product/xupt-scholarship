package service

import (
	"encoding/json"
	"xupt-scholarship/db"
	"xupt-scholarship/global"
	"xupt-scholarship/mvc_struct"
	utils "xupt-scholarship/utils"
)

// AddProcessStep 添加ProcessStep
func AddProcessStep(id int, task mvc_struct.ProcessStepSchedule) {
	var stepHistory []mvc_struct.ProcessHistoryItem
	var currentStep mvc_struct.ProcessHistoryItem
	var procedure db.Procedure
	result := db.Mysql.First(&procedure, id)
	if result.Error == nil {
		json.Unmarshal(procedure.History, &stepHistory)
		currentStep = mvc_struct.ProcessHistoryItem{
			StartAt: utils.GetCurrentTime(),
			Step:    task.Step,
		}
		stepHistory = append(stepHistory, currentStep)
		history, _ := json.Marshal(stepHistory)
		step, _ := json.Marshal(currentStep)

		data := map[string]interface{}{"history": history, "current_step": step}

		result := db.Mysql.Model(&procedure).Where("id = ?", id).Updates(data)
		if result.Error != nil {
			panic(result.Error)
		}
	} else {
		panic(result.Error)
	}
}

// UpdateProcessStep 更新ProcessStep
func UpdateProcessStep(task mvc_struct.ProcessStepSchedule) {
	var stepHistory []mvc_struct.ProcessHistoryItem
	var currentStep mvc_struct.ProcessHistoryItem
	var procedure db.Procedure
	result := db.Mysql.Last(&procedure)
	if result.Error == nil {
		json.Unmarshal(procedure.History, &stepHistory)
		currentStep = mvc_struct.ProcessHistoryItem{
			StartAt: utils.GetCurrentTime(),
			Step:    task.Step,
		}
		stepHistory = append(stepHistory, currentStep)
		history, _ := json.Marshal(stepHistory)
		step, _ := json.Marshal(currentStep)

		data := map[string]interface{}{"history": history, "current_step": step}

		result := db.Mysql.Model(&procedure).Where("id = ?", procedure.ID).Updates(data)
		if result.Error != nil {
			panic(result.Error)
		}
	} else {
		panic(result.Error)
	}
}

// HandleProcessTask 处理ProcessTask
func HandleProcessTask(processList []mvc_struct.ProcessStepSchedule, mentions []string, processId int) {
	var initProcess = processList[0]
	initProcess.Status = global.ProcessInit
	NewTimerSchedule(initProcess, UpdateProcessStep)
	for i := 0; i < len(processList); i++ {
		NewTimerSchedule(processList[i], UpdateProcessStep)
	}
}
