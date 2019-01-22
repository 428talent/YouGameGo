package service

import (
	"github.com/astaxie/beego/orm"
	"yougame.com/yougame-server/models"
)

type TagQueryBuilder struct {
	ResourceQueryBuilder
	game      []interface{}
	searchKey interface{}
	NameOption
}

func (b *TagQueryBuilder) ApiQuery() (*int64, interface{}, error) {
	return b.Query()
}

func (b *TagQueryBuilder) WithGame(game ...interface{}) {
	b.game = append(b.game, game...)
}
func (b *TagQueryBuilder) Search(key interface{}) {
	b.searchKey = key
}
func (b *TagQueryBuilder) Query() (*int64, []*models.Tag, error) {
	condition := b.build()
	if len(b.ids) > 0 {
		condition = condition.And("id", b.ids...)
	}
	if b.pageOption == nil {
		b.pageOption = &PageOption{
			Page:     1,
			PageSize: 10,
		}
	}
	if len(b.names) > 0 {
		condition = condition.And("name__in", b.names...)
	}
	if len(b.game) > 0 {
		condition = condition.And("Games__Game__Id__in", b.game...)
	}
	if b.searchKey != nil {
		condition = condition.And("name__icontains", b.searchKey)
	}

	return models.GetTagList(func(o orm.QuerySeter) orm.QuerySeter {
		setter := o.SetCond(condition).Limit(b.pageOption.PageSize).Offset(b.pageOption.Offset())
		if len(b.orders) > 0 {
			setter = setter.OrderBy(b.orders...)
		}
		return setter

	})
}
