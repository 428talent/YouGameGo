package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Wallet struct {
	Id      int
	User    *User `orm:"reverse(one)"`
	Balance float64
	Enable  bool
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func (w *Wallet) Update(o orm.Ormer, fields ...string) error {
	_, err := o.Update(w)
	return err
}


func GetWalletList(filter func(o orm.QuerySeter) orm.QuerySeter) (*int64, []*Wallet, error) {
	o := orm.NewOrm()
	var walletList []*Wallet
	seter := o.QueryTable("wallet")
	_, err := filter(seter).All(&walletList)
	if err != nil {
		return nil, nil, err
	}
	count, err := filter(seter).Count()
	return &count, walletList, err
}