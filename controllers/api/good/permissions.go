package good

import (
	"github.com/sirupsen/logrus"
	"yougame.com/yougame-server/security"
)

const (
	UpdateGoodPermissionName = "UpdateGood"
	CreateGoodPermissionName = "CreateGood"
	DeleteGoodGoodPermissionName = "DeleteGood"
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

type DeleteGoodPermission struct {
}

func (p *DeleteGoodPermission) CheckPermission(context map[string]interface{}) bool {
	claims := context["claims"].(*security.UserClaims)
	if err := security.CheckClaimsPermission(*claims, DeleteGoodGoodPermissionName); err != nil {
		logrus.Error(err)
		return false
	} else {
		return true
	}
}
type CreateGoodPermission struct {

}

func (p *CreateGoodPermission) CheckPermission(context map[string]interface{}) bool {
	claims := context["claims"].(*security.UserClaims)
	if err := security.CheckClaimsPermission(*claims, CreateGoodPermissionName); err != nil {
		logrus.Error(err)
		return false
	} else {
		return true
	}
}
