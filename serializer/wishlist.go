package serializer

import (
	"fmt"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/util"
)

const (
	DefaultWishListTemplateType = "DefaultWishListTemplateType"
)

type WishListModel struct {
	Id      int        `json:"id"`
	UserId  int        `json:"user_id"`
	GameId  int        `json:"game_id"`
	Created string     `json:"created"`
	Links   []*ApiLink `json:"links"`
}

func (s *WishListModel) Serialize(model interface{}, context map[string]interface{}) {
	wishlistItem := model.(*models.WishList)

	s.Id = wishlistItem.Id
	s.UserId = wishlistItem.UserId
	s.GameId = wishlistItem.Game.Id
	s.Created = util.FormatApiTime(wishlistItem.Created)
	site := context["site"].(string)
	s.Links = []*ApiLink{
		&ApiLink{
			Rel:  RelUser,
			Href: fmt.Sprintf("%s/api/user/%d", site, wishlistItem.UserId),
			Type: "GET",
		},
		&ApiLink{
			Rel:  RelGame,
			Href: fmt.Sprintf("%s/api/game/%d", site, wishlistItem.Game.Id),
			Type: "GET",
		},
	}
}

func NewWishlistTemplate(templateType string) Template {
	switch templateType {
	case templateType:
		return &WishListModel{}
	}
	return &WishListModel{}
}
