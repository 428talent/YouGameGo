package wishlist

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strconv"
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/serializer"
	"yougame.com/yougame-server/util"
)

type ApiWishListController struct {
	api.ApiController
}

func (c *ApiWishListController) GetWishList() {
	page, pageSize := c.GetPage()
	c.GetPage()
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
	var serializerDataList []interface{}
	for _, item := range wishlist {
		serializeData := &serializer.WishListModel{}
		serializeData.SerializeData(item, util.GetSiteAndPortUrl(c.Controller))
		intf := reflect.ValueOf(serializeData).Interface()
		serializerDataList = append(serializerDataList, intf)
	}

	c.ServerPageResult(serializerDataList,*count,page,pageSize)
}

func (c *ApiWishListController) DeleteWishList() {

}
