package cart

import (
	"github.com/sirupsen/logrus"
	"yougame.com/yougame-server/service"
)

type DuplicateCartItemValidator struct {
}

func (v *DuplicateCartItemValidator) Check(context map[string]interface{}) bool {
	userId := context["userId"].(int64)
	goodId := context["goodId"].(int64)

	cartQueryBuilder := service.CartQueryBuilder{}
	cartQueryBuilder.InUser(userId)
	cartQueryBuilder.WithGood(goodId)
	cartQueryBuilder.WithEnable("visit")
	count, _, err := cartQueryBuilder.Query()
	if err != nil {
		logrus.Debug(err)
		return false
	}
	if count > 0 {
		return false
	}
	return true

}
