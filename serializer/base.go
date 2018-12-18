package serializer

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"
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

type Template interface {
	CustomSerialize(convertTag string, value interface{}) interface{}
	Serialize(model interface{},context map[string]interface{})
}




func SerializeModelData(data interface{}, template interface{}) interface{} {
	dataRef := reflect.ValueOf(data).Elem()
	templateRef := reflect.ValueOf(template).Elem()
	for fieldIdx := 0; fieldIdx < templateRef.NumField(); fieldIdx++ {
		sourceName := reflect.TypeOf(template).Elem().Field(fieldIdx).Tag.Get("source")
		if len(sourceName) == 0 {
			sourceName = reflect.TypeOf(template).Elem().Field(fieldIdx).Name
		}
		sourceType := reflect.TypeOf(template).Elem().Field(fieldIdx).Tag.Get("source_type")
		dataFieldValue := dataRef
		for _, fieldString := range strings.Split(sourceName, ".") {
			matchMethodRegex, _ := regexp.Compile(`(.*?)\(\)\[(\d+)\]$`)
			if matchMethodRegex.MatchString(fieldString) {
				methodName := matchMethodRegex.FindStringSubmatch(fieldString)[1]
				returnValueIndex, _ := strconv.Atoi(matchMethodRegex.FindStringSubmatch(fieldString)[2])
				if dataFieldValue.Kind() == reflect.Ptr {
					dataFieldValue = dataFieldValue.Elem()
				}
				param := make([]reflect.Value, 0)
				returnValues := dataFieldValue.MethodByName(methodName).Call(param)
				dataFieldValue = returnValues[returnValueIndex]

			} else {
				if dataFieldValue.Kind() == reflect.Ptr {
					dataFieldValue = dataFieldValue.Elem()
				}
				dataFieldValue = dataFieldValue.FieldByName(fieldString)
			}

			//check it method call

		}
		converter := reflect.TypeOf(template).Elem().Field(fieldIdx).Tag.Get("converter")
		if len(converter) > 0 {
			customTemplate := template.(Template)
			result := customTemplate.CustomSerialize(converter, dataFieldValue.Interface())
			dataFieldValue = reflect.ValueOf(result)
		}
		switch sourceType {
		case "int":
			templateRef.Field(fieldIdx).SetInt(dataFieldValue.Int())
		case "string":
			templateRef.Field(fieldIdx).SetString(dataFieldValue.String())
		case "float":
			templateRef.Field(fieldIdx).SetFloat(dataFieldValue.Float())
		}

	}

	return nil
}

func SerializeMultipleTemplate(items interface{},template Template,context map[string]interface{})interface{}{
	result := make([]interface{},0)
	itemListRef := reflect.ValueOf(items)
	for itemIdx := 0; itemIdx < itemListRef.Len();itemIdx++  {
		itemTemplate := reflect.New(reflect.TypeOf(template).Elem())
		tmp := itemTemplate.Interface().(Template)
		tmp.Serialize(itemListRef.Index(itemIdx).Interface(),context)
		result = append(result, tmp)
	}
	return result
}