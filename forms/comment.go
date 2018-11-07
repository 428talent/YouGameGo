package forms

type CreateCommentForm struct {
	Content string `form:"content"`
	GoodId int `form:"goodId"`
	Evaluation string `form:"evaluation"`
}