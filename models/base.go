package models

import "github.com/astaxie/beego/orm"

type Model interface {
	GetList(filter func(o orm.QuerySeter) orm.QuerySeter, md interface{}) (count int64, err error)
}

type DataModel interface {
	Query(id int64) error
	Save(o orm.Ormer) error
	Delete(o orm.Ormer) error
	Update(id int64, o orm.Ormer, fields ...string) error
}
