package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type CartItem struct {
	Id      int
	Good    *Good `orm:"rel(fk)"`
	Enable bool
	User *User `orm:"rel(fk)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

func (c *CartItem) QueryCartById() error {
	o := orm.NewOrm()
	err := o.Read(c)
	return err
}

func (c *CartItem) ReadGood() error {
	o := orm.NewOrm()
	_, err := o.LoadRelated(c, "Good")
	return err
}

func GetCartList(filter func(o orm.QuerySeter) orm.QuerySeter) (int64,[]*CartItem, error) {
	o := orm.NewOrm()
	var cartList []*CartItem
	seter := o.QueryTable("cart_item")
	_, err := filter(seter).All(&cartList)
	count, err := filter(seter).Count()
	return count,cartList, err
}

func (c *CartItem) Save() error {
	o := orm.NewOrm()
	_, err := o.Insert(c)
	return err
}

func (c *CartItem) Delete() error {
	o := orm.NewOrm()
	_, err := o.Delete(c, "id")
	return err
}
func (c *CartItem) DeleteAll() error {
	o := orm.NewOrm()
	_, err := o.Delete(c, "user_id")
	return err
}
