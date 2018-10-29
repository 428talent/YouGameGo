package serializer

import (
	"yougame.com/yougame-server/models"
)

type OrderSerializer struct {
	Id      int
	Goods   []*OrderItemSerializer
	State   string
	Created int64
	Updated int64
}
type OrderItemSerializer struct {
	Id       int
	GoodName string
	Name     string
	GameId   int
	Price    float64
	BandPic  string
	Created  int64
}

func SerializeOrder(data models.Order, template interface{}) (interface{}, error) {
	switch template.(type) {
	case OrderSerializer:
		err := data.ReadOrderGoods()
		if err != nil {
			return nil, err
		}
		var goodList []*OrderItemSerializer
		for _, orderGood := range data.Goods {
			err = orderGood.ReadGood()
			if err != nil {
				return nil, err
			}
			err = orderGood.Good.ReadGame()
			if err != nil {
				return nil, err
			}
			err = orderGood.Good.Game.ReadGameBand()
			if err != nil {
				return nil, err
			}
			goodList = append(goodList, &OrderItemSerializer{
				Id:       orderGood.Id,
				GoodName: orderGood.Name,
				Name:     orderGood.Good.Game.Name,
				Price:    orderGood.Price,
				GameId:   orderGood.Good.Game.Id,
				BandPic:  orderGood.Good.Game.Band.Path,
				Created:  orderGood.Created.Unix(),
			})
		}
		order := OrderSerializer{
			Id:      data.Id,
			State:   data.State,
			Goods:   goodList,
			Created: data.Created.Unix(),
			Updated: data.Updated.Unix(),
		}
		return order, nil

	}
	return nil, nil
}
func SerializeOrderList(data []*models.Order, template interface{}) ([]*interface{}, error) {
	var result []*interface{}
	for _, order := range data {
		item, err := SerializeOrder(*order, template)
		if err != nil {
			return nil, err
		}
		result = append(result, &item)
	}
	return result, nil
}
