package serializer

import (
	"fmt"
	"yougame.com/yougame-server/models"
)

const (
	DefaultCommentTemplateType = "DefaultCommentTemplateType"
)

func NewCommentTemplate(templateType string) Template {
	return &DefaultCommentSerializeTemplate{}
}

type DefaultCommentSerializeTemplate struct {
	Id       int        `json:"id" source:"Id" source_type:"int"`
	UserId   int        `json:"user_id" source:"User.Id" source_type:"int"`
	GoodId   int        `json:"good_id" source:"Good.Id" source_type:"int"`
	Content  string     `json:"content" source:"Content" source_type:"string"`
	Rating   int        `json:"rating" source:"Rating" source_type:"int"`
	CreateAt string     `json:"create_at" source:"Created" source_type:"string" converter:"time"`
	UpdateAt string     `json:"update_at" source:"Updated" source_type:"string" converter:"time"`
	Link     []*ApiLink `json:"link"`
}

func (t *DefaultCommentSerializeTemplate) Serialize(model interface{}, context map[string]interface{}) {
	data := model.(*models.Comment)
	SerializeModelData(model, t)
	site := context["site"].(string)
	t.Link = []*ApiLink{
		{
			Rel:  "user",
			Href: fmt.Sprintf("%s/api/user/%d", site, data.User.Id),
			Type: "GET",
		},
		{
			Rel:  "good",
			Href: fmt.Sprintf("%s/api/good/%d", site, data.Good.Id),
			Type: "GET",
		},
	}
}

type CommentSummarySerializeTemplate struct {
	Rating []*models.CommentRatingCountResult `json:"rating"`
}
