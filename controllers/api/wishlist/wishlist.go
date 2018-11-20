package wishlist

import (
	"reflect"
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/serializer"
	"yougame.com/yougame-server/service"
	"yougame.com/yougame-server/util"
)

type ApiWishListController struct {
	api.ApiController
}

func (c *ApiWishListController) GetWishList() {
	var err error
	defer api.CheckError(func(e error) {
		api.HandleApiError(c.Controller, e)
	})

	page, pageSize := c.GetPage()
	queryContext := make(map[string]interface{})
	userId, _ := c.GetInt64("user", 0)
	if userId != 0 {
		queryContext["user"] = userId
	}
	count, wishlist, err := service.GetWishList(queryContext, page, pageSize)
	if err != nil {
		panic(err)
	}

	results := make([]interface{},0)
	for _, item := range wishlist {
		results = append(results, reflect.ValueOf(*item).Interface())
	}
	serializerDataList := serializer.SerializeMultipleData(&serializer.WishListModel{},results,util.GetSiteAndPortUrl(c.Controller))
	c.ServerPageResult(serializerDataList, count, page, pageSize)
}

func (c *ApiWishListController) DeleteWishList() {

}
