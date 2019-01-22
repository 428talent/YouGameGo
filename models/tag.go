package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Tag struct {
	Id      int
	Name    string
	Enable  bool
	Games   []*Game   `orm:"reverse(many)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func (t *Tag) Query(id int64) error {
	t.Id = int(id)
	err := orm.NewOrm().Read(t)
	return err
}

func (t *Tag) Save(o orm.Ormer) error {
	_, err := o.Insert(t)
	return err
}

func (t *Tag) Delete(o orm.Ormer) error {
	t.Enable = false
	_, err := o.Update(t)
	return err
}

func (t *Tag) Update(id int64, o orm.Ormer, fields ...string) error {
	t.Id = int(id)
	_, err := o.Update(t, fields...)
	return err
}

func GetTagList(filter func(o orm.QuerySeter) orm.QuerySeter) (*int64, []*Tag, error) {
	o := orm.NewOrm()
	var tagList []*Tag
	seter := o.QueryTable("tag")
	_, err := filter(seter).All(&tagList)
	count, err := filter(seter).Count()
	return &count, tagList, err
}
