package models

import "time"

type Requirement struct {
	Id int
	Name string
	Value string
	Game    *Game     `orm:"rel(fk)"`
	Created time.Time    `orm:"auto_now_add;type(datetime)"`
	Updated time.Time    `orm:"auto_now;type(datetime)"`
}
