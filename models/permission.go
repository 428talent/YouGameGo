package models

import "time"

type Permission struct {
	Id         int
	Name       string
	UserGroups []*UserGroup `orm:"reverse(many)"`
	Created    time.Time    `orm:"auto_now_add;type(datetime)"`
	Enable     bool
}
