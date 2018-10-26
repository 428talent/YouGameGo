package parser

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

type CreateOrderParser struct {
	Goods []int64 `json:"goods"`
}

func (p *CreateOrderParser) Parse(c beego.Controller) error {
	err := json.Unmarshal(c.Ctx.Input.RequestBody, p)
	return err
}
