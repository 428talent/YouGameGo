package service

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
	"yougame.com/yougame-server/models"
)

type GameQueryBuilder struct {
	ResourceQueryBuilder
	searchName        string
	goods             []interface{}
	gameCollectionIds []interface{}
	tags              []interface{}
	priceStart        *float64
	priceEnd          *float64
	releaseTimeStart  string
	releaseTimeEnd    string
}

func (b *GameQueryBuilder) Delete() error {
	condition := orm.NewCondition()
	if len(b.ids) > 0 {
		condition = condition.And("id__in", b.ids...)
	} else {
		return nil
	}
	err := models.DeleteGameMultiple(func(o orm.QuerySeter) orm.QuerySeter {
		return o.SetCond(condition)
	})
	return err
}

func (b *GameQueryBuilder) InGameCollection(id ...interface{}) {
	b.gameCollectionIds = append(b.gameCollectionIds, id)
}
func (b *GameQueryBuilder) InGood(id ...interface{}) {
	b.goods = append(b.goods, id)
}
func (b *GameQueryBuilder) InTag(id ...interface{}) {
	b.tags = append(b.tags, id)
}
func (b *GameQueryBuilder) ApiQuery() (*int64, interface{}, error) {
	return b.Query()
}

func (b *GameQueryBuilder) SearchWithName(key string) {
	b.searchName = key
}
func (b *GameQueryBuilder) InPriceStart(value float64) {
	b.priceStart = &value
}
func (b *GameQueryBuilder) InPriceEnd(value float64) {
	b.priceEnd = &value
}

func (b *GameQueryBuilder) InReleaseTimeStart(value string) {
	b.releaseTimeStart = value
}
func (b *GameQueryBuilder) InReleaseTimeEnd(value string) {
	b.releaseTimeEnd = value
}
func (b *GameQueryBuilder) Query() (*int64, []*models.Game, error) {
	condition := b.build()

	if len(b.searchName) > 0 {
		condition = condition.And("name__icontains", b.searchName)
	}
	if len(b.goods) > 0 {
		condition = condition.And("Goods__id__in", b.goods)
	}
	if len(b.tags) > 0 {
		condition = condition.And("Tags__tag_id__in", b.tags)
	}
	if len(b.gameCollectionIds) > 0 {
		condition = condition.And("Collections__game_collection_id__in", b.gameCollectionIds...)
	}

	if b.priceStart != nil {
		condition = condition.And("price__gte", *b.priceStart)
	}
	if b.priceEnd != nil {
		condition = condition.And("price__lte", *b.priceEnd)
	}


	if len(b.releaseTimeStart) > 0 {
		condition = condition.And("release_time__gte", b.releaseTimeStart)
	}
	if len(b.releaseTimeEnd) > 0 {
		condition = condition.And("release_time__lte", b.releaseTimeEnd)
	}

	return models.GetGameList(func(o orm.QuerySeter) orm.QuerySeter {
		querySetter := o.SetCond(condition).Limit(b.pageOption.PageSize).Offset(b.pageOption.Offset())
		if len(b.orders) > 0 {
			querySetter = querySetter.OrderBy(b.orders...)
		}
		return querySetter
	})
}

func GetGameBand(gameId int, imageType string) (*models.Image, error) {
	game := models.Game{Id: gameId}
	err := game.QueryById()
	if err != nil {
		return nil, err
	}
	imageQueryBuilder := ImageQueryBuilder{}
	if imageType == "android" {
		imageQueryBuilder.WithName(fmt.Sprint("Band:", game.Id, ":android"))
	} else {
		imageQueryBuilder.WithName(fmt.Sprint("Band:", game.Id))
	}

	count, imageList, err := imageQueryBuilder.Query()
	if err != nil {
		panic(err)
	}
	if *count == 0 {
		return nil, NotFound
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
	o := orm.NewOrm()
	err = game.Save(o)
	if err != nil {
		return nil, err
	}
	return game, nil
}

func AddGameTags(gameId int, tagIds ...int) error {
	o := orm.NewOrm()

	m2m := o.QueryM2M(&models.Game{Id: gameId}, "Tags")
	for _, tagId := range tagIds {
		_, err := m2m.Add(&models.Tag{Id: tagId})
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteGameTags(gameId int, tagIds ...int) error {
	o := orm.NewOrm()

	m2m := o.QueryM2M(&models.Game{Id: gameId}, "Tags")
	for _, tagId := range tagIds {
		_, err := m2m.Remove(&models.Tag{Id: tagId})
		if err != nil {
			return err
		}
	}
	return nil
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

func GetGameWithUserInventory(userId int, pageOption PageOption) (int64, []*models.Game, error) {
	return models.GetGameWithInventory(userId, int(pageOption.PageSize), int(pageOption.Offset()))
}
