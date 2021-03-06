package game

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/parser"
	"yougame.com/yougame-server/request"
	"yougame.com/yougame-server/security"
	"yougame.com/yougame-server/serializer"
	"yougame.com/yougame-server/service"
	"yougame.com/yougame-server/util"
)

type Controller struct {
	api.ApiController
}

func (c *Controller) CreateGame() {
	c.WithErrorContext(func() {
		view := api.CreateView{
			Controller: &c.ApiController,
			Parser:     &parser.CreateGameRequestBody{},
			Permissions: []api.PermissionInterface{
				CreateGamePermission{},
			},
			Model:         &models.Game{},
			ModelTemplate: serializer.NewGameTemplate(serializer.AdminGameTemplateType),
			OnPrepareSave: func(c *api.CreateView) {
				model := c.Model.(*models.Game)
				requestParser := c.Parser.(*parser.CreateGameRequestBody)
				releaseTime, err := time.Parse("2006/1/2", requestParser.ReleaseTime)
				if err != nil {
					panic(api.ParseJsonDataError)
				}
				model.ReleaseTime = releaseTime
			},
		}
		err := view.Exec()
		if err != nil {
			panic(err)
		}
	})
}

func (c *Controller) GetGameBand() {
	imageQueryBuilder := service.ImageQueryBuilder{}
	c.WithErrorContext(func() {
		listView := api.ObjectView{
			Controller:    &c.ApiController,
			LookUpField:   "-",
			QueryBuilder:  &imageQueryBuilder,
			ModelTemplate: &serializer.ImageAdminTemplate{},
			SetFilter: func(builder service.ApiQueryBuilder) {
				gameIdParam := c.Controller.Ctx.Input.Param(":id")
				imageQueryBuilder.WithName("band:" + gameIdParam)

			},
		}
		err := listView.Exec()
		if err != nil {
			panic(err)
		}
	})
}

func (c *Controller) UploadGameBand() {
	c.WithErrorContext(func() {
		gameId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
		if err != nil {
			panic(err)
		}
		_, err = security.ParseAuthHeader(c.Controller)
		if err != nil {
			panic(service.NoAuthError)
		}

		f, h, err := c.GetFile("image")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		path := "static/upload/img/" + util.EncodeFileName(h.Filename)
		err = c.SaveToFile("image", path)
		if err != nil {
			panic(err)
		}
		imageType := c.GetString("type", "desktop")
		image, err := service.SaveGameBangImage(int64(gameId), path, imageType)
		if err != nil {
			panic(err)
		}
		imageTemplate := serializer.ImageTemplate{}
		imageTemplate.Serialize(image, map[string]interface{}{})
		c.Data["json"] = imageTemplate
		c.ServeJSON()
	})

}
func (c *Controller) UploadGamePreviewImage() {
	var err error
	defer api.CheckError(func(e error) {
		logrus.Error(err)
		api.HandleApiError(c.Controller, err)
	})
	gameId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		panic(err)
		return
	}
	_, err = security.ParseAuthHeader(c.Controller)
	if err != nil {
		panic(err)
	}

	game := models.Game{Id: gameId}
	err = game.QueryById()
	if err != nil {
		panic(err)
	}
	f, h, err := c.GetFile("image")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	path := "static/upload/img/" + util.EncodeFileName(h.Filename)
	err = c.SaveToFile("image", path)
	if err != nil {
		panic(err)
	}
	image, err := game.SavePreviewImage(path)
	template := serializer.ImageAdminTemplate{}
	template.Serialize(image, map[string]interface{}{})
	c.Data["json"] = template
	c.ServeJSON()
}

func (c *Controller) GetGood() {
	c.WithErrorContext(func() {
		gameId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
		if err != nil {
			panic(err)
		}
		listView := api.ListView{
			Controller:    &c.ApiController,
			ModelTemplate: serializer.NewGoodSerializeTemplate(serializer.DefaultGoodTemplateType),
			QueryBuilder:  &service.GoodQueryBuilder{},
			SetFilter: func(builder service.ApiQueryBuilder) {
				goodQueryBuilder, _ := builder.(*service.GoodQueryBuilder)
				goodQueryBuilder.InGameId(gameId)
			},
		}
		err = listView.Exec()
		if err != nil {
			panic(err)
		}
	})
}

func (c *Controller) AddTags() {
	c.WithErrorContext(func() {
		gameId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
		if err != nil {
			panic(err)
		}
		_, err = security.ParseAuthHeader(c.Controller)
		if err != nil {
			panic(err)
		}
		var requestBodyStruct parser.AddGameTagRequestBody

		err = requestBodyStruct.Parse(c.Ctx.Input.RequestBody)
		if err != nil {
			panic(err)

		}

		err = service.AddGameTags(gameId, requestBodyStruct.Tags...)
		if err != nil {
			panic(err)
		}
		tagTemplate := &serializer.CommonApiResponseBody{
			Success: true,
		}
		c.Data["json"] = tagTemplate
		c.ServeJSON()
	})

}
func (c *Controller) DeleteTags() {
	c.WithErrorContext(func() {
		gameId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
		if err != nil {
			panic(err)
		}
		_, err = security.ParseAuthHeader(c.Controller)
		if err != nil {
			panic(err)
		}
		var requestBodyStruct parser.AddGameTagRequestBody

		err = requestBodyStruct.Parse(c.Ctx.Input.RequestBody)
		if err != nil {
			panic(err)

		}

		err = service.DeleteGameTags(gameId, requestBodyStruct.Tags...)
		if err != nil {
			panic(err)
		}
		tagTemplate := &serializer.CommonApiResponseBody{
			Success: true,
		}
		c.Data["json"] = tagTemplate
		c.ServeJSON()
	})

}

func (c *Controller) AddGood() {
	c.WithErrorContext(func() {
		gameId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
		if err != nil {
			panic(err)
		}
		_, err = c.GetAuth()
		if err != nil {
			panic(api.ClaimsNoFoundError)
		}
		var requestBodyStruct request.AddGoodRequestBody
		err = json.Unmarshal(c.Ctx.Input.RequestBody, &requestBodyStruct)
		if err != nil {
			panic(err)
		}

		game := models.Game{Id: gameId}
		err = game.QueryById()
		if err != nil {
			panic(err)
		}
		good := models.Good{
			Name:  requestBodyStruct.Name,
			Price: requestBodyStruct.Price,
			Game:  &game,
		}
		err = game.AddGood(&good)
		if err != nil {
			panic(err)
		}
		goodTemplate := serializer.GoodSerializeTemplate{}
		goodTemplate.Serialize(&good, map[string]interface{}{
			"site": util.GetSiteAndPortUrl(c.Controller),
		})
		c.Data["json"] = goodTemplate
		c.ServeJSON()
	})

}

func (c *Controller) PutGame() {
	c.WithErrorContext(func() {
		claims, err := c.GetAuth()
		if err != nil {
			panic(api.ClaimsNoFoundError)
		}

		// check permission
		err = c.CheckPermission([]api.PermissionInterface{
			&UpdateGamePermission{},
		}, map[string]interface{}{
			"claims": *claims,
		})
		if err != nil {
			panic(err)
		}

		gameId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
		if err != nil {
			panic(err)
		}
		requestBody := parser.ModifyGameRequestBody{}
		err = json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
		if err != nil {
			panic(api.ParseJsonDataError)
		}

		// parse time
		game, err := requestBody.ApplyToGame(int64(gameId))
		if err != nil {
			panic(api.ParseJsonDataError)
		}
		err = service.UpdateGame(game, "name", "publisher", "intro", "release_time", "price")
		if err != nil {
			panic(err)
		}
		//parse json

		if game.Band != nil {
			err = game.ReadGameBand()
			if err != nil {
				panic(err)
			}
		} else {
			game.Band = &models.Image{
				Path: "",
			}
		}

		serializeTemplate := serializer.NewGameTemplate(serializer.AdminGameTemplateType)
		serializeTemplate.Serialize(game, map[string]interface{}{
			"site": util.GetSiteAndPortUrl(c.Controller),
		})
		c.Data["json"] = serializeTemplate
		c.ServeJSON()
	})
}
func (c *Controller) DeleteGame() {
	c.WithErrorContext(func() {
		deleteView := api.DeleteView{
			Controller: &c.ApiController,
			Model:      &models.Game{},
			Permissions: []api.PermissionInterface{
				&DeleteGamePermission{},
			},
		}
		err := deleteView.Exec()
		if err != nil {
			panic(err)
		}
	})
}
func (c *Controller) PatchGame() {
	c.WithErrorContext(func() {
		claims, err := c.GetAuth()
		if err != nil {
			panic(api.ClaimsNoFoundError)
		}

		// check permission
		err = c.CheckPermission([]api.PermissionInterface{
			&UpdateGamePermission{},
		}, map[string]interface{}{
			"claims": *claims,
		})
		if err != nil {
			panic(err)
		}

		gameId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
		if err != nil {
			panic(err)
		}
		requestBody := parser.ModifyGameRequestBody{}
		err = json.Unmarshal(c.Ctx.Input.RequestBody, &requestBody)
		if err != nil {
			panic(api.ParseJsonDataError)
		}

		// parse time
		updateFields := util.GetUpdateModelField(&requestBody)
		game, err := requestBody.ApplyToGame(int64(gameId))
		if err != nil {
			panic(api.ParseJsonDataError)
		}
		err = service.UpdateGame(game, updateFields...)
		if err != nil {
			panic(err)
		}

		//parse json
		if game.Band != nil {
			err = game.ReadGameBand()
			if err != nil {
				panic(err)
			}
		} else {
			game.Band = &models.Image{
				Path: "",
			}
		}

		serializeTemplate := serializer.NewGameTemplate(serializer.AdminGameTemplateType)
		serializeTemplate.Serialize(game, map[string]interface{}{
			"site": util.GetSiteAndPortUrl(c.Controller),
		})
		c.Data["json"] = serializeTemplate
		c.ServeJSON()
	})
}
func (c *Controller) UpdateGame() {
	c.WithErrorContext(func() {
		updateView := api.UpdateView{
			Controller: &c.ApiController,
			Parser:     &parser.ModifyGameRequestBody{},
			Model:      &models.Game{},
			Permissions: []api.PermissionInterface{
				&UpdateGamePermission{},
			},
			ModelTemplate: serializer.NewGameTemplate(serializer.AdminGameTemplateType),
		}
		err := updateView.Exec()
		if err != nil {
			panic(err)
		}
	})
}

func (c *Controller) GetGame() {
	c.WithErrorContext(func() {

		claims, err := c.GetAuth()
		if claims != nil && err == nil {
			// login user
			_ = c.User.ReadUserGroup()
			for _, group := range c.User.UserGroups {
				if group.Name == security.UserGroupAdmin {
					c.Role = security.UserGroupAdmin

				}
			}
		} else {
			c.Role = security.Anonymous
		}
		var renderResult interface{}
		switch c.Role {
		case security.UserGroupAdmin:
			view := AdminGetGameView{
				Controller: c,
			}
			renderResult = view.Render()
		case security.Anonymous:
			view := DefaultGetGameView{
				Controller: c,
			}
			renderResult = view.Render()
		default:
			view := DefaultGetGameView{
				Controller: c,
			}
			renderResult = view.Render()
		}

		c.Data["json"] = renderResult
		c.ServeJSON()
	})
}

func (c *Controller) GetGamePreview() {
	c.WithErrorContext(func() {
		listView := api.ListView{
			Controller:    &c.ApiController,
			QueryBuilder:  &service.ImageQueryBuilder{},
			ModelTemplate: &serializer.ImageTemplate{},
			Init: func() {
				c.GetAuth()
			},
			SetFilter: func(builder service.ApiQueryBuilder) {
				imageQueryBuilder := builder.(*service.ImageQueryBuilder)
				gameId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
				if err != nil {
					panic(err)
				}
				imageQueryBuilder.WithName(fmt.Sprintf("Preview:%d", gameId))
				enable := "visit"
				if security.CheckUserGroup(c.User, security.UserGroupAdmin) {
					enable = c.GetString("enable", "visit")
					if enable != "all" {
						imageQueryBuilder.WithEnable(enable)
					}
				}

			},
		}
		err := listView.Exec()
		if err != nil {
			panic(err)
		}
	})
}

func (c *Controller) GetTags() {
	c.WithErrorContext(func() {
		gameId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
		if err != nil {
			panic(err)
		}
		page, pageSize := c.GetPage()
		tagQueryBuilder := service.TagQueryBuilder{}
		tagQueryBuilder.SetPage(page, pageSize)
		tagQueryBuilder.WithGame(gameId)
		count, tags, err := tagQueryBuilder.Query()
		if err != nil {
			panic(err)
		}
		template := serializer.DefaultTagTemplate{}
		c.ServerPageResult(serializer.SerializeMultipleTemplate(tags, &template, map[string]interface{}{}), *count, page, pageSize)
	})
}

func (c *Controller) GetGameList() {
	c.WithErrorContext(func() {
		listView := api.ListView{
			Controller:    &c.ApiController,
			QueryBuilder:  &service.GameQueryBuilder{},
			ModelTemplate: serializer.NewGameTemplate(serializer.DefaultGameTemplateType),
			GetTemplate: func() serializer.Template {
				if c.Role == security.UserGroupAdmin {
					return serializer.NewGameTemplate(serializer.AdminGameTemplateType)
				}
				return serializer.NewGameTemplate(serializer.DefaultGameTemplateType)
			},
			Init: func() {
				c.GetAuth()
				c.Role = security.Anonymous
				if security.CheckUserGroup(c.User, security.UserGroupAdmin) {
					c.Role = security.UserGroupAdmin
				}
			},
			SetFilter: func(builder service.ApiQueryBuilder) {
				gameQueryBuilder := builder.(*service.GameQueryBuilder)
				util.FilterByParam(&c.Controller, "order", builder, "ByOrder", false)
				enable := "visit"
				if security.CheckUserGroup(c.User, security.UserGroupAdmin) {
					enable = c.GetString("enable", "visit")
				}
				gameQueryBuilder.WithEnable(enable)
				util.FilterByParam(&c.Controller, "name", builder, "SearchWithName", true)
				util.FilterByParam(&c.Controller, "tag", builder, "InTag", false)
				util.FilterByParam(&c.Controller, "id", gameQueryBuilder, "InId", false)
				util.FilterByParam(&c.Controller, "collection", gameQueryBuilder, "InGameCollection", false)
				util.FilterByParam(&c.Controller, "good", gameQueryBuilder, "InGood", false)
				priceStartParam := c.GetString("priceStart", "")
				if len(priceStartParam) > 0 {
					priceStart, err := strconv.ParseFloat(priceStartParam, 64)
					if err != nil {
						panic(err)
					}
					gameQueryBuilder.InPriceStart(priceStart)

				}

				priceEndParam := c.GetString("priceEnd", "")
				if len(priceEndParam) > 0 {
					priceEnd, err := strconv.ParseFloat(priceEndParam, 64)
					if err != nil {
						panic(err)
					}
					gameQueryBuilder.InPriceEnd(priceEnd)
				}

				util.FilterByParam(&c.Controller, "releaseTimeStart", gameQueryBuilder, "InReleaseTimeStart", true)
				util.FilterByParam(&c.Controller, "releaseTimeEnd", gameQueryBuilder, "InReleaseTimeEnd", true)

			},
		}
		err := listView.Exec()
		if err != nil {
			panic(err)
		}
	})
}

func (c *Controller) DeleteGameMultiple() {
	multipleView := api.DeleteMultipleView{
		Controller: &c.ApiController,
		Builder:    &service.GameQueryBuilder{},
		SetFilter: func(v *api.DeleteMultipleView) {
			builder := v.Builder.(*service.GameQueryBuilder)
			type deleteDatasRequestBody struct {
				Ids []interface{} `json:"ids"`
			}
			requestBody := &deleteDatasRequestBody{}

			err := json.Unmarshal(v.Controller.Ctx.Input.RequestBody, requestBody)
			if err != nil {
				panic(api.ParseJsonDataError)
			}
			builder.InId(requestBody.Ids...)
		},
	}
	err := multipleView.Exec()
	if err != nil {
		panic(err)
	}
}

func (c *Controller) UpdateGameMultiple() {
	multipleView := api.UpdateMultipleView{
		Controller: &c.ApiController,
		Model:      &models.Game{},
	}
	err := multipleView.Exec()
	if err != nil {
		panic(err)
	}
}
