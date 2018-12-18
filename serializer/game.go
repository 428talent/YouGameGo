package serializer

import (
	"fmt"
	"yougame.com/yougame-server/models"
)

type Game struct {
	Id            int                 `json:"id"`
	Name          string              `json:"name"`
	ReleaseTime   int64               `json:"release_time"`
	Publisher     string              `json:"publisher"`
	Band          string              `json:"band"`
	Intro         string              `json:"intro"`
	Tags          []*GameTag          `json:"tags"`
	PreviewImages []*GamePreviewImage `json:"preview_images"`
	Goods         []*GameGood         `json:"goods"`
}
type GameTag struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type GamePreviewImage struct {
	Id   int    `json:"id"`
	Path string `json:"path"`
}
type GameGood struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (g *Game) Serialize(game models.Game) {
	g.Id = game.Id
	g.Name = game.Name
	g.ReleaseTime = game.ReleaseTime.Unix()
	g.Band = game.Band.Path
	g.Intro = game.Intro
	var previewImages []*GamePreviewImage
	g.Publisher = game.Publisher
	for _, image := range game.PreviewImages {
		previewImages = append(previewImages, &GamePreviewImage{
			Id:   image.Id,
			Path: image.Path,
		})
	}
	g.PreviewImages = previewImages

	var tags []*GameTag
	for _, tag := range game.Tags {
		tags = append(tags, &GameTag{
			Id:   tag.Id,
			Name: tag.Name,
		})
	}
	g.Tags = tags

	var goods []*GameGood
	for _, good := range game.Goods {
		goods = append(goods, &GameGood{
			Id:    good.Id,
			Name:  good.Name,
			Price: good.Price,
		})
	}
	g.Goods = goods

}

type GoodModel struct {
	Id     int        `json:"id"`
	GameId int        `json:"game_id"`
	Name   string     `json:"name"`
	Price  float64    `json:"price"`
	Link   []*ApiLink `json:"link"`
}

func (g *GoodModel) SerializeData(model interface{}, site string) interface{} {
	good := model.(*models.Good)
	goodModel := GoodModel{
		Id:     good.Id,
		GameId: good.Game.Id,
		Name:   good.Name,
		Price:  good.Price,
		Link: []*ApiLink{{
			Rel:  "game",
			Href: fmt.Sprintf("%s/api/game/%d", site, good.Game.Id),
			Type: "GET",
		},},
	}
	return goodModel
}

type GoodSerializeTemplate struct {
	Id     int        `json:"id" source:"Id" source_type:"int"`
	GameId int        `json:"game_id" source:"Game.Id" source_type:"int"`
	Name   string     `json:"name" source:"Name" source_type:"string"`
	Price  float64    `json:"price" source:"Price" source_type:"float"`
	Link   []*ApiLink `json:"link"`
}

func (t *GoodSerializeTemplate) CustomSerialize(convertTag string, value interface{}) interface{} {
	return value
}

func (t *GoodSerializeTemplate) Serialize(model interface{}, context map[string]interface{}) {
	data := model.(*models.Good)
	SerializeModelData(model, t)
	site := context["site"].(string)
	t.Link = []*ApiLink{
		{
			Rel:  "good",
			Href: fmt.Sprintf("%s/api/game/%d", site, data.Game.Id),
			Type: "GET",
		},
	}
}

type GameTemplate struct {
	Id          int    `json:"id"  source_type:"int"`
	Name        string `json:"name" source_type:"string"`
	ReleaseTime int64  `json:"release_time" source:"ReleaseTime.Unix()[0]" source_type:"int"`
	Publisher   string `json:"publisher" source_type:"string"`
	Intro       string `json:"intro" source_type:"string"`
	Link        []*ApiLink
}

func (t *GameTemplate) CustomSerialize(convertTag string, value interface{}) interface{} {
	return value
}

func (t *GameTemplate) Serialize(model interface{}, context map[string]interface{}) {
	data := model.(*models.Game)
	SerializeModelData(model,t)
	site := context["site"].(string)
	t.Link = []*ApiLink{
		{
			Rel:  "goods",
			Href: fmt.Sprintf("%s/api/game/%d/goods", site, data.Id),
			Type: "GET",
		},
		{
			Rel:  "band",
			Href: fmt.Sprintf("%s/api/game/%d/band", site, data.Id),
			Type: "GET",
		},
		{
			Rel:  "tags",
			Href: fmt.Sprintf("%s/api/game/%d/tags", site, data.Id),
			Type: "GET",
		},
		{
			Rel:  "preview_images",
			Href: fmt.Sprintf("%s/api/game/%d/preview", site, data.Id),
			Type: "GET",
		},
	}
}

