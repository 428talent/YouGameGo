package parser

import "yougame.com/yougame-server/validate"

type CreateCommentModel struct {
	Content    string `json:"content" valid:"Required;MinSize(15);MaxSize(500)"`
	Evaluation string `json:"evaluation"`
}

func (c *CreateCommentModel) Validate() error {
	return validate.ValidateData(*c)
}
