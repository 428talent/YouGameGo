package api_admin_game

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	"github.com/jinzhu/now"
	"yougame.com/yougame-server/models"
)

type GameSerializer struct {
	requestBody []byte
}

func (s *GameSerializer) produce() (newGame models.Game, err error) {

	if s.requestBody == nil {
		return newGame, errors.New("requestBody is null")
	}
	//解析json
	requestData := CreateGameRequest{}
	err = json.Unmarshal(s.requestBody, &requestData)
	if err != nil {
		beego.Error(err)
	}
	//检查输入值
	validation := GameCreateValidate{
		requestData: requestData,
	}
	err = validation.CheckValidation()
	if err != nil {
		return
	}

	//将值转换为model
	releaseTime, err := now.Parse(requestData.ReleaseTime)
	if err != nil {
		return
	}
	newGame = models.Game{
		Name:        requestData.Name,
		Price:       requestData.Price,
		ReleaseTime: releaseTime,
		Publisher:   requestData.Publisher,
		Intro:       requestData.Intro,
		Enable:      true,
	}
	return
}
