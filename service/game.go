package service

import (
	"github.com/astaxie/beego/orm"
	"time"
	"yougame.com/yougame-server/models"
)

type GameQueryBuilder struct {
	ids        []interface{}
	pageOption *PageOption
}

func (b *GameQueryBuilder) InId(ids ...int) {
	b.ids = append(b.ids, ids)
}

func (b *GameQueryBuilder) Query() (*int64, []*models.Game, error) {
	condition := orm.NewCondition()
	if len(b.ids) > 0 {
		condition = condition.And("id", b.ids...)
	}
	if b.pageOption == nil {
		b.pageOption = &PageOption{
			Page:     1,
			PageSize: 10,
		}
	}
	return models.GetGameList(func(o orm.QuerySeter) orm.QuerySeter {
		return o.SetCond(condition).Limit(b.pageOption.Page).Offset(b.pageOption.Offset())
	})
}

func CreateNewGame(name string, price float32, intro string, publisher string, releaseTime time.Time) (game *models.Game, err error) {
	game = &models.Game{
		Name:        name,
		Price:       price,
		Intro:       intro,
		Publisher:   publisher,
		ReleaseTime: releaseTime,
	}
	err = game.Save()
	if err != nil {
		return nil, err
	}
	return game, nil
}
