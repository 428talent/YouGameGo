package serializer

import (
	"fmt"
	"yougame.com/yougame-server/models"
)

const (
	DefaultCartTemplateType = "DefaultCartTemplate"
)

type CartTemplate struct {
	Id      int        `json:"id" source:"Id" source_type:"int"`
	GoodId  int        `json:"good_id" source:"Good.Id" source_type:"int"`
	UserId  int        `json:"user_id" source:"User.Id" source_type:"int"`
	Created string     `json:"created" source:"Created" source_type:"string" converter:"time"`
	Link    []*ApiLink `json:"link"`
}

func (t *CartTemplate) CustomSerialize(convertTag string, value interface{}) interface{} {
	return value
}

func (t *CartTemplate) Serialize(model interface{}, context map[string]interface{}) {
	data := model.(*models.CartItem)
	SerializeModelData(model, t)
	site := context["site"].(string)
	t.Link = []*ApiLink{{
		Rel:  "good",
		Href: fmt.Sprintf("%s/api/good/%d", site, data.Good.Id),
		Type: "GET",
	}, {
		Rel:  "user",
		Href: fmt.Sprintf("%s/api/user/%d", site, data.User.Id),
		Type: "GET",
	}}
}

func NewCartTemplate(templateType string) Template {
	switch templateType {
	case DefaultGameTemplateType:
		return &CartTemplate{}
	}
	return &CartTemplate{}
}
