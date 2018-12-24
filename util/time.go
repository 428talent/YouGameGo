package util

import "time"

func FormatApiTime(timeValue time.Time) string {
	return timeValue.Format(time.RFC3339)
}
