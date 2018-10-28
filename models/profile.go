package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type Profile struct {
	Id       int
	User     *User  `orm:"reverse(one)"`
	Nickname string `orm:"unique"`
	Email    string
	Avatar   string
	Updated  time.Time `orm:"auto_now;type(datetime)"`
}

func ReadProfile(user *User)  {
	o := orm.NewOrm()
	beego.Debug(user.Profile)
	if user.Profile != nil {
		err := o.Read(user.Profile)
		if err != nil {
			beego.Error(err)
		}
	}

}

func (p *Profile) ChangeUserProfile(email string, nickname string) error {
	o := orm.NewOrm()
	fields := make([]string, 0)

	if len(email) != 0 {
		p.Email = email
		fields = append(fields, "email")
	}

	if len(nickname) != 0 {
		p.Nickname = nickname
		fields = append(fields, "nickname")
	}
	_, err := o.Update(p, fields...)
	return err
}

func (p *Profile) SaveAvatar(path string) error {
	o := orm.NewOrm()
	p.Avatar = path
	_, err := o.Update(p, "avatar")
	return err
}
