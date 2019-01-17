package service

import (
	"github.com/astaxie/beego/orm"
	"yougame.com/yougame-server/models"
)

type InventoryQueryBuilder struct {
	ResourceQueryBuilder
	userIds []interface{}
	goodIds []interface{}
	gameIds []interface{}
}

func (builder *InventoryQueryBuilder) ApiQuery() (*int64, interface{}, error) {
	return builder.Query()
}

func (builder *InventoryQueryBuilder) BelongUser(id ...interface{}) {
	builder.userIds = append(builder.userIds, id...)
}

func (builder *InventoryQueryBuilder) InGood(id ...interface{}) {
	builder.goodIds = append(builder.goodIds, id...)
}

func (builder *InventoryQueryBuilder) InGame(id ...interface{}) {
	builder.gameIds = append(builder.gameIds, id...)
}
func (builder *InventoryQueryBuilder) Query() (*int64, []*models.InventoryItem, error) {
	condition := builder.build()
	if len(builder.userIds) > 0 {
		condition = condition.And("user_id__in", builder.userIds...)
	}
	if len(builder.goodIds) > 0 {
		condition = condition.And("good_id__in", builder.goodIds...)
	}
	if len(builder.gameIds) > 0 {
		condition = condition.And("Good__Game__id__in", builder.gameIds)
	}
	return models.GetInventoryItemList(func(o orm.QuerySeter) orm.QuerySeter {
		querySetter := o.SetCond(condition).Limit(builder.pageOption.PageSize).Offset(builder.pageOption.Offset())
		if len(builder.orders) > 0 {
			querySetter = querySetter.OrderBy(builder.orders...)
		}
		return querySetter
	})
}
