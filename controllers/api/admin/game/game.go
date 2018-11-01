package api_admin_game

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/sirupsen/logrus"
	"strconv"
	"yougame.com/letauth/controllers"
	"yougame.com/letauth/security"
	"yougame.com/letauth/util"
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/request"
)

type GameController struct {
	beego.Controller
}

type CreateGameRequest struct {
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	ReleaseTime string  `json:"release_time"`
	Publisher   string  `json:"publisher"`
	Intro       string  `json:"intro"`
}

//func (c *GameController) Get(){
//	gameId := c.GetInt("id")
//
//}
func (c *GameController) Post() {
	_, err := security.ParseAuthHeader(c.Controller)
	if err != nil {
		beego.Error(err)
	}
	serializer := GameSerializer{
		requestBody: c.Ctx.Input.RequestBody,
	}
	newGame, err := serializer.produce()
	if err != nil {
		panic(api.ParseJsonDataError)
	}
	err = newGame.Save()
	if err != nil {
		panic(err)
	}
	defer func() {
		troubleMaker := recover()
		if troubleMaker != nil {
			err = troubleMaker.(error)
			switch err {
			default:
				logrus.Error(err)
				api.HandleApiError(c.Controller, err)
			}
		} else {
			c.Data["json"] = newGame
			c.ServeJSON()
		}

	}()

}

func (c *GameController) UploadGameBand() {
	gameId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		beego.Error(err)
		controllers.AbortServerError(c.Controller)
		return
	}
	_, err = security.ParseAuthHeader(c.Controller)
	if err != nil {
		beego.Error(err)
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

	//err = os.Remove(user.Profile.Avatar)
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
	gameId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		beego.Error(err)
		controllers.AbortServerError(c.Controller)
		return
	}
	_, err = security.ParseAuthHeader(c.Controller)
	if err != nil {
		beego.Error(err)
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

	//err = os.Remove(user.Profile.Avatar)
	path := "static/upload/img/" + util.EncodeFileName(h.Filename)
	err = c.SaveToFile("image", path)
	if err != nil {
		beego.Error(err)
	}
	err = game.SavePreviewImage(path)
	beego.Error(err)
	c.Data["json"] = game
	c.ServeJSON()
}

func (c *GameController) AddTags() {
	gameId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		beego.Error(err)
		controllers.AbortServerError(c.Controller)
		return
	}
	_, err = security.ParseAuthHeader(c.Controller)
	if err != nil {
		beego.Error(err)
	}
	var requestBodyStruct request.AddGameTagRequestBody
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &requestBodyStruct)
	if err != nil {
		beego.Error(err)
		return
	}

	game := models.Game{Id: gameId}
	err = game.QueryById()
	if err != nil {
		beego.Error(err)
	}
	err = game.SaveTags(requestBodyStruct.Tags)
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = game
	c.ServeJSON()
}

func (c *GameController) AddGood() {
	gameId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		beego.Error(err)
		controllers.AbortServerError(c.Controller)
		return
	}
	_, err = security.ParseAuthHeader(c.Controller)
	if err != nil {
		beego.Error(err)
	}
	var requestBodyStruct request.AddGoodRequestBody
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &requestBodyStruct)
	if err != nil {
		beego.Error(err)
		return
	}

	game := models.Game{Id: gameId}
	err = game.QueryById()
	if err != nil {
		beego.Error(err)
	}
	good := models.Good{
		Name:  requestBodyStruct.Name,
		Price: requestBodyStruct.Price,
		Game:  &game,
	}
	err = game.AddGood(good)
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = game
	c.ServeJSON()
}

func (c *GameController) GetGame() {
	gameId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		beego.Error(err)
		controllers.AbortServerError(c.Controller)
		return
	}

	game := models.Game{
		Id: gameId,
	}
	err = game.QueryById()
	if err != nil {
		beego.Error(err)
		controllers.AbortServerError(c.Controller)
		return
	}

	game.ReadGameBand()
	game.ReadGamePreviewImage()
	game.ReadTags()
	game.ReadGoods()
	c.Data["json"] = game
	c.ServeJSON()
}
