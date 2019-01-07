package image

import (
	"github.com/sirupsen/logrus"
	"yougame.com/yougame-server/security"
)

const (
	DeleteImagePermissionName = "DeleteImage"
	UpdateImagePermissionName = "UpdateImage"
)

type DeleteImagePermission struct {
}

func (p *DeleteImagePermission) CheckPermission(context map[string]interface{}) bool {
	claims := context["claims"].(*security.UserClaims)
	if err := security.CheckClaimsPermission(*claims, DeleteImagePermissionName); err != nil {
		logrus.Error(err)
		return false
	} else {
		return true
	}
}

type UpdateImagePermission struct {
}

func (p *UpdateImagePermission) CheckPermission(context map[string]interface{}) bool {
	claims := context["claims"].(*security.UserClaims)
	if err := security.CheckClaimsPermission(*claims, UpdateImagePermissionName); err != nil {
		logrus.Error(err)
		return false
	} else {
		return true
	}
}
