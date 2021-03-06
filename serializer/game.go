package serializer

import (
	"fmt"
	"time"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/util"
)

const (
	DefaultGameTemplateType = "GameTemplate"
	AdminGameTemplateType   = "AdminGameTemplate"
)

func init() {

	AddCustomConverter("time", func(value interface{}) interface{} {
		timeValue, _ := value.(time.Time)
		return util.FormatApiTime(timeValue)
	})
	AddCustomConverter("date", func(value interface{}) interface{} {
		timeValue, _ := value.(time.Time)
		return util.FormatDate(timeValue)
	})

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
		}},
	}
	return goodModel
}

type GameTemplate struct {
	Id          int        `json:"id"  source_type:"int"`
	Name        string     `json:"name" source_type:"string"`
	Price       string    `json:"price" source:"Price" source_type:"string" converter:"money"`
	ReleaseTime string     `json:"release_time" source:"ReleaseTime" source_type:"string" converter:"date"`
	Publisher   string     `json:"publisher" source_type:"string"`
	Intro       string     `json:"intro" source_type:"string"`
	Link        []*ApiLink `json:"link"`
}

func (t *GameTemplate) Serialize(model interface{}, context map[string]interface{}) {
	data := model.(*models.Game)
	SerializeModelData(model, t)
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

type AdminGameTemplate struct {
	Id          int        `json:"id"  source_type:"int"`
	Name        string     `json:"name" source_type:"string"`
	ReleaseTime string     `json:"release_time" source:"ReleaseTime" source_type:"string" converter:"date"`
	Publisher   string     `json:"publisher" source_type:"string"`
	Intro       string     `json:"intro" source_type:"string"`
	Price       string    `json:"price" source:"Price" source_type:"string" converter:"money"`
	Created     string     `json:"created" source:"Created" source_type:"string" converter:"time"`
	Updated     string     `json:"updated" source:"Updated" source_type:"string" converter:"time"`
	Enable      bool       `json:"enable" source:"Enable" source_type:"bool"`
	Link        []*ApiLink `json:"link"`
}

func (t *AdminGameTemplate) CustomSerialize(convertTag string, value interface{}) interface{} {
	return value
}

func (t *AdminGameTemplate) Serialize(model interface{}, context map[string]interface{}) {
	data := model.(*models.Game)
	SerializeModelData(model, t)
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

func NewGameTemplate(templateType string) Template {
	switch templateType {
	case DefaultGameTemplateType:
		return &GameTemplate{}
	case AdminGameTemplateType:
		return &AdminGameTemplate{}
	default:
		return &GameTemplate{}
	}
}
