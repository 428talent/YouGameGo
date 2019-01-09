package parser

type DeleteWishlistItems struct {
	Items []int `json:"items"`
}

type CreateWishlistRequestBody struct {
	GameId int64 `json:"game_id"`
}