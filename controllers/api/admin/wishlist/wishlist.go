package wishlist

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/serializer"
	"yougame.com/yougame-server/util"
)

type ApiWishListController struct {
	beego.Controller
}

func (c *ApiWishListController) GetWishList() {
	page, pageSize := util.ParsePageRequest(c.Controller)
	userId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		beego.Error(err)
	}
	count, wishlist, err := models.GetWishList(func(o orm.QuerySeter) orm.QuerySeter {
		return o.Filter("user_id", userId).Limit(pageSize).Offset((page - 1) * pageSize)
	})
	if err != nil {
		beego.Error(err)
	}

	results := serializer.SerializeWishListMultiple(wishlist, serializer.WishListSerializer{})
	response := util.PageResponse{
		PageSize: pageSize,
		Page:     page,
		Count:    *count,
		Result:   results,
	}
	c.Data["json"] = response
	c.ServeJSON()
}
