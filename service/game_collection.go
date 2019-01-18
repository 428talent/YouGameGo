package service

import (
	"github.com/astaxie/beego/orm"
	"yougame.com/yougame-server/models"
)

type GameCollectionQueryBuilder struct {
	ResourceQueryBuilder
}

func (builder *GameCollectionQueryBuilder) ApiQuery() (*int64, interface{}, error) {
	return builder.Query()
}
func (builder *GameCollectionQueryBuilder) Query() (*int64, []*models.GameCollection, error) {
	condition := builder.build()
	return models.GetGameCollectionList(func(o orm.QuerySeter) orm.QuerySeter {
		querySetter := o.SetCond(condition).Limit(builder.pageOption.PageSize).Offset(builder.pageOption.Offset())
		if len(builder.orders) > 0 {
			querySetter = querySetter.OrderBy(builder.orders...)
		}
		return querySetter
	})
}

func AddGameToCollection(collectionId int, games ...int) error {
	gameCollection := models.GameCollection{}
	err := gameCollection.Query(int64(collectionId))
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	for _, gameId := range games {
		err = gameCollection.AddGame(o, gameId)
		if err != nil {
			return err
		}
	}
	return nil
}
