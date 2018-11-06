package order

import (
	"yougame.com/yougame-server/security"
)

const (
	GetOwnOrderPermissionName   = "GetOwnOrderOrder"
	GetOtherOrderPermissionName = "GetOtherOrderPermission"
)

type GetOwnOrderPermission struct{}

func (p GetOwnOrderPermission) CheckPermission(context map[string]interface{}) bool {
	claims := context["claims"].(security.UserClaims)
	orderUserId := context["orderUserId"].(int)
	if orderUserId != claims.UserId {
		if err := security.CheckClaimsPermission(claims, GetOwnOrderPermissionName); err != nil {
			return false
		} else {
			return true
		}
	} else {
		if err := security.CheckClaimsPermission(claims, GetOtherOrderPermissionName); err != nil {
			return false
		} else {
			return true
		}
	}
}
