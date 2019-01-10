package cart

import (
	"github.com/astaxie/beego"
	"strconv"
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/parser"
	"yougame.com/yougame-server/security"
	"yougame.com/yougame-server/serializer"
	"yougame.com/yougame-server/service"
	"yougame.com/yougame-server/util"
)

type ApiCartController struct {
	api.ApiController
}

func (c ApiCartController) GetCartList() {
	var err error
	defer api.CheckError(func(e error) {
		beego.Debug(e)
		api.HandleApiError(c.Controller, e)
	})
	claims, err := c.GetAuth()
	if err != nil {
		panic(security.ReadAuthorizationFailed)
	}
	if claims == nil {
		panic(security.ReadAuthorizationFailed)
	}

	cartUserId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		panic(err)
	}

	permissionContext := map[string]interface{}{
		"claims":     *claims,
		"cartUserId": cartUserId,
	}
	permissions := []api.PermissionInterface{
		GetOtherCartPermission{},
		GetSelfCartPermission{},
	}
	err = c.CheckPermission(permissions, permissionContext)
	if err != nil {
		panic(err)
	}
	queryBuilder := service.CartQueryBuilder{}
	page, pageSize := c.GetPage()
	queryBuilder.SetPage(page, pageSize)
	user, err := models.GetUserById(claims.UserId)
	if err != nil {
		panic(security.ReadAuthorizationFailed)
	}
	queryBuilder.InUser(user.Id)
	count, cartItems, err := queryBuilder.Query()
	if err != nil {
		panic(err)
	}

	results := serializer.SerializeMultipleTemplate(cartItems, serializer.NewCartTemplate(serializer.DefaultCartTemplateType), map[string]interface{}{
		"site": util.GetSiteAndPortUrl(c.Controller),
	})

	c.ServerPageResult(results, count, page, pageSize)

}

func (c *ApiCartController) Create() {
	c.WithErrorContext(func() {
		createView := api.CreateView{
			Controller:    &c.ApiController,
			Parser:        &parser.CreateCartRequestBody{},
			Model:         &models.CartItem{},
			ModelTemplate: serializer.NewCartTemplate(serializer.DefaultCartTemplateType),
			OnPrepareSave: func(c *api.CreateView) {
				dataModel := c.Model.(*models.CartItem)
				requestParser := c.Parser.(*parser.CreateCartRequestBody)
				dataModel.Good = &models.Good{
					Id: int(requestParser.GoodId),
				}
				dataModel.User = c.Controller.User
				dataModel.Enable = true
			},
			Validate: func(v *api.CreateView) {
				requestParser := v.Parser.(*parser.CreateCartRequestBody)
				checkDuplicateCartItemValidator := DuplicateCartItemValidator{}
				result := checkDuplicateCartItemValidator.Check(map[string]interface{}{
					"userId": int64(c.User.Id),
					"goodId": requestParser.GoodId,
				})
				if !result {
					panic(api.InvalidateError)
				}
			},
		}
		err := createView.Exec()
		if err != nil {
			panic(err)
		}

	})
}
