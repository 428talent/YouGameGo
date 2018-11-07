package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Comment struct {
	Id         int
	Good       *Good `orm:"rel(fk)"`
	User       *User `orm:"rel(fk)"`
	Content string
	Evaluation string
	Enable     bool
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func (comment *Comment) Save(o orm.Ormer) error {
	_, err := o.Insert(comment)
	return err
}
