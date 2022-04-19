package service

import (
	"encoding/json"
	"time"
	"xupt-scholarship/db"
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

// NewProcessTaskSchedule 创建定时评定流程
func NewProcessTaskSchedule(task mvc_struct.ProcessStepSchedule, index int, processId int) {
	var temp *time.Timer
	if task.Duration > 0 {
		// 提前8小时通知
		//temp = time.NewTimer(time.Duration(task.Duration-8) * time.Hour)
		go func() {
			<-temp.C
			//if task.Status == "end" {
			//	SendEmail("end", task.Name, task.NotifyList, task.Date[1])
			//} else {
			//	SendEmail("start", task.Name, task.NotifyList, task.Date[0], task.Date[1])
			//}
			AddProcessStep(processId, task)
		}()
	}
}

// NewProcessSchedule 创建定时评定流程
func NewProcessSchedule(task mvc_struct.ProcessStepSchedule, mentions []string, processId int) {
	var temp *time.Timer
	//if task.Duration > 0 {
	// 提前8小时通知
	//temp = time.NewTimer(time.Duration(task.Duration-48) * time.Hour)
	go func() {
		<-temp.C
		//SendEmail("init", "init", mentions, task.Date[0])
		AddProcessStep(processId, task)
	}()
	//}
}

// HandleProcessTask 处理ProcessTask
func HandleProcessTask(processList []mvc_struct.ProcessStepSchedule, mentions []string, processId int) {
	NewProcessSchedule(processList[0], mentions, processId)
	for i := 0; i < len(processList); i++ {
		NewProcessTaskSchedule(processList[i], i, processId)
	}
}
