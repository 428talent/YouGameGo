package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Image struct {
	Id     int
	Name   string
	Path   string
	Type   string
	Enable bool

	Preview []*Game   `orm:"reverse(many)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

func GetImageList(filter func(o orm.QuerySeter) orm.QuerySeter) (*int64, []*Image, error) {
	o := orm.NewOrm()
	var imageList []*Image
	seter := o.QueryTable("image")
	_, err := filter(seter).All(&imageList)
	if err != nil {
		return nil, nil, err
	}
	count, err := filter(seter).Count()
	return &count, imageList, err
}
