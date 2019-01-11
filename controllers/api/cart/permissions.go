package cart

import (
	"yougame.com/yougame-server/security"
	"yougame.com/yougame-server/service"
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

type DeleteCartPermission struct {
}

func (p *DeleteCartPermission) CheckPermission(context map[string]interface{}) bool {
	cartId := context["cartId"].(int)
	claims := context["claims"].(*security.UserClaims)

	cartQuery := service.CartQueryBuilder{}
	cartQuery.InId(cartId)
	cartQuery.InUser(claims.UserId)
	count, _, _ := cartQuery.Query()
	if count < 1 {
		return false
	}
	return true
}
