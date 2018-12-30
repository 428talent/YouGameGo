package service

import "github.com/astaxie/beego/orm"

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
	ApiQuery() (*int64,interface{},error)
	InId(id ...interface{})
	SetPage(page int64, pageSize int64)
}