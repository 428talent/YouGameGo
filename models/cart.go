package models

import "time"

type CartItem struct {
	Id      int
	Good    *Good `orm:"rel(fk)"`
	UserId  int
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}
