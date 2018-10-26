package models

import "time"

type Wallet struct {
	Id      int
	UserId  int
	Balance float64
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}
