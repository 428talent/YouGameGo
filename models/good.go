package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Good struct {
	Id       int
	Name     string
	Price    float64
	Enable   bool
	Users    []*User    `orm:"reverse(many)"`
	Comments []*Comment `orm:"reverse(many)"`
	Game     *Game      `orm:"rel(fk)"`
	Created  time.Time  `orm:"auto_now_add;type(datetime)"`
	Updated  time.Time  `orm:"auto_now;type(datetime)"`
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

func GetGoodList(filter func(o orm.QuerySeter) orm.QuerySeter) (*int64,[]*Good, error) {
	o := orm.NewOrm()
	var goodList []*Good
	setter := filter(o.QueryTable("good"))
	_, err := setter.All(&goodList)
	if err != nil {
		return nil,nil, err
	}
	count, err := setter.Count()
	return &count,goodList, nil
}

func (g *Good) Update(o orm.Ormer, fields ...string) error {
	_, err := o.Update(g, fields...)
	return err
}
