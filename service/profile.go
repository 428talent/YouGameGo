package service

import (
	"github.com/astaxie/beego/orm"
	"yougame.com/yougame-server/models"
)

type UserProfileQueryBuilder struct {
	ResourceQueryBuilder
	userIds []interface{}
}

func (b *UserProfileQueryBuilder) InUser(userId ...interface{}) {
	b.userIds = append(b.userIds, userId...)
}

func (b *UserProfileQueryBuilder) Query() (*int64, []*models.Profile, error) {
	condition := b.build()
	if len(b.userIds) > 0 {
		condition = condition.And("User__id__in", b.userIds...)
	}
	return models.GetProfileList(func(o orm.QuerySeter) orm.QuerySeter {
		setter := o.SetCond(condition).Limit(b.pageOption.PageSize).Offset(b.pageOption.Offset())
		if len(b.orders) > 0 {
			setter = setter.OrderBy(b.orders...)
		}
		return setter
	})

}

func (b *UserProfileQueryBuilder) ApiQuery() (*int64, interface{}, error) {
	return b.Query()
}
