package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Tag struct {
	Id     int
	Name   string
	Enable bool

	Games    []*Game   `orm:"reverse(many)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func GetTagList(filter func(o orm.QuerySeter) orm.QuerySeter) (*int64,[]*Tag, error) {
	o := orm.NewOrm()
	var tagList []*Tag
	seter := o.QueryTable("tag")
	_, err := filter(seter).All(&tagList)
	count, err := filter(seter).Count()
	return &count, tagList, err
}
