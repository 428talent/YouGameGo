package service

import (
	"github.com/astaxie/beego/orm"
	"yougame.com/yougame-server/models"
)

type GetOrderListBuilder struct {
	user     int64
	state    []string
	page     int64
	pageSize int64
}

func (builder *GetOrderListBuilder) SetUser(userId int64) *GetOrderListBuilder {
	builder.user = userId
	return builder
}
func (builder *GetOrderListBuilder) SetState(state []string) *GetOrderListBuilder {
	builder.state = state
	return builder
}
func (builder *GetOrderListBuilder) build() *orm.Condition {
	cond := orm.NewCondition()
	if builder.user != 0 {
		cond = cond.And("user_id", builder.user)
	}
	if len(builder.state) > 0 {
		stateCond := orm.NewCondition()
		for _, state := range builder.state {
			stateCond = stateCond.Or("state", state)
		}
		cond = cond.AndCond(stateCond)
	}
	if builder.page == 0 {
		builder.page = 1
	}
	if builder.pageSize == 0 {
		builder.page = 10
	}
	return cond
}

func GetOrderList(builder GetOrderListBuilder) (int64, []*models.Order, error) {
	count, orders, err := models.GetOrderList(func(o orm.QuerySeter) orm.QuerySeter {
		cond := builder.build()
		return o.SetCond(cond).Limit(builder.pageSize).Offset((builder.page - 1) * builder.pageSize)
	})
	return count, orders, err
}
