package parser

import "yougame.com/yougame-server/validate"

type CreateCommentModel struct {
	Content string `json:"content" valid:"Required;MinSize(15);MaxSize(500)"`
	Rating  int64  `json:"rating"`
	GoodId  int64  `json:"good_id"`
}

func (c *CreateCommentModel) Validate() error {
	return validate.ValidateData(*c)
}

type UpdateCommentParser struct {
	Content string `json:"content" valid:"Required;MinSize(15);MaxSize(500)"`
	Rating  int64  `json:"rating"`
}