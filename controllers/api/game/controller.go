package game

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/sirupsen/logrus"
	"reflect"
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

type GameController struct {
	api.ApiController
}

func (c *GameController) CreateGame() {
	var err error
	defer api.CheckError(func(e error) {
		logrus.Error(e)
		api.HandleApiError(c.Controller, err)
	})
	// read authentic
	claims, err := c.GetAuth()
	if err != nil {
		panic(security.ReadAuthorizationFailed)
	}

	//check permission
	permissionContext := map[string]interface{}{
		"claims": *claims,
	}
	permission := []api.PermissionInterface{
		CreateGamePermission{},
	}
	err = c.CheckPermission(permission, permissionContext)
	if err != nil {
		panic(err)
	}
	//parse request body
	requestBodyStruct := parser.CreateGameRequestBody{}
	if err = requestBodyStruct.Parse(c.Ctx.Input.RequestBody); err != nil {
		panic(api.ParseJsonDataError)
	}
	releaseTime, err := time.Parse("2006/1/2", requestBodyStruct.ReleaseTime)
	if err != nil {
		panic(api.ParseJsonDataError)
	}
	//handle
	game, err := service.CreateNewGame(
		requestBodyStruct.Name,
		requestBodyStruct.Price,
		requestBodyStruct.Intro,
		requestBodyStruct.Publisher,
		releaseTime,
	)
	if err != nil {
		panic(err)
	}

	template := serializer.GameTemplate{}
	template.Serialize(game, map[string]interface{}{
		"site": util.GetSiteAndPortUrl(c.Controller),
	})
	c.Data["json"] = template
	c.ServeJSON()

}

func (c *GameController) UploadGameBand() {
	gameId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		panic(err)
	}
	_, err = security.ParseAuthHeader(c.Controller)
	if err != nil {
		panic(service.NoAuthError)
	}

	game := models.Game{Id: gameId}
	err = game.QueryById()
	if err != nil {
		beego.Error(err)
	}
	f, h, err := c.GetFile("image")
	if err != nil {
		beego.Error(err)
	}
	defer f.Close()

	path := "static/upload/img/" + util.EncodeFileName(h.Filename)
	err = c.SaveToFile("image", path)
	if err != nil {
		beego.Error(err)
	}
	image, err := game.SaveGameBangImage(path)
	imageTemplate := serializer.ImageTemplate{}
	imageTemplate.Serialize(image, map[string]interface{}{})
	c.Data["json"] = imageTemplate
	c.ServeJSON()
}

func (c *GameController) UploadGamePreviewImage() {
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
	err = game.SavePreviewImage(path)
	c.Data["json"] = game
	c.ServeJSON()
}
func (c *GameController) GetGood() {
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

func (c *GameController) AddTags() {
	var err error
	defer api.CheckError(func(e error) {
		logrus.Error(err)
		api.HandleApiError(c.Controller, err)
	})
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

	tags, err := service.AddGameTags(gameId, requestBodyStruct.Tags...)
	if err != nil {
		panic(err)
	}
	tagTemplate := serializer.TagTemplate{}
	c.Data["json"] = serializer.SerializeMultipleTemplate(tags, &tagTemplate, map[string]interface{}{})
	c.ServeJSON()
}

func (c *GameController) AddGood() {
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

func (c *GameController) PutGame() {
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

func (c *GameController) PatchGame() {
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
		beego.Debug(updateFields)
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
func (c *GameController) GetGame() {
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
		}

		c.Data["json"] = renderResult
		c.ServeJSON()
	})
}

func (c *GameController) GetGameBand() {
	c.WithErrorContext(func() {
		gameId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
		if err != nil {
			panic(err)
		}
		image, err := service.GetGameBand(gameId)
		if err != nil {
			panic(err)
		}
		template := serializer.ImageTemplate{}
		template.Serialize(image, map[string]interface{}{})
		c.Data["json"] = template
		c.ServeJSON()
	})
}

func (c *GameController) GetGamePreview() {
	c.WithErrorContext(func() {
		gameId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
		if err != nil {
			panic(err)
		}
		page, pageSize := c.GetPage()
		count, image, err := service.GetGamePreview(gameId, page, pageSize)
		if err != nil {
			panic(err)
		}
		template := serializer.ImageTemplate{}
		c.ServerPageResult(serializer.SerializeMultipleTemplate(image, &template, map[string]interface{}{}), *count, page, pageSize)
	})
}

func (c *GameController) GetTags() {
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
		template := serializer.TagTemplate{}
		c.ServerPageResult(serializer.SerializeMultipleTemplate(tags, &template, map[string]interface{}{}), *count, page, pageSize)
	})
}

func (c *GameController) GetGameList() {
	c.WithErrorContext(func() {
		listView := api.ListView{
			Controller:    &c.ApiController,
			QueryBuilder:  &service.GameQueryBuilder{},
			ModelTemplate: serializer.NewGameTemplate(serializer.DefaultGameTemplateType),
			OnGetResult: func(gameList interface{}) {
				if c.Role != security.UserGroupAdmin {
					return
				}
				ref := reflect.ValueOf(gameList)
				for idx := 0; idx < ref.Len(); idx++ {
					game := ref.Index(idx).Interface().(*models.Game)
					if game.Band == nil {
						game.Band = &models.Image{
							Path: "",
						}
					} else {
						err := game.ReadGameBand()
						if err != nil {
							beego.Debug(err)
						}
					}

				}
			},
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
				orders := c.GetStrings("order")
				if len(orders) > 0 {
					gameQueryBuilder.ByOrder(orders...)
				}

				if security.CheckUserGroup(c.User, security.UserGroupAdmin) {
					enable := c.GetString("enable", "visit")
					if enable != "all" {
						gameQueryBuilder.WithEnable(enable)
					}
				}

				if name := c.GetString("name"); len(name) > 0 {
					gameQueryBuilder.SearchWithName(name)
				}

			},
		}
		err := listView.Exec()
		if err != nil {
			panic(err)
		}
	})
}
