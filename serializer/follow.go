package serializer

import (
	"fmt"
	"yougame.com/yougame-server/models"
)

const (
	DefaultFollowTemplateType = "DefaultFollowTemplate"
)

type FollowSerializeTemplate struct {
	Id          int        `json:"id" source:"Id" source_type:"int"`
	UserId      int        `json:"user_id" source:"User.Id" source_type:"int"`
	FollowingId int        `json:"following_id" source:"Following.Id" source_type:"int"`
	Link        []*ApiLink `json:"link"`
}

func (t *FollowSerializeTemplate) Serialize(model interface{}, context map[string]interface{}) {
	data := model.(*models.Follow)
	SerializeModelData(model, t)
	site := context["site"].(string)
	t.Link = []*ApiLink{
		{
			Rel:  "user",
			Href: fmt.Sprintf("%s/api/user/%d", site, data.User.Id),
			Type: "GET",
		},
		{
			Rel:  "follow",
			Href: fmt.Sprintf("%s/api/user/%d", site, data.Following.Id),
			Type: "GET",
		},
	}
}

func NewFollowSerializeTemplate(templateType string) Template {
	return &FollowSerializeTemplate{}
}
