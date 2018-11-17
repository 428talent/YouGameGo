package serializer

import "yougame.com/yougame-server/models"

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
