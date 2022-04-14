package service

import (
	"time"
	"xupt-scholarship/mvc_struct"
)

// NewProcessTaskSchedule 创建定时评定流程
func NewProcessTaskSchedule(task mvc_struct.ProcessTask, index int) {
	var temp *time.Timer
	if task.Duration > 0 {
		// 提前8小时通知
		temp = time.NewTimer(time.Duration(task.Duration-8) * time.Hour)
		go func() {
			<-temp.C
			if task.Type == "end" {
				SendEmail("end", task.Name, task.NotifyList, task.Date[1])
			} else {
				SendEmail("start", task.Name, task.NotifyList, task.Date[0], task.Date[1])
			}
		}()
	}
}

// NewProcessSchedule 创建定时评定流程
func NewProcessSchedule(task mvc_struct.ProcessTask, mentions []string) {
	var temp *time.Timer
	if task.Duration > 0 {
		// 提前8小时通知
		temp = time.NewTimer(time.Duration(task.Duration-48) * time.Hour)
		go func() {
			<-temp.C
			SendEmail("init", "init", mentions, task.Date[0])
		}()
	}
}

// HandleProcessTask 处理ProcessTask
func HandleProcessTask(processList []mvc_struct.ProcessTask, mentions []string) {
	NewProcessSchedule(processList[0], mentions)
	for i := 0; i < len(processList); i++ {
		NewProcessTaskSchedule(processList[i], i)
	}
}
