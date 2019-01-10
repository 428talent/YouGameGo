package service

import (
	"github.com/astaxie/beego/orm"
	"yougame.com/yougame-server/models"
)

type CartQueryBuilder struct {
	pageOption *PageOption
	ids        []interface{}
	userIds    []interface{}
	goodIds    []interface{}
	enable     string
}

func (b *CartQueryBuilder) ApiQuery() (*int64, interface{}, error) {
	count, result, err := b.Query()
	return &count, result, err
}

func (b *CartQueryBuilder) InId(id ...interface{}) {
	b.ids = append(b.ids, id...)
}
func (b *CartQueryBuilder) WithEnable(visibility string) {
	b.enable = visibility
}
func (b *CartQueryBuilder) InUser(id ...interface{}) {
	b.userIds = append(b.userIds, id...)
}
func (b *CartQueryBuilder) WithGood(id ...interface{}) {
	b.goodIds = append(b.goodIds, id...)
}
func (b *CartQueryBuilder) SetPage(page int64, pageSize int64) {
	b.pageOption = &PageOption{
		Page:     page,
		PageSize: pageSize,
	}
}

func (b *CartQueryBuilder) Query() (int64, []*models.CartItem, error) {
	cond := orm.NewCondition()

	if len(b.userIds) > 0 {
		cond = cond.And("user_id__in", b.userIds...)
	}
	if len(b.goodIds) > 0 {
		cond = cond.And("good_id__in", b.goodIds...)
	}

	if len(b.enable) > 0 {
		switch b.enable {
		case "visit":
			cond = cond.And("enable", true)
		case "remove":
			cond = cond.And("enable", false)
		}

	}
	if b.pageOption == nil {
		b.pageOption = &PageOption{
			Page:     1,
			PageSize: 10,
		}
	}
	count, cartItems, err := models.GetCartList(func(o orm.QuerySeter) orm.QuerySeter {
		return o.SetCond(cond).Limit(b.pageOption.PageSize).Offset(b.pageOption.Offset())
	})
	return count, cartItems, err
}
