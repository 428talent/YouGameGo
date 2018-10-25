package models

import "time"

type Good struct {
	Id int
	Name string
	Price float64
	Game *Game `orm:"rel(fk)"`
	Created       time.Time `orm:"auto_now_add;type(datetime)"`
	Updated       time.Time `orm:"auto_now;type(datetime)"`
}
