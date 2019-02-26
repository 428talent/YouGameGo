package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type UserGroup struct {
	Id          int
	Name        string
	Permissions []*Permission `orm:"rel(m2m)"`
	Users       []*User       `orm:"reverse(many)"`
	Created     time.Time     `orm:"auto_now_add;type(datetime)"`
	Enable      bool
}

func (g *UserGroup) Query(id int64) error {
	g.Id = int(id)
	err := orm.NewOrm().Read(g)
	return err
}

func (g *UserGroup) Save(o orm.Ormer) error {
	_, err := o.Insert(g)
	return err
}

func (g *UserGroup) Delete(o orm.Ormer) error {
	g.Enable = false
	_, err := o.Update(g, "enable")
	return err
}

func (g *UserGroup) Update(id int64, o orm.Ormer, fields ...string) error {
	g.Id = int(id)
	_, err := o.Update(g, fields...)
	return err
}

func GetUserGroupList(filter func(o orm.QuerySeter) orm.QuerySeter) (int64, []*UserGroup, error) {
	o := orm.NewOrm()
	var userGroupList []*UserGroup
	seter := o.QueryTable("user_group")
	_, err := filter(seter).All(&userGroupList)
	count, err := filter(seter).Count()
	return count, userGroupList, err
}
