package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
	"yougame.com/yougame-server/util"
)

//用户模块
type User struct {
	Id           int
	Username     string `orm:"unique"`
	Password     string
	LastLogin    time.Time `orm:"null"`
	Enable       bool
	UserGroups   []*UserGroup   `orm:"rel(m2m)"`
	ShoppingCart []*CartItem    `orm:"reverse(many)"`
	Orders       []*Order       `orm:"reverse(many)"`
	Comments     []*Comment     `orm:"reverse(many)"`
	Good         []*Good        `orm:"rel(m2m)"`
	Transactions []*Transaction `orm:"reverse(many)"`
	Created      time.Time      `orm:"auto_now_add;type(datetime)"`
	Updated      time.Time      `orm:"auto_now;type(datetime)"`
	Profile      *Profile       `orm:"null;rel(one);on_delete(set_null)"`
	Wallet       *Wallet        `orm:"null;rel(one);on_delete(set_null)"`
}

func (u *User) TableName() string {
	return "auth_user"
}

func GetUserById(userId int) (*User, error) {
	o := orm.NewOrm()
	user := &User{Id: userId}
	if err := o.Read(user); err != nil {
		return nil, err
	}
	return user, nil
}
func GetAllUser(page int64, pageSize int64) (*int64, []*User, error) {
	o := orm.NewOrm()
	var userList []*User
	count, err := o.QueryTable("auth_user").Count()
	if err != nil {
		return nil, nil, err
	}
	_, err = o.QueryTable("auth_user").Limit(pageSize).Offset((page - 1) * pageSize).All(&userList)
	if err != nil {
		return nil, nil, err
	}
	return &count, userList, nil
}
func CheckUserValidate(loginUser *User) bool {
	encryptPassword, err := util.EncryptSha1WithSalt(loginUser.Password)
	if err != nil {
		beego.Error(err)
		return false
	}
	loginUser.Password = *encryptPassword
	o := orm.NewOrm()
	beego.Debug(*encryptPassword)

	beego.Debug(loginUser.Username)
	if err = o.Read(loginUser, "username", "password"); err == orm.ErrNoRows {
		beego.Error(err)
		return false
	}
	return true
}

func (u *User) ReadProfile() error {
	o := orm.NewOrm()
	err := o.Read(u.Profile)
	return err
}
func (u *User) ReadWallet() error {
	o := orm.NewOrm()
	err := o.Read(u.Wallet)
	return err
}
func (u *User) ReadCart(offset int64, limit int64, order string) error {
	o := orm.NewOrm()
	_, err := o.LoadRelated(u, "ShoppingCart", 3, limit, offset, order)
	return err
}

func (u *User) ReadOrders(offset int64, limit int64, order string) error {
	o := orm.NewOrm()
	_, err := o.LoadRelated(u, "Orders", 3, limit, offset, order)
	return err
}

func (u *User) ReadUserGroup() error {
	o := orm.NewOrm()
	_, err := o.LoadRelated(u, "UserGroups")
	return err
}

func GetUserList(filter func(o orm.QuerySeter) orm.QuerySeter) (int64, []*User, error) {
	o := orm.NewOrm()
	var userList []*User
	seter := o.QueryTable("auth_user")
	_, err := filter(seter).All(&userList)
	count, err := filter(seter).Count()
	return count, userList, err
}
