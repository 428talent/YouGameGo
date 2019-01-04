package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"os"
	"time"
)

type Game struct {
	Id            int       `json:"id"`
	Name          string    `json:"name"`
	Price         float32   `json:"price"`
	ReleaseTime   time.Time `json:"release_time"`
	Publisher     string    `json:"publisher"`
	Enable        bool      `json:"enable"`
	Band          *Image    `orm:"null;rel(one);on_delete(set_null)"`
	Intro         string
	Tags          []*Tag    `orm:"rel(m2m)"`
	PreviewImages []*Image  `orm:"rel(m2m)"`
	Goods         []*Good   `orm:"reverse(many)"`
	Created       time.Time `orm:"auto_now_add;type(datetime)"`
	Updated       time.Time `orm:"auto_now;type(datetime)"`
}

func (g *Game) Update(id int64, o orm.Ormer, fields ...string) error {
	g.Id = int(id)
	_, err := o.Update(g, fields...)
	return err
}

func (g *Game) Delete(o orm.Ormer) error {
	g.Enable = false
	_, err := o.Update(g, "enable")
	return err
}

//保存游戏信息
func (g *Game) Save() error {
	o := orm.NewOrm()
	_, err := o.Insert(g)
	return err
}

func (g *Game) Query(id int64) error {
	g.Id = int(id)
	o := orm.NewOrm()
	err := o.Read(g)
	return err

}

func (g *Game) QueryById() error {
	o := orm.NewOrm()
	err := o.Read(g)
	return err

}

func (g *Game) SaveGameBangImage(path string) (*Image, error) {
	o := orm.NewOrm()

	if g.Band == nil {
		image := Image{
			Path: path,
			Type: "Band",
			Name: fmt.Sprintf("band:%d", g.Id),
		}
		_, err := o.Insert(&image)
		if err != nil {
			return nil, err
		}
		g.Band = &image
		_, err = o.Update(g, "Band")
		if err != nil {
			return nil, err
		}
		return nil, err
	}
	err := o.Read(g.Band)
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat(g.Band.Path); os.IsExist(err) {
		err = os.Remove(g.Band.Path)
		if err != nil {
			return nil, err
		}
	}

	g.Band.Path = path
	_, err = o.Update(g.Band)
	return g.Band, err
}

func GetGameList(filter func(o orm.QuerySeter) orm.QuerySeter) (*int64, []*Game, error) {
	o := orm.NewOrm()
	var gameList []*Game
	seter := o.QueryTable("game")
	_, err := filter(seter).All(&gameList)
	if err != nil {
		return nil, nil, err
	}
	count, err := filter(seter).Count()
	return &count, gameList, err
}

func (g *Game) ReadGameBand() (err error) {
	o := orm.NewOrm()
	err = o.Read(g.Band)
	return err
}

func (g *Game) SavePreviewImage(path string) error {
	o := orm.NewOrm()
	image := Image{
		Type: "Preview",
		Path: path,
		Name: fmt.Sprintf("preview:%d", g.Id),
	}
	imageId, err := o.Insert(&image)
	if err != nil {
		return err
	}
	image.Id = int(imageId)
	m2m := o.QueryM2M(g, "PreviewImages")
	_, err = m2m.Add(image)
	return err
}

func (g *Game) ReadGamePreviewImage() error {
	o := orm.NewOrm()
	_, err := o.LoadRelated(g, "PreviewImages")
	return err
}

func (g *Game) SaveTags(names []string) error {
	o := orm.NewOrm()
	var tags []*Tag
	for _, tagName := range names {
		tag := Tag{
			Name: tagName,
		}
		tagId, err := o.Insert(&tag)
		if err != nil {
			return err
		}
		tag.Id = int(tagId)
		tags = append(tags, &tag)
	}
	m2m := o.QueryM2M(g, "Tags")
	_, err := m2m.Add(tags)
	return err
}

func (g *Game) ReadTags() error {
	o := orm.NewOrm()
	_, err := o.LoadRelated(g, "Tags")
	return err
}

func (g *Game) AddGood(good *Good) error {
	o := orm.NewOrm()
	id, err := o.Insert(good)
	good.Id = int(id)
	return err
}

func (g *Game) ReadGoods() error {
	o := orm.NewOrm()
	_, err := o.LoadRelated(g, "Goods")
	return err
}

func SearchGame(key string) ([]*Game, error) {
	o := orm.NewOrm()
	var gameList []*Game
	_, err := o.QueryTable("game").Filter("name__icontains", key).All(&gameList)
	return gameList, err
}

func (g *Game) UpdateGame(o orm.Ormer, fields ...string) error {
	_, err := o.Update(g, fields...)
	return err
}
