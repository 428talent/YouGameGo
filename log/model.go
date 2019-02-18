package log

import "time"

type Payload struct {
	Message     string    `json:"message"`
	Level       string    `json:"level"`
	Time        time.Time `json:"log_time"`
	Application string    `json:"application"`
	Instance    string    `json:"instance"`
}

func NewPayload(message string, level string, time time.Time) *Payload {
	return &Payload{Message: message, Level: level, Time: time}
}
