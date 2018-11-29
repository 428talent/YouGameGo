package service

import "github.com/astaxie/beego/orm"

type PageOption struct {
	Page int64
	PageSize int64
}

type ResourcesQueryBuilder interface {
	build() *orm.Condition
}

