package serializer

import (
	"fmt"
	"yougame.com/yougame-server/models"
)

const (
	// default order template
	DefaultOrderTemplateType = "DefaultOrderTemplateType"
)

type OrderTemplate struct {
	Id      int        `json:"id" source:"Id" source_type:"int"`
	State   string     `json:"state" source:"State" source_type:"string"`
	UserId  int        `json:"user_id" source:"User.Id" source_type:"int"`
	Created string     `json:"created" source:"Created" source_type:"string" converter:"time"`
	Updated string     `json:"updated" source:"Updated" source_type:"string" converter:"time"`
	Link    []*ApiLink `json:"link"`
}

func (t *OrderTemplate) Serialize(model interface{}, context map[string]interface{}) {
	order := model.(*models.Order)
	SerializeModelData(order, t)
	site := context["site"].(string)
	t.Link = append(t.Link, &ApiLink{
		Rel:  "goods",
		Href: fmt.Sprintf("%s/api/order/%d/goods", site, order.Id),
		Type: "GET",
	}, &ApiLink{
		Rel:  "user",
		Href: fmt.Sprintf("%s/api/user/%d", site, order.User.Id),
		Type: "GET",
	})
}

func NewOrderTemplate(templateType string) Template {
	return &OrderTemplate{}
}

const (
	DefaultOrderGoodTemplateType = "DefaultOrderGoodTemplateType"
)

func NewOrderGoodTemplate(templateType string) Template {
	return &OrderItemTemplate{}
}

type OrderItemTemplate struct {
	Id       int        `json:"id" source:"Id" source_type:"int"`
	GoodName string     `json:"good_name" source:"GoodName" source_type:"string"`
	Name     string     `json:"name" source:"Name" source_type:"string"`
	GameId   int        `json:"game_id"source:"Game.Id" source_type:"int"`
	Price    float64    `json:"price" source:"Price" source_type:"float"`
	BandPic  string     `json:"band_pic" source:"BandPic" source_type:"string"`
	Created  int64      `json:"created" source:"Created" source_type:"string" converter:"time"`
	Link     []*ApiLink `json:"link"`
}

func (t *OrderItemTemplate) Serialize(model interface{}, context map[string]interface{}) {
	orderGood := model.(*models.OrderGood)
	SerializeModelData(orderGood, t)
	site := context["site"].(string)
	t.Link = append(t.Link, &ApiLink{
		Rel:  "order",
		Href: fmt.Sprintf("%s/api/order/%d", site, orderGood.Order.Id),
		Type: "GET",
	})
}
