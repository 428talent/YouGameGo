package serializer

import "yougame.com/yougame-server/models"

type WishListSerializer struct {
	Id      int   `json:"id"`
	UserId  int   `json:"user_id"`
	Game    int   `json:"game_id"`
	Created int64 `json:"created"`
}

func SerializeWishList(data models.WishList, template interface{}) interface{} {
	switch template.(type) {
	case WishListSerializer:
		return WishListSerializer{
			Id:      data.Id,
			UserId:  data.UserId,
			Game:    data.Game.Id,
			Created: data.Created.Unix(),
		}
	}
	return nil
}

func SerializeWishListMultiple(data []*models.WishList, template interface{}) []*interface{} {
	var result []*interface{}
	for _, wishlist := range data {
		item := SerializeWishList(*wishlist, template)
		result = append(result, &item)
	}
	return result
}
