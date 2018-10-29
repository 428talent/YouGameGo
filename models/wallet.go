package models

import "time"

type Wallet struct {
	Id      int
	User    *User `orm:"reverse(one)"`
	Balance float64
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}
