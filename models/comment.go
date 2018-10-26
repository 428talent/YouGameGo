package models

import "time"

type Comment struct {
	Id         int
	Game       *Game `orm:"rel(fk)"`
	Evaluation string
	Created    time.Time `orm:"auto_now_add;type(datetime)"`
	Updated    time.Time `orm:"auto_now;type(datetime)"`
}
