package models

import "time"

type UserGroup struct {
	Id          int
	Name        string
	Permissions []*Permission `orm:"rel(m2m)"`
	Users       []*User       `orm:"reverse(many)"`
	Created     time.Time     `orm:"auto_now_add;type(datetime)"`
	Enable      bool
}
