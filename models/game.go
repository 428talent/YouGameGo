package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Game struct {
	Id          int
	Name        string
	Price       float32
	ReleaseTime time.Time
	Publisher   string
	Enable      bool
	Intro       string
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
	Updated     time.Time `orm:"auto_now;type(datetime)"`
}

//保存游戏信息
func (g *Game) Save() error {
	o := orm.NewOrm()
	_, err := o.Insert(g)
	return err
}
