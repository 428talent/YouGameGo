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

//func (c *DefaultCommentSerializeTemplate) SerializeData(model interface{}, site string) interface{} {
//	comment := model.(models.Comment)
//	serializeData := DefaultCommentSerializeTemplate{
//		Id:         comment.Id,
//		UserId:     comment.User.Id,
//		GoodId:     comment.Good.Id,
//		Content:    comment.Content,
//		Evaluation: comment.Evaluation,
//		CreateAt:   comment.Created.Unix(),
//		UpdateAt:   comment.Updated.Unix(),
//		Link: []*ApiLink{
//			{
//				Rel:  "user",
//				Href: fmt.Sprintf("%s/api/user/%d", site, comment.User.Id),
//				Type: "GET",
//			},
//			{
//				Rel:  "good",
//				Href: fmt.Sprintf("%s/api/good/%d", site, comment.Good.Id),
//				Type: "GET",
//			},
//		},
//	}
//	return serializeData
//}
