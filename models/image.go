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

func (i *Image) Query(id int64) error {
	o := orm.NewOrm()
	i.Id = int(id)
	err := o.Read(i)
	return err
}

func (i *Image) Save(o orm.Ormer) error {
	_, err := o.Insert(i)
	return err
}

func (i *Image) Delete(o orm.Ormer) error {
	i.Enable = false
	_, err := o.Update(i, "enable")
	return err
}

func (i *Image) Update(id int64, o orm.Ormer, fields ...string) error {
	return nil
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
