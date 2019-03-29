package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Follow struct {
	Id        int
	User      *User     `orm:"rel(fk)"`
	Following *User     `orm:"rel(fk)"`
	Enable    bool      `mapstructure:"enable"`
	Created   time.Time `orm:"auto_now_add;type(datetime)"`
	Updated   time.Time `orm:"auto_now;type(datetime)"`
}

func (f *Follow) Query(id int64) error {
	f.Id = int(id)
	o := orm.NewOrm()
	err := o.Read(f)
	return err
}

func (f *Follow) Save(o orm.Ormer) error {
	_, err := o.Insert(f)
	return err
}

func (f *Follow) Delete(o orm.Ormer) error {
	_, err := o.Delete(f, "id")
	return err
}

func (f *Follow) Update(id int64, o orm.Ormer, fields ...string) error {
	f.Id = int(id)
	_, err := o.Update(f, fields...)
	return err
}

func GetFollowList(filter func(o orm.QuerySeter) orm.QuerySeter) (*int64, []*Follow, error) {
	o := orm.NewOrm()
	var followList []*Follow
	seter := o.QueryTable("follow")
	_, err := filter(seter).All(&followList)
	if err != nil {
		return nil, nil, err
	}
	count, err := filter(seter).Count()
	return &count, followList, err
}
