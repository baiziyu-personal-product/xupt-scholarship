package service

import (
	"fmt"
	"time"
)

var ProcessTask []*time.Timer

func NewProcessSchedule(taskType string) {
	temp := time.NewTimer(5 * time.Second)
	go func() {
		<-temp.C
		fmt.Println(" 1 s 时间到 ", time.Now().Unix())
	}()

}
