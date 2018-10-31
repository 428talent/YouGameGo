package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Transaction struct {
	Id      int
	Type    string
	Balance float64
	Amount  float64
	Order   *Order    `orm:"null;rel(one);on_delete(set_null)"`
	User    *User     `orm:"rel(fk)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

func (t *Transaction) Save(o orm.Ormer) error {
	transaction, err := o.Insert(t)
	t.Id = int(transaction)
	return err
}
