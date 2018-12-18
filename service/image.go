package service

import (
	"github.com/astaxie/beego/orm"
	"yougame.com/yougame-server/models"
)

type ImageQueryBuilder struct {
	pageOption *PageOption
	ids        []interface{}
	names      []interface{}
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

func (q *ImageQueryBuilder) Query() (*int64, []*models.Image, error) {
	condition := orm.NewCondition()
	if len(q.ids) > 0 {
		condition.And("id__in", q.ids...)
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
