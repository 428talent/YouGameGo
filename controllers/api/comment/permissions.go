package comment

import (
	"github.com/sirupsen/logrus"
	"yougame.com/yougame-server/security"
)

const CreateCommentPermissionCode = "CreateComment"

type CreateCommentPermission struct{}

func (p CreateCommentPermission) CheckPermission(context map[string]interface{}) bool {
	claims := context["claims"].(security.UserClaims)
	if err := security.CheckClaimsPermission(claims, CreateCommentPermissionCode); err != nil {
		logrus.Debug(err)
		return false
	} else {
		return true
	}

}
