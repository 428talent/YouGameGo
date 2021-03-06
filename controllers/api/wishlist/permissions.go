package wishlist

import (
	"github.com/astaxie/beego"
	"yougame.com/yougame-server/security"
	"yougame.com/yougame-server/service"
)

type DeleteWishlistPermission struct{}

func (p DeleteWishlistPermission) DeleteWishlistPermission(context map[string]interface{}) bool {
	claims := context["claims"].(security.UserClaims)
	ids := context["ids"].([]int)
	beego.Debug(ids)
	queryBuilder := service.WishListQueryBuilder{}
	queryBuilder.BelongToUser(claims.UserId)
	for _, id := range ids {
		queryBuilder.InId(id)
	}
	count, _, err := queryBuilder.GetWishList()
	if err != nil {
		return false
	}
	beego.Debug()
	if int(count) != len(ids) {
		return false
	}
	return true
}

type DeleteWishlistItemPermission struct {
}

func (p *DeleteWishlistItemPermission) CheckPermission(context map[string]interface{}) bool {
	claims := context["claims"].(*security.UserClaims)
	id := context["id"].(int)
	queryBuilder := service.WishListQueryBuilder{}
	queryBuilder.BelongToUser(claims.UserId)
	queryBuilder.InId(id)
	count, _, _ := queryBuilder.GetWishList()
	if count == 0 {
		return false
	}
	return true
}
