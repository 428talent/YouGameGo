package order

import "yougame.com/yougame-server/service"

type GoodValidate struct {
}

func (v *GoodValidate) Check(context map[string]interface{}) bool {
	goodList := context["goodList"].([]int64)
	goodQueryBuilder := service.GoodQueryBuilder{}
	for _, id := range goodList {
		goodQueryBuilder.InId(id)
	}
	goodQueryBuilder.WithEnable("visit")
	count, _, err := goodQueryBuilder.Query()
	if err != nil {
		return false
	}
	if int(*count) == len(goodList) {
		return true
	}
	return false

}
