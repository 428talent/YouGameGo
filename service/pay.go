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
	if err != nil {
		return err
	}

	dbTransaction := func() error {
		if err = order.QueryById(); err != nil {
			return err
		}
		user, err := models.GetUserById(order.User.Id)
		if err != nil {
			return err
		}
		if order.State != "Created" {
			return WrongOrderStateError
		}
		if err = order.ReadOrderGoods(); err != nil {
			return err
		}
		totalPrice := 0.0
		for _, orderGood := range order.Goods {
			totalPrice += orderGood.Price
		}
		if err = user.ReadWallet(); err != nil {
			return err
		}
		if totalPrice > user.Wallet.Balance {
			return NotSufficientFundsError
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
			return err
		}
		order.State = models.OrderStateDone
		err = order.Update(int64(order.Id), o, "State")
		if err != nil {
			return err
		}
		user.Wallet.Balance += transaction.Amount
		err = user.Wallet.Update(o, "Balance")
		if err != nil {
			return err
		}
		return nil
	}

	err = dbTransaction()
	if err != nil {
		return err
	}
	err = o.Commit()
	if err != nil {
		return err
	}

	return nil
}
