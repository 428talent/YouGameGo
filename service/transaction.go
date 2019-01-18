package service

import (
	"github.com/astaxie/beego/orm"
	"yougame.com/yougame-server/models"
)

type TransactionQueryBuilder struct {
	ResourceQueryBuilder
	UserIdOption
}

func (builder *TransactionQueryBuilder) Query() (*int64, []*models.Transaction, error) {
	condition := builder.build()
	if len(builder.userIds) > 0 {
		condition = condition.And("User__id__in", builder.userIds...)
	}
	return models.GetTransactionList(func(o orm.QuerySeter) orm.QuerySeter {
		querySetter := o.SetCond(condition).Limit(builder.pageOption.PageSize).Offset(builder.pageOption.Offset())
		if len(builder.orders) > 0 {
			querySetter = querySetter.OrderBy(builder.orders...)
		}
		return querySetter
	})
}
func (builder *TransactionQueryBuilder) ApiQuery() (*int64, interface{}, error) {
	return builder.Query()
}
