package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Good struct {
	Id             int
	Name           string           `mapstructure:"name"`
	Price          float64          `mapstructure:"name"`
	Enable         bool             `mapstructure:"enable"`
	Users          []*User          `orm:"reverse(many)"`
	InventoryItems []*InventoryItem `orm:"reverse(many)"`
	Comments       []*Comment       `orm:"reverse(many)"`
	Game           *Game            `orm:"rel(fk)"`
	Created        time.Time        `orm:"auto_now_add;type(datetime)"`
	Updated        time.Time        `orm:"auto_now;type(datetime)"`
}


func (g *Good) DeleteMultiple(o orm.Ormer, ids []interface{}) error {
	condition := orm.NewCondition()
	condition = condition.And("id__in", ids...)
	_, err := o.QueryTable("good").SetCond(condition).Update(orm.Params{
		"enable": false,
	})
	return err
}

func (g *Good) Query(id int64) error {
	o := orm.NewOrm()
	g.Id = int(id)
	err := o.Read(g)
	return err
}

func (g *Good) Save(o orm.Ormer) error {
	_, err := o.Insert(g)
	return err
}

func (g *Good) Delete(o orm.Ormer) error {
	g.Enable = false
	_, err := o.Update(g, "enable")
	return err
}

func (g *Good) Update(id int64, o orm.Ormer, fields ...string) error {
	g.Id = int(id)
	_, err := o.Update(g, fields...)
	return err
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

func GetGoodList(filter func(o orm.QuerySeter) orm.QuerySeter) (*int64, []*Good, error) {
	o := orm.NewOrm()
	var goodList []*Good
	setter := filter(o.QueryTable("good"))
	_, err := setter.All(&goodList)
	if err != nil {
		return nil, nil, err
	}
	count, err := setter.Count()
	return &count, goodList, nil
}

func DeleteGoodMultiple(filter func(o orm.QuerySeter) orm.QuerySeter) error {
	o := orm.NewOrm()
	setter := filter(o.QueryTable("good"))
	_, err := setter.Update(orm.Params{
		"enable": false,
	})
	return err
}

//func (g *Good) Update(o orm.Ormer, fields ...string) error {
//	_, err := o.Update(g, fields...)
//	return err
//}
