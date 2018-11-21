package serializer

import (
	"fmt"
	"yougame.com/yougame-server/models"
)

type OrderModel struct {
	Id      int        `json:"id"`
	State   string     `json:"state"`
	UserId  int      `json:"user_id"`
	Created int64      `json:"created"`
	Updated int64      `json:"updated"`
	Link    []*ApiLink `json:"link"`
}
type OrderItemSerializer struct {
	Id       int     `json:"id"`
	GoodName string  `json:"good_name"`
	Name     string  `json:"name"`
	GameId   int     `json:"game_id"`
	Price    float64 `json:"price"`
	BandPic  string  `json:"band_pic"`
	Created  int64   `json:"created"`
}

func (o *OrderModel) SerializeData(model interface{}, site string) interface{} {
	order := model.(models.Order)
	serializeData := OrderModel{
		Id:      order.Id,
		State:   string(order.State),
		UserId:  order.User.Id,
		Created: order.Created.Unix(),
		Updated: order.Updated.Unix(),
	}
	serializeData.Link = append(serializeData.Link, &ApiLink{
		Rel:  "goods",
		Href: fmt.Sprintf("%s/api/order/%d/goods", site, order.Id),
		Type: "GET",
	}, &ApiLink{
		Rel:  "user",
		Href: fmt.Sprintf("%s/api/user/%d", site, order.User.Id),
		Type: "GET",
	})
	return serializeData
}
