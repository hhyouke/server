package utils

import "time"

var (
	timeFormat = "2006-01-02 15:04:05"
	defaultTZ  = "Local"
)

// PtrTime returns the pointer of time.Time
func PtrTime(t time.Time) *time.Time {
	return &t
}

// FormattedLocalNow output formatted local now time
func FormattedLocalNow() string {
	loc, _ := time.LoadLocation(defaultTZ)
	return time.Now().In(loc).Format(timeFormat)
}
