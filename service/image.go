package service

import (
	"github.com/astaxie/beego/orm"
	"yougame.com/yougame-server/models"
)

type ImageQueryBuilder struct {
	pageOption *PageOption
	ids        []interface{}
	names      []interface{}
	enable     string
}

func (q *ImageQueryBuilder) ApiQuery() (*int64, interface{}, error) {
	return q.Query()
}

func (q *ImageQueryBuilder) SetPage(page int64, pageSize int64) {
	q.pageOption = &PageOption{Page: page, PageSize: pageSize}
}

func (q *ImageQueryBuilder) InId(id ...interface{}) {
	q.ids = append(q.ids, id...)
}

func (q *ImageQueryBuilder) WithName(name ...interface{}) {
	q.names = append(q.names, name...)
}
func (b *ImageQueryBuilder) WithEnable(visibility string) {
	b.enable = visibility
}
func (q *ImageQueryBuilder) Query() (*int64, []*models.Image, error) {
	condition := orm.NewCondition()
	if len(q.ids) > 0 {
		condition = condition.And("id__in", q.ids...)
	}

	if len(q.names) > 0 {
		condition = condition.And("name__in", q.names...)
	}
	if len(q.enable) > 0 {
		switch q.enable {
		case "visit":
			condition = condition.And("enable", true)
		case "remove":
			condition = condition.And("enable", false)
		}

	}
	if q.pageOption == nil {
		q.pageOption = &PageOption{
			Page:     1,
			PageSize: 20,
		}
	}
	return models.GetImageList(func(o orm.QuerySeter) orm.QuerySeter {
		return o.SetCond(condition).Limit(q.pageOption.PageSize).Offset(q.pageOption.Offset())
	})
}
