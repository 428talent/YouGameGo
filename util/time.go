package util

import "time"

func FormatApiTime(timeValue time.Time) string {
	return timeValue.Format(time.RFC3339)
}

func FormatDate(timeValue time.Time) string {
	return timeValue.Format("2006-1-2")
}

func ParseDate(timeString string) (time.Time,error) {
	return time.Parse("2006-1-2",timeString)
}