package serializer

import (
	"fmt"
	"yougame.com/yougame-server/models"
)

type WishListModel struct {
	Id      int        `json:"id"`
	UserId  int        `json:"user_id"`
	GameId  int        `json:"game_id"`
	Created int64      `json:"created"`
	Links   []*ApiLink `json:"links"`
}

func (s *WishListModel) SerializeData(model interface{}, site string) interface{} {
	wishlistItem := model.(models.WishList)
	serializeData := WishListModel{
		Id:      wishlistItem.Id,
		UserId:  wishlistItem.UserId,
		GameId:  wishlistItem.Game.Id,
		Created: wishlistItem.Created.Unix(),
	}

	serializeData.Links = append(serializeData.Links, &ApiLink{
		Rel:  RelUser,
		Href: fmt.Sprintf("%s/api/user/%d", site, wishlistItem.UserId),
		Type: "GET",
	})
	serializeData.Links = append(serializeData.Links, &ApiLink{
		Rel:  RelGame,
		Href: fmt.Sprintf("%s/api/game/%d", site, wishlistItem.Game.Id),
		Type: "GET",
	})
	return serializeData
}
