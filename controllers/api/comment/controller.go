package comment

import (
	"strconv"
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/parser"
	"yougame.com/yougame-server/serializer"
	"yougame.com/yougame-server/service"
	"yougame.com/yougame-server/util"
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

				//commentQueryBuilder := builder.(*service.CommentQueryBuilder)
				util.FilterByParam(&c.Controller, "game", builder, "SetGame", false)
				util.FilterByParam(&c.Controller, "good", builder, "SetGood", false)
				util.FilterByParam(&c.Controller, "user", builder, "SetUser", false)
				util.FilterByParam(&c.Controller, "rating", builder, "WithRating", false)
				util.FilterByParam(&c.Controller, "order", builder, "ByOrder", false)
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
func (c *ApiCommentController) Update() {
	c.WithErrorContext(func() {
		updateView := api.UpdateView{
			Controller:    &c.ApiController,
			Parser:        &parser.UpdateCommentParser{},
			Model:         &models.Comment{},
			ModelTemplate: serializer.NewCommentTemplate(serializer.DefaultCommentTemplateType),
		}
		err := updateView.Exec()
		if err != nil {
			panic(err)
		}
	})
}
func (c *ApiCommentController) GetCommentSummary() {
	c.WithErrorContext(func() {
		gameId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
		if err != nil {
			panic(err)
		}
		result, err := service.GetCommentSummary(gameId)
		c.Data["json"] = result
		c.ServeJSON()
	})
}
