package cart

import (
	"yougame.com/yougame-server/security"
)

type GetOtherCartPermission struct {
}

func (p GetOtherCartPermission) CheckPermission(context map[string]interface{}) bool {
	claims := context["claims"].(security.UserClaims)
	err := security.CheckClaimsPermission(claims, "GetOtherCart")
	if err != nil {
		return false
	}
	return true

}

type GetSelfCartPermission struct {
}

func (p GetSelfCartPermission) CheckPermission(context map[string]interface{}) bool {
	cartUserId := context["cartUserId"].(int)
	claims := context["claims"].(security.UserClaims)
	return cartUserId == claims.UserId
}
