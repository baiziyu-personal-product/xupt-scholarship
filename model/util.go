package model

import (
	"time"
	"xupt-scholarship/global"
	"xupt-scholarship/mvc_struct"
	"xupt-scholarship/service"
)

const timeLayout = "2006-01-02"

func GetPageLimit(pageCount int, pageIndex int) (int, int) {
	return pageCount * (pageIndex - 1), pageCount * pageIndex
}

func GetCurrentYear(startDate string) int64 {
	date := startDate
	currentYear, _, _ := time.Now().Date()
	currentLocation := time.Now().Location()
	if startDate == "" {
		firstOfYear := time.Date(currentYear, time.January, 1, 0, 0, 0, 0, currentLocation)
		return firstOfYear.Unix()
	}
	temp, _ := time.ParseInLocation(timeLayout, date, currentLocation)
	return temp.Unix()
}

func GetIsLate(startDate string) bool {
	return GetDateDurationByHour(startDate) < 24
}

func GetDateDurationByHour(startDate string) float64 {
	currentLocation := time.Now().Location()

	start, _ := time.ParseInLocation(timeLayout, startDate, currentLocation)
	return start.Sub(time.Now()).Hours()
}

// DispatchProcessNoticeEvent 创建流程通知事件
func DispatchProcessNoticeEvent(processInfo mvc_struct.ProcessFormData, processId int) {
	var processTasks []mvc_struct.ProcessStepSchedule
	var mentions []string
	var processList = global.ProcessStepList
	for i, v := range processInfo.Form {
		mentions = append(mentions, v.Mentions...)
		processTasks = append(
			processTasks,
			mvc_struct.ProcessStepSchedule{
				Name:       processList[i],
				Step:       v.Step,
				Duration:   GetDateDurationByHour(v.Date[0]),
				NotifyList: v.Mentions,
				Date:       v.Date,
				Status:     global.ProcessStart,
			},
			mvc_struct.ProcessStepSchedule{
				Name:       processList[i],
				Step:       v.Step,
				Duration:   GetDateDurationByHour(v.Date[1]),
				NotifyList: v.Mentions,
				Date:       v.Date,
				Status:     global.ProcessEnd,
			})
	}
	mentions = removeDuplicationMap(mentions)
	service.HandleProcessTask(processTasks, mentions, processId)
}

// removeDuplicationMap 去重
func removeDuplicationMap(arr []string) []string {
	set := make(map[string]struct{}, len(arr))
	j := 0
	for _, v := range arr {
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
		arr[j] = v
		j++
	}

	return arr[:j]
}
