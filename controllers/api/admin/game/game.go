package api_admin_game

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/jinzhu/now"
	"you_game_go/models"
)

type GameController struct {
	beego.Controller
}

type CreateGameRequest struct {
	Name        string
	Price       float32
	ReleaseTime string
	Publisher   string
	Intro       string
}

func (c *GameController) Post() {
	requestData := CreateGameRequest{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &requestData)
	if err != nil {
		beego.Error(err)
	}
	releaseTime, err := now.Parse(requestData.ReleaseTime)
	if err != nil {
		beego.Error(err)
	}
	newGame := models.Game{
		Name:        requestData.Name,
		Price:       requestData.Price,
		ReleaseTime: releaseTime,
		Publisher:   requestData.Publisher,
		Intro:       requestData.Intro,
		Enable:      true,
	}
	err = newGame.Save()
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = newGame
	c.ServeJSON()
}
