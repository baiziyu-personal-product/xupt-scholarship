package model

import (
	"time"
)

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
	currentLocation := time.Now().Location()
	timeLayout := "2006-01-02"
	start, _ := time.ParseInLocation(timeLayout, startDate, currentLocation)
	return start.Sub(time.Now()).Hours() < 24
}
