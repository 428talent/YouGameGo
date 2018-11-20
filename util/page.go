package util

import (
	"github.com/astaxie/beego"
)

func ParsePageRequest(c beego.Controller) (int64, int64) {
	page, err := c.GetInt64("page", 1)
	if err != nil {
		beego.Error(err)
		page = 1
	}
	pageSize, err := c.GetInt64("pageSize", 10)
	if err != nil {
		beego.Error(err)
		pageSize = 10
	}
	return page, pageSize
}

type PageResponse struct {
	Count    int64       `json:"count"`
	Page     int64       `json:"page"`
	PageSize int64       `json:"page_size"`
	NextPage *string      `json:"next_page"`
	PrevPage *string      `json:"prev_page"`
	Result   interface{} `json:"result"`
}


