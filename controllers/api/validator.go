package api

type RequestValidator interface {
	Check(context map[string]interface{}) bool
	GetMessage() string
}