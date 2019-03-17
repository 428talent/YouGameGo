package tag

import (
	"yougame.com/yougame-server/parser"
	"yougame.com/yougame-server/service"
)

type DuplicateTagValidator struct {

}

func (v *DuplicateTagValidator) Check(context map[string]interface{}) bool {
	requestParser := context["parser"].(*parser.CreateTagRequestBody)
	builder := service.TagQueryBuilder{}
	builder.WithName(requestParser.Name)
	count, _, err := builder.Query()
	if err != nil {
		return false
	} else {
		if *count != 0 {
			return false
		}
	}
	return true
}


func (v *DuplicateTagValidator) GetMessage() string {
	return "Duplicate tag name"
}
