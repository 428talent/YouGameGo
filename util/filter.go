package util

import (
	"github.com/astaxie/beego"
	"reflect"
)

func FilterByParam(controller *beego.Controller, name string, queryBuilder interface{}, methodName string) {
	params := controller.GetStrings(name)
	builderRef := reflect.ValueOf(queryBuilder)
	filterMethodRef := builderRef.MethodByName(methodName)
	inputs := make([]reflect.Value, len(params))
	for i := range params {
		inputs[i] = reflect.ValueOf(params[i])
	}
	filterMethodRef.Call(inputs)
}
