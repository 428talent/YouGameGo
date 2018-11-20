package serializer

import (
	"reflect"
)

const (
	RelUser = "user"
	RelGame = "game"
)

type Model interface {
	SerializeData(model interface{}, site string) interface{}
}

type CommonApiResponseBody struct {
	Success bool        `json:"success"`
	Payload interface{} `json:"payload"`
}

type ApiLink struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
	Type string `json:"type"`
}

func SerializeData(template Model, model interface{}, site string) interface{} {
	serializeModel := reflect.New(reflect.ValueOf(template).Type()).Interface().(Model)
	serializeModel.SerializeData(model, site)
	return serializeModel
}

func SerializeMultipleData(template Model, models []interface{}, site string) []interface{} {
	dataList := make([]interface{}, 0)
	for _, model := range models {
		dataList = append(dataList, template.SerializeData(model, site))
	}
	return dataList
}
