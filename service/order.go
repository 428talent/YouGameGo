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
func (builder *GetOrderListBuilder) SetPage(page int64) *GetOrderListBuilder {
	builder.page = page
	return builder
}
func (builder *GetOrderListBuilder) SetPageSize(pageSize int64) *GetOrderListBuilder {
	builder.pageSize = pageSize
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

func GetOrderById(id int) (*models.Order, error) {
	order := models.Order{Id: id,}
	err := order.QueryById()
	return &order, err
}

type GetOrderGoodListOption struct {
	order int64
	page  PageOption
}

func (option *GetOrderGoodListOption) SetOrder(orderId int64) *GetOrderGoodListOption {
	option.order = orderId
	return option
}
func (option *GetOrderGoodListOption) SetPage(pageOption PageOption) *GetOrderGoodListOption {
	option.page = pageOption
	return option
}
func (option *GetOrderGoodListOption) build() *orm.Condition {
	cond := orm.NewCondition()
	if option.order != 0 {
		cond = cond.And("order_id", option.order)
	}
	if option.page.Page == 0 {
		option.page.Page = 1
	}
	if option.page.PageSize == 0 {
		option.page.PageSize = 10
	}
	return cond
}
func (option *GetOrderGoodListOption) Query(md interface{}) (int64, error) {
	modelStruct :=  models.OrderGood{}
	count, err :=modelStruct.GetList(func(o orm.QuerySeter) orm.QuerySeter {
		cond := option.build()
		return o.SetCond(cond).Limit(option.page.PageSize).Offset((option.page.Page - 1) * option.page.PageSize)
	}, md)
	return count, err
}
