package service

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/sirupsen/logrus"
	"yougame.com/yougame-server/mail"
	"yougame.com/yougame-server/models"
)

var (
	NotSufficientFundsError = errors.New("NotSufficientFunds")
	WrongOrderStateError    = errors.New("WrongOrderStateError")
	AlreadyInInventoryError = errors.New("already in inventory")
)

func PayOrder(order models.Order) error {
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		return err
	}

	dbTransaction := func() (*models.User, *models.Transaction, *models.Order, [] *models.OrderGood, error) {
		if err = order.QueryById(); err != nil {
			return nil, nil, nil, nil, err
		}
		user, err := models.GetUserById(order.User.Id)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		// check order state
		if order.State != "Created" {
			return nil, nil, nil, nil, WrongOrderStateError
		}

		orderGoodQueryBuilder := OrderGoodQueryBuilder{}
		orderGoodQueryBuilder.WishOrderId(order.Id)
		_, orderGoodList, err := orderGoodQueryBuilder.Query()
		if err != nil {
			return nil, nil, nil, nil, err
		}
		// check inventory
		inventQueryBuilder := InventoryQueryBuilder{}
		inventQueryBuilder.BelongUser(user.Id)
		for _, orderGood := range orderGoodList {
			inventQueryBuilder.InGood(orderGood.Good.Id)
		}
		inventoryCount, _, err := inventQueryBuilder.Query()
		if err != nil {
			return nil, nil, nil, nil, err
		}
		if *inventoryCount != 0 {
			return nil, nil, nil, nil, AlreadyInInventoryError
		}

		// check sufficient
		totalPrice := 0.0
		for _, orderGood := range orderGoodList {
			totalPrice += orderGood.Price
		}
		if err = user.ReadWallet(); err != nil {
			return nil, nil, nil, nil, err
		}
		if totalPrice > user.Wallet.Balance {
			return nil, nil, nil, nil, NotSufficientFundsError
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
			return nil, nil, nil, nil, err
		}
		order.State = models.OrderStateDone
		err = order.Update(int64(order.Id), o, "State")
		if err != nil {
			return nil, nil, nil, nil, err
		}
		user.Wallet.Balance += transaction.Amount
		err = user.Wallet.Update(o, "Balance")
		if err != nil {
			return nil, nil, nil, nil, err
		}
		//add inventory
		inventoryItems := make([]*models.InventoryItem, 0)
		for _, orderGood := range orderGoodList {
			inventoryItem := &models.InventoryItem{
				Good:   orderGood.Good,
				User:   user,
				Enable: true,
			}
			inventoryItems = append(inventoryItems, inventoryItem)
		}
		_, err = o.InsertMulti(len(inventoryItems), inventoryItems)
		if err != nil {
			return nil, nil, nil, nil, err
		}

		return user, &transaction, &order, orderGoodList, nil
	}

	user, transaction, orderResult, orderGoods, err := dbTransaction()
	if err != nil {
		return err
	}
	err = o.Commit()
	if err != nil {
		return err
	}
	// send success mail
	err = user.ReadProfile()
	if err != nil {
		logrus.Error(err)
	}
	if len(user.Profile.Email) > 0 {
		goodQueryBuilder := GoodQueryBuilder{}
		for _, orderGood := range orderGoods {
			goodQueryBuilder.InId(orderGood.Id)
		}

		sql := `
		select order_good.name,order_good.price,game.name as game_name from order_good
		inner join good
		inner join game
		where game.id = good.game_id and
      	order_good.good_id = good.id and
      	order_good.order_id = ?`
		type queryResult struct {
			Name     string
			Price    float64
			GameName string
		}
		var resultList []queryResult
		_, err = o.Raw(sql, order.Id).QueryRows(&resultList)
		if err != nil {
			logrus.Error(err)
		}
		totalPrice := 0.0
		mailGoodItems := make([]*mail.ReceiptMailGoods, 0)
		for _, good := range resultList {
			totalPrice += good.Price
			mailGoodItems = append(mailGoodItems, &mail.ReceiptMailGoods{
				GameName: good.GameName,
				GoodName: good.Name,
				Price:    good.Price,
			})
		}
		mailModel := mail.ReceiptMailModel{
			Name:            user.Username,
			Items:           mailGoodItems,
			OrderId:         int64(orderResult.Id),
			TransactionTime: &transaction.Created,
			TotalPrice:      totalPrice,
			TransactionId:   transaction.Id,
		}
		err = mail.SendReceiptMail(mailModel, user.Profile.Email)
		if err != nil {
			logrus.Error(err)
		}
	}

	return nil
}
