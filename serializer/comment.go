package serializer

import (
	"fmt"
	"yougame.com/yougame-server/models"
)

type CommentSerializeModel struct {
	Id         int    `json:"id"`
	UserId     int    `json:"user_id"`
	GoodId     int    `json:"good_id"`
	Content    string `json:"content"`
	Evaluation string `json:"evaluation"`
	CreateAt   int64  `json:"create_at"`
	UpdateAt   int64  `json:"update_at"`
	Link       []*ApiLink `json:"link"`
}

func (c *CommentSerializeModel) SerializeData(model interface{}, site string) interface{} {
	comment := model.(models.Comment)
	serializeData := CommentSerializeModel{
		Id:         comment.Id,
		UserId:     comment.User.Id,
		GoodId:     comment.Good.Id,
		Content:    comment.Content,
		Evaluation: comment.Evaluation,
		CreateAt:   comment.Created.Unix(),
		UpdateAt:   comment.Updated.Unix(),
		Link: []*ApiLink{
			{
				Rel:  "user",
				Href: fmt.Sprintf("%s/api/user/%d", site, comment.User.Id),
				Type: "GET",
			},
			{
				Rel:  "good",
				Href: fmt.Sprintf("%s/api/good/%d", site, comment.Good.Id),
				Type: "GET",
			},
		},
	}
	return serializeData
}
