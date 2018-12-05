package service

import (
	"github.com/astaxie/beego/orm"
	"yougame.com/yougame-server/models"
)

type CartQueryBuilder struct {
	pageOption *PageOption
	ids        []interface{}
	userIds    []interface{}
}

func (b *CartQueryBuilder) InId(id ...interface{}) {
	b.ids = append(b.ids, id...)
}
func (b *CartQueryBuilder) InUser(id ...interface{}) {
	b.ids = append(b.ids, id...)
}
func (b *CartQueryBuilder) SetPage(page int64, pageSize int64) {
	b.pageOption = &PageOption{
		Page:     page,
		PageSize: pageSize,
	}
}

func (b *CartQueryBuilder) Query() (int64, []*models.CartItem, error) {
	cond := orm.NewCondition()
	count, cartItems, err := models.GetCartList(func(o orm.QuerySeter) orm.QuerySeter {
		return o.SetCond(cond).Limit(b.pageOption.PageSize).Offset(b.pageOption.Offset())
	})
	return count, cartItems, err
}
