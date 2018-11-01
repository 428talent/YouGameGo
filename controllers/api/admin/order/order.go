package order

import (
	"github.com/astaxie/beego"
	"strconv"
	"yougame.com/letauthsdk/auth"
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/parser"
	"yougame.com/yougame-server/security"
	"yougame.com/yougame-server/serializer"
	"yougame.com/yougame-server/service"
	"yougame.com/yougame-server/util"
)

type ApiOrderController struct {
	beego.Controller
}

func (c *ApiOrderController) CreateOrder() {
	claims, err := auth.ParseAuthHeader(c.Controller, security.AppSecret)
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
	claims, err := auth.ParseAuthHeader(c.Controller, security.AppSecret)
	if err != nil {
		beego.Error(err)
		return
	}
	if claims == nil {
		return
	}
	page, pageSize := util.ParsePageRequest(c.Controller)
	user, err := models.GetUserById(claims.UserId)
	if err != nil {
		beego.Error(err)
		return
	}
	if err = user.ReadOrders((page-1)*pageSize, pageSize, "-created"); err != nil {
		beego.Error(err)
	}
	serializedData, err := serializer.SerializeOrderList(user.Orders, serializer.OrderSerializer{})
	if err != nil {
		beego.Error(err)
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
	claims, err := auth.ParseAuthHeader(c.Controller, security.AppSecret)
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
