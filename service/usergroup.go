package service

import (
	"github.com/astaxie/beego/orm"
	"yougame.com/yougame-server/models"
)

type UserGroupQueryBuilder struct {
	ResourceQueryBuilder
}

func (b *UserGroupQueryBuilder) ApiQuery() (*int64, interface{}, error) {
	return b.Query()
}

func (b *UserGroupQueryBuilder) Query() (*int64, []*models.UserGroup, error) {
	condition := b.build()
	count, groups, err := models.GetUserGroupList(func(o orm.QuerySeter) orm.QuerySeter {
		querySetter := o.SetCond(condition).Limit(b.pageOption.PageSize).Offset(b.pageOption.Offset())
		if len(b.orders) > 0 {
			querySetter = querySetter.OrderBy(b.orders...)
		}
		return querySetter
	})
	return &count, groups, err
}
