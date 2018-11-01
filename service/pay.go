package service

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"yougame.com/yougame-server/models"
)

var (
	NotSufficientFundsError = errors.New("NotSufficientFunds")
	WrongOrderStateError    = errors.New("WrongOrderStateError")
)

func PayOrder(order models.Order) error {
	o := orm.NewOrm()
	err := o.Begin()
	if err = order.QueryById(); err != nil {
		panic(err)
	}
	user, err := models.GetUserById(order.User.Id)
	if err != nil {
		panic(err)
	}
	if order.State != "Created" {
		panic(WrongOrderStateError)
	}
	if err = order.ReadOrderGoods(); err != nil {
		panic(err)
	}
	totalPrice := 0.0
	for _, orderGood := range order.Goods {
		totalPrice += orderGood.Price
	}
	if err = user.ReadWallet(); err != nil {
		panic(err)
	}
	if totalPrice > user.Wallet.Balance {
		panic(NotSufficientFundsError)
	}
	transaction := models.Transaction{
		Type:    "Order",
		Balance: user.Wallet.Balance,
		Amount:  -totalPrice,
		Order:   &order,
		User:    user,
	}
	err = transaction.Save(o)
	if err != nil {
		panic(err)
	}
	order.State = models.OrderStateDone
	err = order.Update(o,"State")
	if err != nil {
		panic(err)
	}
	user.Wallet.Balance += transaction.Amount
	err = user.Wallet.Update(o,"Balance")
	if err != nil {
		panic(err)
	}
	defer func() {
		reco := recover()
		if reco != nil{
			err = reco.(error)
			err = o.Rollback()
		}else{
			err = o.Commit()
		}
	}()

	return err
}
