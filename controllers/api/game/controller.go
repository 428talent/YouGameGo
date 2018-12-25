package game

import (
	"encoding/json"
	"github.com/astaxie/beego"
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
	permission := []api.ApiPermissionInterface{
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
	var err error
	defer api.CheckError(func(e error) {
		logrus.Error(e)
		api.HandleApiError(c.Controller, err)
	})
	goodId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		panic(err)
	}
	good, err := service.GetGoodById(goodId)
	if err != nil {
		panic(err)
	}

	serializeTemplate := serializer.NewGameTemplate(serializer.DefaultGameTemplateType)
	serializeTemplate.Serialize(good, map[string]interface{}{
		"site": util.GetSiteAndPortUrl(c.Controller),
	})
	c.Data["json"] = serializeTemplate
	c.ServeJSON()
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
