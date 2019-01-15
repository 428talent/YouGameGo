package comment

import (
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/parser"
	"yougame.com/yougame-server/serializer"
	"yougame.com/yougame-server/service"
)

type ApiCommentController struct {
	api.ApiController
}

func (c *ApiCommentController) GetCommentList() {
	c.WithErrorContext(func() {
		listView := api.ListView{
			Controller:    &c.ApiController,
			QueryBuilder:  &service.CommentQueryBuilder{},
			ModelTemplate: serializer.NewCommentTemplate(serializer.DefaultCommentTemplateType),
			SetFilter: func(builder service.ApiQueryBuilder) {
				commentQueryBuilder := builder.(*service.CommentQueryBuilder)
				for _, gameId := range c.GetStrings("game") {
					commentQueryBuilder.SetGame(gameId)
				}
				for _, goodId := range c.GetStrings("good") {
					commentQueryBuilder.SetGood(goodId)
				}
				for _, userId := range c.GetStrings("user") {
					commentQueryBuilder.SetUser(userId)
				}
			},
		}
		err := listView.Exec()
		if err != nil {
			panic(err)
		}
	})

}

func (c *ApiCommentController) CreateComment() {
	c.WithErrorContext(func() {
		createView := api.CreateView{
			Controller:    &c.ApiController,
			Parser:        &parser.CreateCommentModel{},
			Model:         &models.Comment{},
			ModelTemplate: serializer.NewCommentTemplate(serializer.DefaultCommentTemplateType),
			OnPrepareSave: func(c *api.CreateView) {
				commentModel := c.Model.(*models.Comment)
				requestData := c.Parser.(*parser.CreateCommentModel)
				commentModel.User = c.Controller.User
				commentModel.Good = &models.Good{
					Id: int(requestData.GoodId),
				}
				commentModel.Enable = true
			},
		}
		err := createView.Exec()
		if err != nil {
			panic(err)
		}
	})
}
