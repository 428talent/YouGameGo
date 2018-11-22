package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

const (
	OrderStateCreated = "Created"
	OrderStateDone    = "Done"
)

type OrderState string

type Order struct {
	Id          int
	Enable      bool
	State       OrderState
	User        *User        `orm:"rel(fk)"`
	Transaction *Transaction `orm:"reverse(one)"`
	Goods       []*OrderGood `orm:"reverse(many)"`
	Created     time.Time    `orm:"auto_now_add;type(datetime)"`
	Updated     time.Time    `orm:"auto_now;type(datetime)"`
}

type OrderGood struct {
	Id      int
	Price   float64
	Name    string
	Order   *Order    `orm:"rel(fk)"`
	Good    *Good     `orm:"rel(fk)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

func (orderGood *OrderGood) GetList(filter func(o orm.QuerySeter) orm.QuerySeter, md interface{}) (count int64, err error) {
	o := orm.NewOrm()
	seter := o.QueryTable("order_good")
	_, err = filter(seter).All(md)
	if err != nil {
		return
	}
	count, err = filter(seter).Count()
	return
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

func GetOrderList(filter func(o orm.QuerySeter) orm.QuerySeter) (int64, []*Order, error) {
	o := orm.NewOrm()
	var orderList []*Order
	seter := o.QueryTable("order")
	_, err := filter(seter).All(&orderList)
	count, err := filter(seter).Count()
	return count, orderList, err
}
func GetOrderGoodList(filter func(o orm.QuerySeter) orm.QuerySeter) (int64, []*OrderGood, error) {
	o := orm.NewOrm()
	var orderGoodList []*OrderGood
	seter := o.QueryTable("order_good")
	_, err := filter(seter).All(&orderGoodList)
	count, err := filter(seter).Count()
	return count, orderGoodList, err
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
func (d *Order) Update(o orm.Ormer, fields ...string) error {
	_, err := o.Update(d, fields...)
	return err
}
