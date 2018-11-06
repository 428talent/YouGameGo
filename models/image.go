package models

import "time"

type Image struct {
	Id      int
	Name    string
	Path    string
	Type    string
	Enable bool

	Preview []*Game `orm:"reverse(many)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}
