package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"yougame.com/yougame-server/security"
	"yougame.com/yougame-server/util"
	"yougame.com/yougame-server/validate"
)

type ApiController struct {
	beego.Controller
	RequestStruct      interface{}
	SerializerTemplate interface{}
}

func (c ApiController) GetAuth() (*security.UserClaims, error) {
	if claims, err := security.ParseAuthHeader(c.Controller); err != nil {
		return nil, err
	} else {
		return claims, err
	}
}

func (c ApiController) ParseRequestBody() error {
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &c.RequestStruct)
	return err
}

func (c ApiController) ValidateRequestData() error {
	err := validate.ValidateData(c.RequestStruct)
	return err
}

func (c ApiController) SerializeData(data interface{}) {
	panic("serialize function not define")
}

func (c ApiController) GetPage() (page int64, pageSize int64) {
	page, pageSize = util.ParsePageRequest(c.Controller)
	return
}
