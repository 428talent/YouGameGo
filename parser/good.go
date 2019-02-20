package parser

type UpdateGoodRequestBody struct {
	Name   string  `field:"name" mapstructure:"name"`
	Price  float64 `field:"price" mapstructure:"price"`
	Enable bool    `field:"enable" mapstructure:"enable"`
}

type CreateGoodRequestBody struct {
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	GameId int64   `json:"game_id"`
}

type DeleteGoodRequestBody struct {
	Ids []int64 `json:"ids"`
}
