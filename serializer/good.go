package serializer

import (
	"fmt"
	"yougame.com/yougame-server/models"
)

const (
	DefaultGoodTemplateType = "DefaultGoodTemplate"
	AdminGoodTemplateType = "AdminGoodTemplate"
)

type GoodSerializeTemplate struct {
	Id     int        `json:"id" source:"Id" source_type:"int"`
	GameId int        `json:"game_id" source:"Game.Id" source_type:"int"`
	Name   string     `json:"name" source:"Name" source_type:"string"`
	Price  float64    `json:"price" source:"Price" source_type:"float"`
	Link   []*ApiLink `json:"link"`
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

type AdminGoodSerializeTemplate struct {
	Id      int        `json:"id" source:"Id" source_type:"int"`
	GameId  int        `json:"game_id" source:"Game.Id" source_type:"int"`
	Name    string     `json:"name" source:"Name" source_type:"string"`
	Price   float64    `json:"price" source:"Price" source_type:"float"`
	Created string     `json:"created" source:"Created" source_type:"string" converter:"time"`
	Updated string     `json:"updated" source:"Updated" source_type:"string" converter:"time"`
	Enable  bool       `json:"enable" source:"Enable" source_type:"bool"`
	Link    []*ApiLink `json:"link"`
}

func (t *AdminGoodSerializeTemplate) Serialize(model interface{}, context map[string]interface{}) {
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
func NewGoodSerializeTemplate(templateType string) Template {
	switch templateType {
	case AdminGoodTemplateType:
		return &AdminGoodSerializeTemplate{}
	}
	return &GoodSerializeTemplate{}
}
