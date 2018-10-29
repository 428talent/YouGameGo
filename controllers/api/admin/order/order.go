package order

import (
	"github.com/astaxie/beego"
	"strconv"
	"yougame.com/letauthsdk/auth"
	ApiError "yougame.com/yougame-server/error"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/parser"
	"yougame.com/yougame-server/security"
	"yougame.com/yougame-server/serializer"
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
		beego.Error(err)

	}
	if claims == nil {

	}
	orderId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		beego.Error(err)

	}
	order := models.Order{Id: orderId}
	if err = order.QueryById(); err != nil {
		beego.Error(err)

	}
	user, err := models.GetUserById(claims.UserId)
	if err != nil {
		beego.Error(err)
		return
	}
	if err = order.ReadOrderGoods(); err != nil {
		beego.Error(err)
		return
	}
	totalPrice := 0.0
	for _, orderGood := range order.Goods {
		totalPrice += orderGood.Price
	}
	if err = user.ReadWallet(); err != nil {
		beego.Error(err)
		return
	}
	if totalPrice > user.Wallet.Balance {
		errorResponse := ApiError.APIErrorResponse{
			Err:    "not sufficient funds",
			Detail: "not sufficient funds",
			Code:   ApiError.NotSufficientFunds,
		}
		errorResponse.ServerError(c.Controller, 400)
		return
	}
	transaction := models.Transaction{
		Type:    "Order",
		Balance: user.Wallet.Balance,
		Amount:  -totalPrice,
		Order:   &order,
		User:    user,
	}
	err = transaction.Save()
	if err != nil {
		beego.Error(err)
		return
	}
	order.State = "Done"
	err = order.Update("State")
	if err != nil {
		beego.Error(err)
		return
	}
	user.Wallet.Balance += transaction.Amount
	err = user.Wallet.Update("Balance")
	if err != nil {
		beego.Error(err)
		return
	}
	c.Data["json"] = &serializer.CommonApiResponseBody{
		Success: true,
	}
	c.ServeJSON()
}
