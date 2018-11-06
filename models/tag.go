package models

import "time"

type Tag struct {
	Id            int
	Name          string
	Enable bool

	Game		[]*Game`orm:"reverse(many)"`
	Created       time.Time `orm:"auto_now_add;type(datetime)"`
	Updated       time.Time `orm:"auto_now;type(datetime)"`
}
