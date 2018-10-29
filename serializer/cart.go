package serializer

import "yougame.com/yougame-server/models"

type CartSerializer struct {
	Id   int `json:"id"`
	Good struct {
		Id    int     `json:"id"`
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}
	Game struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Band string `json:"band"`
	}
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
				Id    int `json:"id"`
				Name  string `json:"name"`
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
