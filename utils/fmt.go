package utils

import (
	"time"
)

func FmtTimeByUnix(unixTime int64) string {
	return time.Unix(unixTime, 0).Format("2006-01-02 15:04:05")
}

func FmtDateByUnix(unixTime int64) string {
	return time.Unix(unixTime, 0).Format("2006-01-02")
}
