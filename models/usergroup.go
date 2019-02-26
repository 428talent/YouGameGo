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
func GetUserGroupList(filter func(o orm.QuerySeter) orm.QuerySeter) (int64, []*UserGroup, error) {
	o := orm.NewOrm()
	var userGroupList []*UserGroup
	seter := o.QueryTable("user_group")
	_, err := filter(seter).All(&userGroupList)
	count, err := filter(seter).Count()
	return count, userGroupList, err
}