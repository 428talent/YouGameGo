package collection

import (
	"strconv"
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/parser"
	"yougame.com/yougame-server/serializer"
	"yougame.com/yougame-server/service"
)

type Controller struct {
	api.ApiController
}

func (c *Controller) GetGameCollectionList() {
	c.WithErrorContext(func() {
		listView := api.ListView{
			Controller:    &c.ApiController,
			QueryBuilder:  &service.GameCollectionQueryBuilder{},
			ModelTemplate: serializer.NewGameCollectionTemplate(serializer.DefaultGameCollectionTemplateType),
			SetFilter: func(builder service.ApiQueryBuilder) {
				gameCollectionQueryBuilder := builder.(*service.GameCollectionQueryBuilder)
				for _, orderParam := range c.GetStrings("order") {
					gameCollectionQueryBuilder.ByOrder(orderParam)
				}
			},
		}
		err := listView.Exec()
		if err != nil {
			panic(err)
		}
	})
}

func (c *Controller) Create() {
	c.WithErrorContext(func() {
		createView := api.CreateView{
			Controller:    &c.ApiController,
			Parser:        &parser.CreateGameCollectionRequestBody{},
			Model:         &models.GameCollection{},
			ModelTemplate: serializer.NewGameCollectionTemplate(serializer.DefaultGameCollectionTemplateType),
		}
		err := createView.Exec()
		if err != nil {
			panic(err)
		}
	})
}

func (c *Controller) DeleteGameCollection() {
	c.WithErrorContext(func() {
		deleteView := api.DeleteView{
			Controller: &c.ApiController,
			Model:      &models.GameCollection{},
		}
		err := deleteView.Exec()
		if err != nil {
			panic(err)
		}
	})
}

func (c *Controller) Update() {
	c.WithErrorContext(func() {
		updateView := api.UpdateView{
			Controller:    &c.ApiController,
			Model:         &models.GameCollection{},
			Parser:        &parser.CreateGameCollectionRequestBody{},
			ModelTemplate: serializer.NewGameCollectionTemplate(serializer.DefaultGameCollectionTemplateType),
		}
		err := updateView.Exec()
		if err != nil {
			panic(err)
		}
	})
}

func (c *Controller) AddGame() {
	c.WithErrorContext(func() {
		requestBody := parser.AddGameRequestBody{}
		err := requestBody.Parse(c.Ctx.Input.RequestBody)
		if err != nil {
			panic(api.ParseJsonDataError)
		}
		collectionId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
		if err != nil {
			panic(err)
		}

		err = service.AddGameToCollection(collectionId, requestBody.Games...)
		if err != nil {
			panic(err)
		}

		responseBody := serializer.CommonApiResponseBody{
			Success: true,
		}
		c.Data["json"] = responseBody
		c.ServeJSON()
	})
}

func (c *Controller) DeleteGame() {
	requestBody := parser.AddGameRequestBody{}
	err := requestBody.Parse(c.Ctx.Input.RequestBody)
	if err != nil {
		panic(api.ParseJsonDataError)
	}
	collectionId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		panic(err)
	}
	err = service.DeleteGameFromCollection(collectionId, requestBody.Games...)
	if err != nil {
		panic(err)
	}

	responseBody := serializer.CommonApiResponseBody{
		Success: true,
	}
	c.Data["json"] = responseBody
	c.ServeJSON()
}
