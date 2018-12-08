package game

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/sirupsen/logrus"
	"strconv"
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

//func (c *GameController) Get(){
//	gameId := c.GetInt("id")
//
//}
func (c *GameController) Post() {
	var err error
	defer api.CheckError(func(e error) {
		logrus.Error(err)
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

	//handle
	newGame := models.Game{
		Name:      requestBodyStruct.Name,
		Price:     requestBodyStruct.Price,
		Intro:     requestBodyStruct.Intro,
		Publisher: requestBodyStruct.Publisher,
	}
	err = newGame.Save()
	if err != nil {
		panic(err)
	}
	c.Data["json"] = newGame
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
	err = game.SaveGameBangImage(path)
	beego.Error(err)
	c.Data["json"] = game
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
		logrus.Error(err)
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
	goodModel := serializer.GoodModel{}
	c.Data["json"] = goodModel.SerializeData(good, util.GetSiteAndPortUrl(c.Controller))
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

	game := models.Game{Id: gameId}
	err = game.QueryById()
	if err != nil {
		panic(err)
	}
	err = game.SaveTags(requestBodyStruct.Tags)
	if err != nil {
		panic(err)
	}
	c.Data["json"] = game
	c.ServeJSON()
}

func (c *GameController) AddGood() {
	c.WithErrorContext(func() {
		gameId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
		if err != nil {
			panic(err)
		}
		_, err = security.ParseAuthHeader(c.Controller)
		if err != nil {
			panic(err)
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
		err = game.AddGood(good)
		if err != nil {
			panic(err)
		}
		c.Data["json"] = game
		c.ServeJSON()
	})

}

func (c *GameController) GetGame() {
	c.WithErrorContext(func() {
		gameId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
		if err != nil {
			panic(err)
		}
		queryBuilder := service.GameQueryBuilder{}
		queryBuilder.InId(gameId)
		count, result, err := queryBuilder.Query()
		if err != nil {
			panic(err)
		}
		if *count == 0 {
			panic(api.ResourceNotFoundError)
		}
		game := result[0]

		if err = game.ReadGameBand(); err != nil {
			panic(err)
		}
		if err = game.ReadGamePreviewImage(); err != nil {
			panic(err)
		}
		if err = game.ReadTags(); err != nil {
			panic(err)
		}
		if err = game.ReadGoods(); err != nil {
			panic(err)
		}
		serializeData := serializer.Game{}
		serializeData.Serialize(*game)
		c.Data["json"] = serializeData
		c.ServeJSON()
	})
}
