package service

import (
	"github.com/astaxie/beego/orm"
	"yougame.com/yougame-server/models"
)

type PermissionQueryBuilder struct {
	ResourceQueryBuilder
	userGroupIds []interface{}
	searchName   string
}

func (b *PermissionQueryBuilder) ApiQuery() (*int64, interface{}, error) {
	return b.Query()
}
func (b *PermissionQueryBuilder) WithUserGroup(ids ...interface{}) {
	b.userGroupIds = append(b.userGroupIds, ids...)
}

func (b *PermissionQueryBuilder) WithName(name string) {
	b.searchName = name
}

func (b *PermissionQueryBuilder) Query() (*int64, []*models.Permission, error) {
	condition := b.build()
	if len(b.userGroupIds) > 0 {
		condition = condition.And("UserGroups__user_group_id__in", b.userGroupIds...)
	}
	if len(b.searchName) > 0 {
		condition = condition.And("name__icontains", b.searchName)
	}
	count, permissionList, err := models.GetPermissionList(func(o orm.QuerySeter) orm.QuerySeter {
		querySetter := o.SetCond(condition).Limit(b.pageOption.PageSize).Offset(b.pageOption.Offset())
		if len(b.orders) > 0 {
			querySetter = querySetter.OrderBy(b.orders...)
		}
		return querySetter
	})
	return &count, permissionList, err
}
