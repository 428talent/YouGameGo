package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Wallet struct {
	Id      int
	User    *User `orm:"reverse(one)"`
	Balance float64
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func (w *Wallet) Update(o orm.Ormer, fields ...string) error {
	_, err := o.Update(w)
	return err
}
