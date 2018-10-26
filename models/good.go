package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Good struct {
	Id      int
	Name    string
	Price   float64
	Game    *Game     `orm:"rel(fk)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func (g *Good) QueryById() error {
	o := orm.NewOrm()
	err := o.Read(g)
	return err
}

func (g *Good) ReadGame() error {
	o := orm.NewOrm()
	_, err := o.LoadRelated(g, "Game")
	return err
}
