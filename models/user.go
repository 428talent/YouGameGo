package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//用户模块
type User struct {
	Id        int
	Username  string    `orm:"unique"`
	Password  string
	LastLogin time.Time `orm:"null"`
	Enable    bool
	Created   time.Time `orm:"auto_now_add;type(datetime)"`
	Updated   time.Time `orm:"auto_now;type(datetime)"`
}

func (u *User) TableName() string {
	return "auth_user"
}

func CreateUserAccount(username string, password string) error {
	o := orm.NewOrm()
	user := User{
		Username: username,
		Password: password,
		Enable:   true,
	}
	_, err := o.Insert(&user)
	return err
}
