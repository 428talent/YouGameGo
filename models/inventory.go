package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type InventoryItem struct {
	Id      int
	Good    *Good `orm:"rel(fk)"`
	User    *User `orm:"rel(fk)"`
	Enable  bool
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

func GetInventoryItemList(filter func(o orm.QuerySeter) orm.QuerySeter) (*int64, []*InventoryItem, error) {
	o := orm.NewOrm()
	var inventoryList []*InventoryItem
	seter := o.QueryTable("inventory_item")
	_, err := filter(seter).All(&inventoryList)
	if err != nil {
		return nil, nil, err
	}
	count, err := filter(seter).Count()
	return &count, inventoryList, err
}