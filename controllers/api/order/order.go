package order

import (
	"github.com/astaxie/beego"
	"github.com/sirupsen/logrus"
	"reflect"
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
	claims, err := security.ParseAuthHeader(c.Controller)
	if err != nil {
		beego.Error(err)
		return
	}
	if claims == nil {
		return
	}
	requestBodyStruct := parser.CreateOrderParser{}
	err = requestBodyStruct.Parse(c.Controller)
	if err != nil {
		beego.Error(err)
		return
	}
	var goodList []*models.OrderGood
	for _, goodId := range requestBodyStruct.Goods {
		goodList = append(goodList, &models.OrderGood{
			Good: &models.Good{
				Id: int(goodId),
			},
		})
	}
	order := models.Order{
		User:  &models.User{Id: claims.UserId},
		State: "Created",
		Goods: goodList,
	}
	err = order.SaveOrder()
	if err != nil {
		beego.Error(err)
	}
	c.ServeJSON()
}

func (c *ApiOrderController) GetOrderList() {
	var err error
	defer api.CheckError(func(e error) {
		logrus.Error(e)
		api.HandleApiError(c.Controller, e)
	})
	claims, err := security.ParseAuthHeader(c.Controller)
	if err != nil {
		panic(err)
	}
	if claims == nil {
		panic(security.ReadAuthorizationFailed)
	}
	//orderUserId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	//if err != nil {
	//	panic(err)
	//}
	page, pageSize := util.ParsePageRequest(c.Controller)
	//permissionContext := map[string]interface{}{
	//	"claims":      *claims,
	//	"orderUserId": orderUserId,
	//}
	//permissions := []api.ApiPermissionInterface{
	//	GetOwnOrderPermission{},
	//}
	//err = c.CheckPermission(permissions, permissionContext)
	//if err != nil {
	//	panic(api.PermissionDeniedError)
	//}
	//query filter
	builder := service.GetOrderListBuilder{}
	builder.SetPage(page).SetPageSize(pageSize)
	if userIdParam := c.GetString("user"); len(userIdParam) != 0 {
		userId, err := strconv.Atoi(userIdParam)
		if err == nil {
			builder.SetUser(int64(userId))
		}
	}
	if states := c.GetStrings("state"); len(states) > 0 {
		builder.SetState(states)
	}
	count, orders, err := service.GetOrderList(builder)
	if err != nil {
		panic(err)
	}
	serializerDataList := serializer.SerializeMultipleTemplate(
		orders,
		serializer.NewOrderTemplate(serializer.DefaultOrderTemplateType),
		map[string]interface{}{
			"site": util.GetSiteAndPortUrl(c.Controller),
		},
	)
	c.ServerPageResult(serializerDataList, count, page, pageSize)
}
func (c *ApiOrderController) GetOrderGoodsWithOrder() {
	var err error
	defer api.CheckError(func(e error) {
		logrus.Error(e)
		api.HandleApiError(c.Controller, e)
	})
	claims, err := security.ParseAuthHeader(c.Controller)
	if err != nil {
		panic(err)
	}
	if claims == nil {
		panic(security.ReadAuthorizationFailed)
	}
	//orderUserId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	//if err != nil {
	//	panic(err)
	//}
	page, pageSize := util.ParsePageRequest(c.Controller)
	//permissionContext := map[string]interface{}{
	//	"claims":      *claims,
	//	"orderUserId": orderUserId,
	//}
	//permissions := []api.ApiPermissionInterface{
	//	GetOwnOrderPermission{},
	//}
	//err = c.CheckPermission(permissions, permissionContext)
	//if err != nil {
	//	panic(api.PermissionDeniedError)
	//}
	//query filter
	option := service.GetOrderGoodListOption{}
	option.SetPage(service.PageOption{
		Page:     page,
		PageSize: pageSize,
	})

	if orderIdParam := c.Ctx.Input.Param(":id"); len(orderIdParam) != 0 {
		orderId, err := strconv.Atoi(orderIdParam)
		if err == nil {
			option.SetOrder(int64(orderId))
		}
	}
	var orderGoods []*models.OrderGood
	count, err := option.Query(&orderGoods)
	if err != nil {
		panic(err)
	}
	serializerDataList := serializer.SerializeMultipleTemplate(
		orderGoods,
		serializer.NewOrderGoodTemplate(serializer.DefaultOrderTemplateType),
		map[string]interface{}{
			"site": util.GetSiteAndPortUrl(c.Controller),
		},
	)
	c.ServerPageResult(serializerDataList, count, page, pageSize)
}
func (c *ApiOrderController) GetOrderGoods() {
	var err error
	defer api.CheckError(func(e error) {
		logrus.Error(e)
		api.HandleApiError(c.Controller, e)
	})
	claims, err := security.ParseAuthHeader(c.Controller)
	if err != nil {
		panic(err)
	}
	if claims == nil {
		panic(security.ReadAuthorizationFailed)
	}
	//orderUserId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	//if err != nil {
	//	panic(err)
	//}
	page, pageSize := util.ParsePageRequest(c.Controller)
	//permissionContext := map[string]interface{}{
	//	"claims":      *claims,
	//	"orderUserId": orderUserId,
	//}
	//permissions := []api.ApiPermissionInterface{
	//	GetOwnOrderPermission{},
	//}
	//err = c.CheckPermission(permissions, permissionContext)
	//if err != nil {
	//	panic(api.PermissionDeniedError)
	//}
	//query filter
	option := service.GetOrderGoodListOption{}
	option.SetPage(service.PageOption{
		Page:     page,
		PageSize: pageSize,
	})

	if orderIdParam := c.GetString("order"); len(orderIdParam) != 0 {
		orderId, err := strconv.Atoi(orderIdParam)
		if err == nil {
			option.SetOrder(int64(orderId))
		}
	}
	var orderGoods []*models.OrderGood
	count, err := option.Query(&orderGoods)
	if err != nil {
		panic(err)
	}
	results := make([]interface{}, 0)
	for _, item := range orderGoods {
		results = append(results, reflect.ValueOf(*item).Interface())
	}
	serializerDataList := serializer.SerializeMultipleTemplate(
		orderGoods,
		serializer.NewOrderGoodTemplate(serializer.DefaultOrderTemplateType),
		map[string]interface{}{
			"site": util.GetSiteAndPortUrl(c.Controller),
		},
	)
	c.ServerPageResult(serializerDataList, count, page, pageSize)
}
func (c *ApiOrderController) PayOrder() {
	var err error
	defer api.CheckError(func(e error) {
		logrus.Error(e)
		switch e {
		case service.NotSufficientFundsError:
			NotSufficientFunds.ServerError(c.Controller)
			return
		case service.WrongOrderStateError:
			WrongOrderState.ServerError(c.Controller)
			return
		default:
			api.HandleApiError(c.Controller, e)
			return
		}
	})
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
		beego.Error(err)
	}
	err = service.PayOrder(order)
	if err != nil {
		panic(err)
	}

	c.Data["json"] = &serializer.CommonApiResponseBody{
		Success: true,
	}
	c.ServeJSON()
}
