package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	"time"
	"you_game_go/util"
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

func CheckUserValidate(loginUser *User) bool {
	appConfig, err := config.NewConfig("ini", "./conf/app_local.conf")
	if err != nil {
		beego.Error(err)
	}
	loginUser.Password = util.EncryptSha1(loginUser.Password + appConfig.String("salt"))
	o := orm.NewOrm()
	if err = o.Read(loginUser, "username", "password"); err == orm.ErrNoRows {
		beego.Error(err)
		return false
	}
	return true
}
