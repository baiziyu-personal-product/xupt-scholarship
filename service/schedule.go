package service

import (
	"time"
	"xupt-scholarship/global"
	"xupt-scholarship/mvc_struct"
)

func handleProcessServiceDurationByStatus(status string, duration float64) *time.Timer {
	if status == global.ProcessInit {
		return time.NewTimer(time.Duration(duration-global.ProcessInitDurationHours) * time.Hour)
	}
	return time.NewTimer(time.Duration(duration-global.ProcessStartDurationHours) * time.Hour)
}

// NewTimerSchedule 奖学金评定流程发起
func NewTimerSchedule(
	task mvc_struct.ProcessStepSchedule,
	updateProcessStep func(mvc_struct.ProcessStepSchedule),
) {
	var temp *time.Timer
	if task.Duration > 0 {
		// 提前8小时通知
		temp = handleProcessServiceDurationByStatus(task.Status, task.Duration)
		info := ProcessBaseEmailOptions{
			Step:      task.Name,
			EndDate:   task.Date[1],
			StartDate: task.Date[0],
			Status:    task.Status,
			Receiver:  task.NotifyList,
		}
		go func() {
			<-temp.C
			SendEmail(info)
			updateProcessStep(task)
		}()
	}
}
