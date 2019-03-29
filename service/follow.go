package service

import (
	"github.com/astaxie/beego/orm"
	"yougame.com/yougame-server/models"
)

type FollowQueryBuilder struct {
	ResourceQueryBuilder
	userIds      []interface{}
	followingIds []interface{}
}

func (b *FollowQueryBuilder) ApiQuery() (*int64, interface{}, error) {
	return b.Query()
}

func (b *FollowQueryBuilder) InUser(userIds ...interface{}) {
	b.userIds = append(b.userIds, userIds...)
}

func (b *FollowQueryBuilder) InFollowing(followingIds ...interface{}) {
	b.followingIds = append(b.followingIds, followingIds...)
}

func (b *FollowQueryBuilder) Query() (*int64, []*models.Follow, error) {
	condition := b.build()
	if len(b.followingIds) > 0 {
		condition = condition.And("following_id__in", b.followingIds...)
	}
	if len(b.userIds) > 0 {
		condition = condition.And("user_id__in", b.userIds...)
	}
	return models.GetFollowList(func(o orm.QuerySeter) orm.QuerySeter {
		querySetter := o.SetCond(condition).Limit(b.pageOption.PageSize).Offset(b.pageOption.Offset())
		if len(b.orders) > 0 {
			querySetter = querySetter.OrderBy(b.orders...)
		}
		return querySetter
	})
}
