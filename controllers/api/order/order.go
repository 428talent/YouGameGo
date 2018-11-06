package order

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/sirupsen/logrus"
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
	orderUserId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		panic(err)
	}
	page, pageSize := util.ParsePageRequest(c.Controller)
	user, err := models.GetUserById(claims.UserId)
	if err != nil {
		panic(err)
	}
	permissionContext := map[string]interface{}{
		"claims":      *claims,
		"orderUserId": orderUserId,
	}
	permissions := []api.ApiPermissionInterface{
		GetOwnOrderPermission{},
	}
	err = c.CheckPermission(permissions, permissionContext)
	if err != nil {
		panic(api.PermissionDeniedError)
	}
	//query filter

	orders, err := models.GetOrderList(func(o orm.QuerySeter) orm.QuerySeter {
		cond := orm.NewCondition().And("user_id", user.Id)
		stateFilter := c.GetStrings("state")
		if len(stateFilter) > 0 {
			stateCond := orm.NewCondition()
			for _, state := range stateFilter {
				stateCond  = stateCond.Or("state",state)
			}
			cond = cond.AndCond(stateCond)
		}
		orderId := c.GetString("orderId")
		if len(orderId) >0 {
			cond = cond.And("id",orderId)
		}

		o = o.SetCond(cond)
		return o
	})
	if err != nil {
		panic(err)
	}
	serializedData, err := serializer.SerializeOrderList(orders, serializer.OrderSerializer{})
	if err != nil {
		panic(err)
	}
	c.Data["json"] = util.PageResponse{
		Page:     page,
		PageSize: pageSize,
		Result:   serializedData,
		Count:    int64(len(serializedData)),
	}
	c.ServeJSON()
}

func (c *ApiOrderController) PayOrder() {
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

	defer func() {
		troubleMaker := recover()
		if troubleMaker != nil {
			err = troubleMaker.(error)
			switch err {
			case service.NotSufficientFundsError:
				NotSufficientFunds.ServerError(c.Controller)
			case service.WrongOrderStateError:
				WrongOrderState.ServerError(c.Controller)
			default:
				api.HandleApiError(c.Controller, err)
			}
		} else {
			c.Data["json"] = &serializer.CommonApiResponseBody{
				Success: true,
			}
			c.ServeJSON()
		}

	}()
}
