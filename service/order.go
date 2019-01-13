package service

import (
	"github.com/astaxie/beego/orm"
	"yougame.com/yougame-server/models"
)

type OrderQueryBuilder struct {
	user  []interface{}
	state []interface{}
	ResourceQueryBuilder
}

func (builder *OrderQueryBuilder) ApiQuery() (*int64, interface{}, error) {
	return builder.Query()
}

func (builder *OrderQueryBuilder) SetPage(page int64, pageSize int64) {
	builder.pageOption = &PageOption{
		PageSize: pageSize,
		Page:     page,
	}
}

func (builder *OrderQueryBuilder) InUser(userId ...interface{}) *OrderQueryBuilder {
	builder.user = append(builder.user, userId...)
	return builder
}
func (builder *OrderQueryBuilder) SetState(state ...interface{}) *OrderQueryBuilder {
	builder.state = append(builder.state, state...)
	return builder
}

func (builder *OrderQueryBuilder) buildQuery() *orm.Condition {
	cond := builder.build()
	if len(builder.user) > 0 {
		cond = cond.And("user_id__in", builder.user...)
	}
	if len(builder.state) > 0 {
		stateCond := orm.NewCondition()
		for _, state := range builder.state {
			stateCond = stateCond.Or("state", state)
		}
		cond = cond.AndCond(stateCond)
	}

	return cond
}

func (builder *OrderQueryBuilder)Query() (*int64, []*models.Order, error) {
	count, orders, err := models.GetOrderList(func(o orm.QuerySeter) orm.QuerySeter {
		cond := builder.buildQuery()
		setter := o.SetCond(cond).Limit(builder.pageOption.PageSize).Offset(builder.pageOption.Offset())
		if len(builder.orders) > 0 {
			setter = setter.OrderBy(builder.orders...)
		}
		return setter
	})
	return &count, orders, err
}

type OrderGoodQueryBuilder struct {
	ResourceQueryBuilder
}

func (builder *OrderGoodQueryBuilder) Query(md interface{}) (*int64, []*models.OrderGood, error) {
	cond := builder.build()
	count, result, err := models.GetOrderGoodList(func(o orm.QuerySeter) orm.QuerySeter {
		setter := o.SetCond(cond).Limit(builder.pageOption.PageSize).Offset(builder.pageOption.Offset())
		if len(builder.orders) > 0 {
			setter = setter.OrderBy(builder.orders...)
		}
		return setter
	})
	return &count, result, err
}

func CreateOrder(order *models.Order, goods []int64) error {
	o := orm.NewOrm()
	order.Enable = true
	order.State = "Create"
	transaction := func() error {
		err := o.Begin()
		if err != nil {
			return err
		}

		err = order.Save(o)
		if err != nil {
			return err
		}

		for _, goodId := range goods {
			good := &models.Good{}
			err := good.Query(goodId)
			if err != nil {
				return err
			}
			orderGood := &models.OrderGood{
				Good:  good,
				Order: order,
				Name:  good.Name,
				Price: good.Price,
			}
			_, err = o.Insert(orderGood)
			if err != nil {
				return err
			}
		}
		return nil

	}
	err := transaction()
	if err != nil {
		err := o.Rollback()
		if err != nil {
			return err
		}
	} else {
		err := o.Commit()
		if err != nil {
			return err
		}
	}
	return nil
}
