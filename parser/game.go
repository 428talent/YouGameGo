package parser

import (
	"encoding/json"
	"yougame.com/yougame-server/models"
	"yougame.com/yougame-server/util"
	"yougame.com/yougame-server/validate"
)

type AddGoodRequestBody struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (r *AddGoodRequestBody) Parse(body []byte) error {
	err := json.Unmarshal(body, r)
	return err
}

type AddGameTagRequestBody struct {
	Tags []string `json:"tags"`
}

func (r *AddGameTagRequestBody) Parse(body []byte) error {
	err := json.Unmarshal(body, r)
	return err
}

type CreateGameRequestBody struct {
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	ReleaseTime string  `json:"release_time"`
	Publisher   string  `json:"publisher"`
	Intro       string  `json:"intro"`
}

func (r *CreateGameRequestBody) Parse(body []byte) error {
	err := json.Unmarshal(body, r)
	return err
}

func (r *CreateGameRequestBody) Validate() error {
	err := validate.ValidateData(*r)
	return err
}

type ModifyGameRequestBody struct {
	Name        string  `json:"name" field:"name"`
	Price       float32 `json:"price" field:"price"`
	ReleaseTime string  `json:"release_time" field:"release_time"`
	Publisher   string  `json:"publisher"  field:"publisher"`
	Intro       string  `json:"intro" field:"intro"`
	Enable      bool    `json:"enable" field:"enable"`
}

func (r *ModifyGameRequestBody) ApplyToGame(gameId int64) (*models.Game, error) {
	game := &models.Game{
		Id:        int(gameId),
		Price:     r.Price,
		Name:      r.Name,
		Publisher: r.Publisher,
		Intro:     r.Intro,
		Enable:    r.Enable,
	}
	if len(r.ReleaseTime) != 0 {
		releaseTime, err := util.ParseDate(r.ReleaseTime)
		if err != nil {
			return nil, err
		}
		game.ReleaseTime = releaseTime
	}

	return game, nil
}
