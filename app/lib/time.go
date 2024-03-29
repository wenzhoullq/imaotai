package lib

import (
	"time"
)

func GetCurrentDayTime() int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().In(loc)
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	dayTime := startOfDay.UnixNano() / 1e6
	return dayTime
}

func Overdue(t int64) bool {
	t1 := time.Unix(t, 0)
	t2 := time.Now()
	if t1.Before(t2) {
		return true
	}
	return false
}
