package game

import (
	"github.com/sirupsen/logrus"
	"yougame.com/yougame-server/security"
)

const (
	CreateGamePermissionName = "CreateGame"
)

type CreateGamePermission struct {}

func (p CreateGamePermission) CheckPermission(context map[string]interface{}) bool {
	claims := context["claims"].(security.UserClaims)
	if err := security.CheckClaimsPermission(claims, CreateGamePermissionName); err != nil {
		logrus.Error(err)
		return false
	} else {
		return true
	}

}
