package model

import (
	"reflect"
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
	timeLayout := "2006-01-02"
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
func DispatchProcessNoticeEvent(processInfo mvc_struct.ProcessFormData) {
	var processTasks []mvc_struct.ProcessTask
	processTypes := reflect.TypeOf(processInfo.Form)
	processValues := reflect.ValueOf(processInfo.Form)
	stepMap := reflect.ValueOf(global.ProcessStepMap)
	stepKey := reflect.TypeOf(global.ProcessStepMap)
	var mentions []string
	for i := 0; i < processTypes.NumField(); i++ {
		stepValue := processValues.Field(i).Interface().(mvc_struct.ProcessStepValue)
		mentions = append(mentions, stepValue.Mentions...)
		processTasks = append(
			processTasks,
			mvc_struct.ProcessTask{
				Name:       stepMap.Field(i).String(),
				Step:       stepKey.Field(i).Name,
				Duration:   GetDateDurationByHour(stepValue.Date[0]),
				NotifyList: stepValue.Mentions,
				Date:       stepValue.Date,
				Type:       "start",
			},
			mvc_struct.ProcessTask{
				Name:       stepMap.Field(i).String(),
				Step:       stepKey.Field(i).Name,
				Duration:   GetDateDurationByHour(stepValue.Date[1]),
				NotifyList: stepValue.Mentions,
				Date:       stepValue.Date,
				Type:       "end",
			})
	}
	mentions = removeDuplicationMap(mentions)
	service.HandleProcessTask(processTasks, mentions)
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
