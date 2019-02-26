package service

import (
	"github.com/astaxie/beego/orm"
	"yougame.com/yougame-server/models"
)

type PermissionQueryBuilder struct {
	ResourceQueryBuilder
}

func (b *PermissionQueryBuilder) ApiQuery() (*int64, interface{}, error) {
	return b.Query()
}

func (b *PermissionQueryBuilder) Query() (*int64, []*models.Permission, error) {
	condition := b.build()
	count, permissionList, err := models.GetPermissionList(func(o orm.QuerySeter) orm.QuerySeter {
		querySetter := o.SetCond(condition).Limit(b.pageOption.PageSize).Offset(b.pageOption.Offset())
		if len(b.orders) > 0 {
			querySetter = querySetter.OrderBy(b.orders...)
		}
		return querySetter
	})
	return &count, permissionList, err
}
