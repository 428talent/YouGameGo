package good

import (
	"github.com/sirupsen/logrus"
	"yougame.com/yougame-server/security"
)

const (
	UpdateGoodPermissionName = "UpdateGood"
)

type UpdateGoodPermission struct {
}

func (p *UpdateGoodPermission) CheckPermission(context map[string]interface{}) bool {
	claims := context["claims"].(*security.UserClaims)
	if err := security.CheckClaimsPermission(*claims, UpdateGoodPermissionName); err != nil {
		logrus.Error(err)
		return false
	} else {
		return true
	}
}
