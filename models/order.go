package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Order struct {
	Id      int
	State   string
	UserId  int
	Goods   []*OrderGood `orm:"reverse(many)"`
	Created time.Time    `orm:"auto_now_add;type(datetime)"`
	Updated time.Time    `orm:"auto_now;type(datetime)"`
}

type OrderGood struct {
	Id      int
	Price   float64
	Name    string
	Order   *Order    `orm:"rel(fk)"`
	Good    *Good     `orm:"rel(fk)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

func (order *Order) SaveOrder() error {
	o := orm.NewOrm()
	orderId, err := o.Insert(order)
	if err != nil {
		return err
	}
	order.Id = int(orderId)
	for _, orderGoods := range order.Goods {
		//get good current price
		err = orderGoods.Good.QueryById()
		if err != nil {
			return err
		}
		orderGoods.Price = orderGoods.Good.Price
		orderGoods.Name = orderGoods.Good.Name
		//set order
		orderGoods.Order = order
		//save
		_, err := o.Insert(orderGoods)
		if err != nil {
			return err
		}
	}
	return nil
}
