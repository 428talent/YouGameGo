package cart

import (
	"github.com/astaxie/beego"
	"yougame.com/letauthsdk/auth"
	ApiError "yougame.com/yougame-server/error"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/security"
	"yougame.com/yougame-server/serializer"
	"yougame.com/yougame-server/util"
)

type ApiCartController struct {
	beego.Controller
}

func (c ApiCartController) GetCartList() {
	claims, err := auth.ParseAuthHeader(c.Controller, security.AppSecret)
	if err != nil {
		beego.Error(err)
		ApiError.ServerNoAuthError(c.Controller)
		return
	}
	if claims == nil {
		ApiError.ServerNoAuthError(c.Controller)
		return
	}
	page, pageSize := util.ParsePageRequest(c.Controller)

	user, err := models.GetUserById(claims.UserId)
	if err != nil {
		beego.Error(err)
	}
	err = user.ReadCart((page-1)*pageSize, pageSize, "-created")
	if err != nil {
		beego.Error(err)
	}
	serializedCartList, err := serializer.SerializeCartList(user.ShoppingCart, serializer.CartSerializer{})
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = &util.PageResponse{
		Page:     page,
		PageSize: pageSize,
		Count:    int64(len(user.ShoppingCart)),
		Result:   serializedCartList,
	}
	c.ServeJSON()

}
