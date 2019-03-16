package tag

import (
	"github.com/sirupsen/logrus"
	"yougame.com/yougame-server/security"
)

const (
	CreateTagPermissionName = "CreateTag"
	DeleteTagPermissionName = "DeleteTag"
	UpdateTagPermissionName = "UpdateTag"
)

type CreateTagPermission struct {
}

func (p *CreateTagPermission) CheckPermission(context map[string]interface{}) bool {
	claims := context["claims"].(*security.UserClaims)
	if err := security.CheckClaimsPermission(*claims, CreateTagPermissionName); err != nil {
		logrus.Error(err)
		return false
	} else {
		return true
	}
}

type DeleteTagPermission struct {
}

func (p *DeleteTagPermission) CheckPermission(context map[string]interface{}) bool {
	claims := context["claims"].(*security.UserClaims)
	if err := security.CheckClaimsPermission(*claims, DeleteTagPermissionName); err != nil {
		logrus.Error(err)
		return false
	} else {
		return true
	}
}

type UpdateTagPermission struct {
}

func (p *UpdateTagPermission) CheckPermission(context map[string]interface{}) bool {
	claims := context["claims"].(*security.UserClaims)
	if err := security.CheckClaimsPermission(*claims, UpdateTagPermissionName); err != nil {
		logrus.Error(err)
		return false
	} else {
		return true
	}
}
