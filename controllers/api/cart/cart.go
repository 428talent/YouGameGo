package cart

import (
	"yougame.com/yougame-server/controllers"
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/security"
	"yougame.com/yougame-server/serializer"
	"yougame.com/yougame-server/util"
)

type ApiCartController struct {
	controllers.ApiController
}

func (c ApiCartController) GetCartList() {
	var err error
	defer api.CheckError(func(e error) {
		api.HandleApiError(c.Controller, err)
	})
	claims, err := c.GetAuth()
	if err != nil {
		panic(security.ReadAuthorizationFailed)
	}
	if claims == nil {
		panic(security.ReadAuthorizationFailed)
	}
	

	page, pageSize := c.GetPage()

	user, err := models.GetUserById(claims.UserId)
	if err != nil {
		panic(security.ReadAuthorizationFailed)
	}

	err = user.ReadCart((page-1)*pageSize, pageSize, "-created")
	if err != nil {
		panic(err)
	}
	serializedCartList, err := serializer.SerializeCartList(user.ShoppingCart, serializer.CartSerializer{})
	if err != nil {
		panic(err)
	}
	c.Data["json"] = &util.PageResponse{
		Page:     page,
		PageSize: pageSize,
		Count:    int64(len(user.ShoppingCart)),
		Result:   serializedCartList,
	}
	c.ServeJSON()

}
