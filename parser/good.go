package parser

type UpdateGoodRequestBody struct {
	Name   string  `field:"name"`
	Price  float64 `field:"price"`
	Enable bool
}

type CreateGoodRequestBody struct {
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	GameId int64   `json:"game_id"`
}

type DeleteGoodRequestBody struct {
	Ids []int64 `json:"ids"`
}