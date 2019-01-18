package service

import (
	"github.com/astaxie/beego/orm"
	"yougame.com/yougame-server/models"
)

type WalletQueryBuilder struct {
	ResourceQueryBuilder
	userIds []interface{}
}

func (builder *WalletQueryBuilder) InUser(userId ...interface{}) {
	builder.userIds = append(builder.userIds, userId...)
}
func (builder *WalletQueryBuilder) Query() (*int64, []*models.Wallet, error) {
	condition := builder.build()
	if len(builder.userIds) > 0 {
		condition = condition.And("User__id__in", builder.userIds...)
	}
	return models.GetWalletList(func(o orm.QuerySeter) orm.QuerySeter {
		querySetter := o.SetCond(condition).Limit(builder.pageOption.PageSize).Offset(builder.pageOption.Offset())
		if len(builder.orders) > 0 {
			querySetter = querySetter.OrderBy(builder.orders...)
		}
		return querySetter
	})
}
func (builder *WalletQueryBuilder) ApiQuery() (*int64, interface{}, error) {
	return builder.Query()
}
