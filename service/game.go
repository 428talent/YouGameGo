package service

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
	"yougame.com/yougame-server/models"
)

type GameQueryBuilder struct {
	ids        []interface{}
	pageOption *PageOption
	orders     []string
}

func (b *GameQueryBuilder) InId(id ...interface{}) {
	b.ids = append(b.ids, id)
}

func (b *GameQueryBuilder) ApiQuery() (*int64, interface{}, error) {
	return b.Query()
}

func (b *GameQueryBuilder) SetPage(page int64, pageSize int64) {
	b.pageOption = &PageOption{
		Page:     page,
		PageSize: pageSize,
	}
}
func (b *GameQueryBuilder) ByOrder(orders ...string) {
	b.orders = append(b.orders, orders...)
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
		querySetter := o.SetCond(condition).Limit(b.pageOption.PageSize).Offset(b.pageOption.Offset())
		if len(b.orders) > 0 {
			querySetter = querySetter.OrderBy(b.orders...)
		}
		return querySetter
	})
}

func GetGameBand(gameId int) (*models.Image, error) {
	game := models.Game{Id: gameId}
	err := game.QueryById()
	if err != nil {
		return nil, err
	}
	imageQueryBuilder := ImageQueryBuilder{}
	imageQueryBuilder.WithName(fmt.Sprint("Band:", game.Id))
	count, imageList, err := imageQueryBuilder.Query()
	if err != nil {
		panic(err)
	}
	if *count == 0 {
		panic(NotFound)
	}
	return imageList[0], nil
}

func GetGamePreview(gameId int, page int64, pageSize int64) (*int64, []*models.Image, error) {
	game := models.Game{Id: gameId}
	err := game.QueryById()
	if err != nil {
		return nil, nil, err
	}
	imageQueryBuilder := ImageQueryBuilder{}
	imageQueryBuilder.SetPage(page, pageSize)
	imageQueryBuilder.WithName(fmt.Sprint("Preview:", game.Id))
	return imageQueryBuilder.Query()
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

func AddGameTags(gameId int, names ...string) ([]*models.Tag, error) {
	o := orm.NewOrm()
	var tags []*models.Tag
	for _, tagName := range names {
		tag := models.Tag{
			Name: tagName,
		}
		tagId, err := o.Insert(&tag)
		if err != nil {
			return nil, err
		}
		tag.Id = int(tagId)
		tags = append(tags, &tag)
	}

	m2m := o.QueryM2M(&models.Game{Id: gameId}, "Tags")
	_, err := m2m.Add(tags)
	return tags, err
}

func UpdateGame(game *models.Game, fields ...string) error {
	o := orm.NewOrm()
	err := game.UpdateGame(o, fields...)
	if err != nil {
		return err
	}
	err = o.Read(game)

	return err
}
