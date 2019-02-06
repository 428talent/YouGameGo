package log

import "time"

type LogPayload struct {
	Message string `json:"message"`
	Level   string `json:"level"`
	Time    time.Time  `json:"log_time"`
}
