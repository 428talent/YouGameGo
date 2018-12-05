package serializer

import (
	"fmt"
	"yougame.com/yougame-server/models"
)

type CartSerializer struct {
	Id   int `json:"id"`
	Good struct {
		Id    int     `json:"id"`
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	} `json:"good"`
	Game struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Band string `json:"band"`
	} `json:"game"`
	Created int64 `json:"created"`
}

func SerializeCart(data models.CartItem, template interface{}) (interface{}, error) {
	switch template.(type) {
	case CartSerializer:
		return &CartSerializer{
			Id: data.Id,
			Game: struct {
				Id   int    `json:"id"`
				Name string `json:"name"`
				Band string `json:"band"`
			}{data.Good.Game.Id, data.Good.Game.Name, data.Good.Game.Band.Path},
			Good: struct {
				Id    int     `json:"id"`
				Name  string  `json:"name"`
				Price float64 `json:"price"`
			}{data.Good.Id, data.Good.Name, data.Good.Price},
			Created: data.Created.Unix(),
		}, nil

	}
	return nil, nil
}

func SerializeCartList(data []*models.CartItem, template interface{}) ([]*interface{}, error) {
	var result []*interface{}
	for _, cart := range data {
		item, err := SerializeCart(*cart, template)
		if err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}

type CartModel struct {
	Id      int        `json:"id"`
	GoodId  int        `json:"good_id"`
	UserId  int        `json:"user_id"`
	Created int64      `json:"created"`
	Link    []*ApiLink `json:"link"`
}

func (c *CartModel) SerializeData(model interface{}, site string) interface{} {
	cartItem := model.(models.CartItem)
	data := CartModel{
		Id:      cartItem.Id,
		GoodId:  cartItem.Good.Id,
		UserId:  cartItem.User.Id,
		Created: cartItem.Created.Unix(),
		Link: []*ApiLink{
			{
				Rel:  "good",
				Href: fmt.Sprintf("%s/api/good/%d", site, cartItem.Good.Id),
				Type: "GET",
			}, {
				Rel:  "user",
				Href: fmt.Sprintf("%s/api/user/%d", site, cartItem.User.Id),
				Type: "GET",
			},
		},
	}
	return data

}
