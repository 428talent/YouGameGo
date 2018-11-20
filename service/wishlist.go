package service

import (
	"github.com/astaxie/beego/orm"
	"yougame.com/yougame-server/models"
)

func GetWishList(filter map[string]interface{}, page int64, pageSize int64) (int64, []*models.WishList, error) {
	count, wishlist, err := models.GetWishList(func(o orm.QuerySeter) orm.QuerySeter {
		if userId, exist := filter["user"]; exist {
			o = o.Filter("user_id", userId.(int64))
		}
		return o.Limit(pageSize).Offset((page - 1) * pageSize)
	})
	return *count, wishlist, err
}
