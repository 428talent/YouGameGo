package service

import (
	"github.com/astaxie/beego/orm"
	"yougame.com/yougame-server/models"
)

type PageOption struct {
	Page     int64
	PageSize int64
}

func (p *PageOption) Offset() int64 {
	return (p.Page - 1) * p.PageSize
}

type ResourcesQueryBuilder interface {
	build() *orm.Condition
}

type QueryBuilder interface {
	SetPage(page int64, pageSize int64)
	InId(id ...interface{})
}

type ApiQueryBuilder interface {
	ApiQuery() (*int64, interface{}, error)
	InId(id ...interface{})
	SetPage(page int64, pageSize int64)
}

func DeleteData(model models.DataModel) error {
	o := orm.NewOrm()
	return model.Delete(o)
}
func UpdateData(id int64, model models.DataModel, fields ...string) error {
	o := orm.NewOrm()
	return model.Update(id, o, fields...)
}
func SaveData(model models.DataModel) error {
	o := orm.NewOrm()
	return model.Save(o)
}
type ResourceQueryBuilder struct {
	ids        []interface{}
	pageOption *PageOption
	enable     string
	orders     []string
}
func (b *ResourceQueryBuilder) build() *orm.Condition {
	condition := orm.NewCondition()
	if len(b.ids) > 0 {
		condition = condition.And("id__in", b.ids...)
	}
	if b.pageOption == nil {
		b.pageOption = &PageOption{
			Page:     1,
			PageSize: 10,
		}
	}
	if len(b.enable) > 0 {
		switch b.enable {
		case "visit":
			condition = condition.And("enable", true)
		case "remove":
			condition = condition.And("enable", false)
		}

	}
	return condition
}



func (b *ResourceQueryBuilder) SetPage(page int64, pageSize int64) {
	b.pageOption = &PageOption{
		PageSize: pageSize,
		Page:     page,
	}
}

func (b *ResourceQueryBuilder) InId(id ...interface{}) {
	b.ids = append(b.ids, id...)
}

func (b *ResourceQueryBuilder) WithEnable(visibility string) {
	b.enable = visibility
}

func (b *ResourceQueryBuilder) ByOrder(orders ...string) {
	b.orders = append(b.orders, orders...)
}


type UserIdOption struct {
	userIds []interface{}
}

func (builder *UserIdOption) InUser(userId ...interface{}) {
	builder.userIds = append(builder.userIds, userId...)
}

type NameOption struct {
	names []interface{}
}

func (builder *NameOption) WithName(name ...interface{}) {
	builder.names = append(builder.names, name...)
}

