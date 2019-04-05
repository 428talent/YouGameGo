package order

import (
	"strconv"
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/parser"
	"yougame.com/yougame-server/security"
	"yougame.com/yougame-server/serializer"
	"yougame.com/yougame-server/service"
	"yougame.com/yougame-server/util"
)

type ApiOrderController struct {
	api.ApiController
}

func (c *ApiOrderController) CreateOrder() {
	c.WithErrorContext(func() {
		createView := api.CreateView{
			Controller:    &c.ApiController,
			Parser:        &parser.CreateOrderParser{},
			Model:         &models.Order{},
			ModelTemplate: serializer.NewOrderTemplate(serializer.DefaultOrderTemplateType),
			Validate: func(v *api.CreateView) {
				request := v.Parser.(*parser.CreateOrderParser)
				if request.Goods != nil {
					goodValidator := GoodValidate{}
					isValid := goodValidator.Check(map[string]interface{}{
						"goodList": request.Goods,
					})
					if !isValid {
						panic(api.InvalidateError)
					}
				}
			},
			OnSave: func(v *api.CreateView) error {
				request := v.Parser.(*parser.CreateOrderParser)
				goodIds := make([]int64, 0)
				if request.UserCart != 0 {
					cartQueryBuilder := service.CartQueryBuilder{}
					cartQueryBuilder.InUser(request.UserCart)
					cartQueryBuilder.WithEnable("visit")
					_, cartItems, err := cartQueryBuilder.Query()
					if err != nil {
						panic(err)
					}
					for _, cartItem := range cartItems {
						goodIds = append(goodIds, int64(cartItem.Good.Id))
					}

				} else {
					goodIds = request.Goods
				}
				orderModel := v.Model.(*models.Order)
				orderModel.User = c.User
				err := service.CreateOrder(orderModel, goodIds)
				if err != nil {
					return err
				}
				return nil
			},
		}
		err := createView.Exec()
		if err != nil {
			panic(err)
		}
	})

}

func (c *ApiOrderController) GetOrderList() {
	c.WithErrorContext(func() {
		listView := api.ListView{
			Controller:    &c.ApiController,
			QueryBuilder:  &service.OrderQueryBuilder{},
			ModelTemplate: serializer.NewOrderTemplate(serializer.DefaultOrderTemplateType),
			Init: func() {
				c.GetAuth()
			},
			SetFilter: func(builder service.ApiQueryBuilder) {
				orderQueryBuilder := builder.(*service.OrderQueryBuilder)
				if security.CheckUserGroup(c.User, security.UserGroupAdmin) {
					util.FilterByParam(&c.Controller, "user", builder, "InUser", false)
					util.FilterByParam(&c.Controller, "enable", builder, "WithEnable", true)
				} else {
					orderQueryBuilder.InUser(c.User.Id)
					orderQueryBuilder.WithEnable("visit")
				}
				util.FilterByParam(&c.Controller, "order", builder, "ByOrder", false)
				util.FilterByParam(&c.Controller, "state", builder, "SetState", false)
				util.FilterByParam(&c.Controller, "id", builder, "InId", false)
			},
		}
		err := listView.Exec()
		if err != nil {
			panic(err)
		}
	})
}

func (c *ApiOrderController) GetOrder() () {
	c.WithErrorContext(func() {
		objectView := api.ObjectView{
			Controller:    &c.ApiController,
			QueryBuilder:  &service.OrderQueryBuilder{},
			ModelTemplate: serializer.NewOrderTemplate(serializer.DefaultOrderTemplateType),
		}

		err := objectView.Exec()
		if err != nil {
			panic(err)
		}
	})
}
func (c *ApiOrderController) GetOrderGoodsWithOrder() {
	//var err error
	//defer api.CheckError(func(e error) {
	//	logrus.Error(e)
	//	api.HandleApiError(c.Controller, e)
	//})
	//claims, err := security.ParseAuthHeader(c.Controller)
	//if err != nil {
	//	panic(err)
	//}
	//if claims == nil {
	//	panic(security.ReadAuthorizationFailed)
	//}
	////orderUserId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	////if err != nil {
	////	panic(err)
	////}
	//page, pageSize := util.ParsePageRequest(c.Controller)
	////permissionContext := map[string]interface{}{
	////	"claims":      *claims,
	////	"orderUserId": orderUserId,
	////}
	////permissions := []api.ApiPermissionInterface{
	////	GetOwnOrderPermission{},
	////}
	////err = c.CheckPermission(permissions, permissionContext)
	////if err != nil {
	////	panic(api.PermissionDeniedError)
	////}
	////query filter
	//option := service.OrderGoodQueryBuilder{}
	//option.SetPage(service.PageOption{
	//	Page:     page,
	//	PageSize: pageSize,
	//})
	//
	//if orderIdParam := c.Ctx.Input.Param(":id"); len(orderIdParam) != 0 {
	//	orderId, err := strconv.Atoi(orderIdParam)
	//	if err == nil {
	//		option.SetOrder(int64(orderId))
	//	}
	//}
	//var orderGoods []*models.OrderGood
	//count, err := option.Query(&orderGoods)
	//if err != nil {
	//	panic(err)
	//}
	//serializerDataList := serializer.SerializeMultipleTemplate(
	//	orderGoods,
	//	serializer.NewOrderGoodTemplate(serializer.DefaultOrderTemplateType),
	//	map[string]interface{}{
	//		"site": util.GetSiteAndPortUrl(c.Controller),
	//	},
	//)
	//c.ServerPageResult(serializerDataList, count, page, pageSize)
}
func (c *ApiOrderController) GetOrderGoods() {
	c.WithErrorContext(func() {
		listView := api.ListView{
			Controller:    &c.ApiController,
			QueryBuilder:  &service.OrderGoodQueryBuilder{},
			ModelTemplate: serializer.NewOrderGoodTemplate(serializer.DefaultOrderGoodTemplateType),
			SetFilter: func(builder service.ApiQueryBuilder) {
				util.FilterByParam(&c.Controller, "orderId", builder, "WithOrderId", false)

			},
		}
		err := listView.Exec()
		if err != nil {
			panic(err)
		}
	})
}
func (c *ApiOrderController) PayOrder() {
	c.WithErrorContext(func() {
		claims, err := security.ParseAuthHeader(c.Controller)
		if err != nil {
			panic(err)

		}
		if claims == nil {
			panic(service.NoAuthError)
		}
		orderId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
		if err != nil {
			panic(err)
		}
		order := models.Order{Id: orderId}
		if err = order.QueryById(); err != nil {
			panic(err)
		}
		err = service.PayOrder(order)
		if err != nil {
			panic(err)
		}

		c.Data["json"] = &serializer.CommonApiResponseBody{
			Success: true,
		}
		c.ServeJSON()
	})
	//var err error
	//defer api.CheckError(func(e error) {
	//	logrus.Error(e)
	//	switch e {
	//	case service.NotSufficientFundsError:
	//		NotSufficientFunds.ServerError(c.Controller)
	//		return
	//	case service.WrongOrderStateError:
	//		WrongOrderState.ServerError(c.Controller)
	//		return
	//	default:
	//		api.HandleApiError(c.Controller, e)
	//		return
	//	}
	//})

}
