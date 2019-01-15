package serializer

import (
	"fmt"
	"yougame.com/yougame-server/models"
)

const (
	DefaultInventoryTemplateType = "DefaultInventoryTemplate"
)

func NewInventoryTemplate(templateType string) Template {
	return &DefaultInventoryTemplate{}
}

type DefaultInventoryTemplate struct {
	Id     int64      `json:"id" source:"Id" source_type:"int"`
	UserId int64      `json:"user_id" source:"User.Id" source_type:"int"`
	GoodId int64      `json:"good_id" source:"Good.Id" source_type:"int"`
	Link   []*ApiLink `json:"link"`
}

func (t *DefaultInventoryTemplate) Serialize(model interface{}, context map[string]interface{}) {
	data := model.(*models.InventoryItem)
	SerializeModelData(model, t)
	site := context["site"].(string)
	t.Link = []*ApiLink{
		{
			Rel:  "good",
			Href: fmt.Sprintf("%s/api/good/%d", site, data.Good.Id),
			Type: "GET",
		},
		{
			Rel:  "user",
			Href: fmt.Sprintf("%s/api/user/%d", site, data.User.Id),
			Type: "GET",
		},
	}
}
