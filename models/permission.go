package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Permission struct {
	Id         int
	Name       string
	UserGroups []*UserGroup `orm:"reverse(many)"`
	Created    time.Time    `orm:"auto_now_add;type(datetime)"`
	Enable     bool
}

func GetPermissionList(filter func(o orm.QuerySeter) orm.QuerySeter) (int64, []*Permission, error) {
	o := orm.NewOrm()
	var permissionList []*Permission
	seter := o.QueryTable("permission")
	_, err := filter(seter).All(&permissionList)
	count, err := filter(seter).Count()
	return count, permissionList, err
}