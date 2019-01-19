package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type GameCollection struct {
	Id      int
	Name    string
	Title   string
	Games   []*Game   `orm:"rel(m2m)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
	Enable  bool
}

func (c *GameCollection) Query(id int64) error {
	c.Id = int(id)
	err := orm.NewOrm().Read(c)
	return err
}

func (c *GameCollection) Save(o orm.Ormer) error {
	_, err := o.Insert(c)
	return err
}

func (c *GameCollection) Delete(o orm.Ormer) error {
	c.Enable = false
	_, err := o.Update(c, "enable")
	return err
}

func (c *GameCollection) Update(id int64, o orm.Ormer, fields ...string) error {
	c.Id = int(id)
	_, err := o.Update(c, fields...)
	return err
}

func GetGameCollectionList(filter func(o orm.QuerySeter) orm.QuerySeter) (*int64, []*GameCollection, error) {
	o := orm.NewOrm()
	var gameCollectionList []*GameCollection
	seter := o.QueryTable("game_collection")
	_, err := filter(seter).All(&gameCollectionList)
	if err != nil {
		return nil, nil, err
	}
	count, err := filter(seter).Count()
	return &count, gameCollectionList, err
}

func (c *GameCollection) AddGame(o orm.Ormer, gameId int) error {
	m2mRel := o.QueryM2M(c, "Games")
	_, err := m2mRel.Add(&Game{Id: gameId})
	return err
}

func (c *GameCollection) DeleteGame(o orm.Ormer, gameId int) error {
	m2mRel := o.QueryM2M(c, "Games")
	_, err := m2mRel.Remove(&Game{Id: gameId})
	return err
}
