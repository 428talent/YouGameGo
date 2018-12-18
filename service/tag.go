package service

import (
	"github.com/astaxie/beego/orm"
	"yougame.com/yougame-server/models"
)

type TagQueryBuilder struct {
	pageOption *PageOption
	ids        []interface{}
	game       []interface{}
}

func (b *TagQueryBuilder) InId(id ...interface{}) {
	b.ids = append(b.ids, id)
}
func (b *TagQueryBuilder) SetPage(page int64, pageSize int64) {
	b.pageOption = &PageOption{
		PageSize: pageSize,
		Page:     page,
	}
}

func (b *TagQueryBuilder) WithGame(game ...interface{}) {
	b.game = append(b.game, game...)
}

func (b *TagQueryBuilder) Query() (*int64, []*models.Tag, error) {
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
	if len(b.game) > 0 {
		condition = condition.And("Games__Game__Id__in", b.game...)
	}
	return models.GetTagList(func(o orm.QuerySeter) orm.QuerySeter {
		return o.SetCond(condition).Limit(b.pageOption.PageSize).Offset(b.pageOption.Offset())

	})
}
