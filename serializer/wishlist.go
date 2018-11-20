package serializer

import (
	"fmt"
	"yougame.com/yougame-server/models"
)

type WishListModel struct {
	Id      int    `json:"id"`
	User    string `json:"user"`
	Game    string `json:"game"`
	Created int64  `json:"created"`
}

func (s *WishListModel) SerializeData(model *models.WishList, site string) {
	s.Id = model.Id
	s.User = fmt.Sprintf("%s/api/user/%d", site, model.UserId)
	s.Game = fmt.Sprintf("%s/api/game/%d", site, model.Game.Id)
	s.Created = model.Created.Unix()
}
