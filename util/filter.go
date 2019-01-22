package util

import (
	"github.com/astaxie/beego"
	"reflect"
)

func FilterByParam(controller *beego.Controller, name string, queryBuilder interface{}, methodName string, single bool) {
	params := controller.GetStrings(name)
	if params == nil {
		return
	}
	builderRef := reflect.ValueOf(queryBuilder)
	filterMethodRef := builderRef.MethodByName(methodName)
	inputs := make([]reflect.Value, len(params))
	if single {
		inputs[0] = reflect.ValueOf(params[0])
	} else {
		for i := range params {
			inputs[i] = reflect.ValueOf(params[i])
		}

	}
	filterMethodRef.Call(inputs)

}
