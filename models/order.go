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
func (order *Order) QueryById() error {
	o := orm.NewOrm()
	return o.Read(order)
}

func GetOrderList(filter func(o orm.QuerySeter) orm.QuerySeter) ([]*Order, error) {
	o := orm.NewOrm()
	var orderList []*Order
	seter := o.QueryTable("order")
	_, err := filter(seter).All(&orderList)
	return orderList, err
}
func (orderGood *OrderGood) QueryById() error {
	o := orm.NewOrm()
	return o.Read(orderGood)
}

func (order *Order) ReadOrderGoods() error {
	o := orm.NewOrm()
	_, err := o.LoadRelated(order, "Goods")
	return err
}

func (orderGood *OrderGood) ReadGood() error {
	o := orm.NewOrm()
	err := o.Read(orderGood)
	return err
}
