package request

type AddGameTagRequestBody struct {
	Tags []string `json:"tags"`
}

type AddGoodRequestBody struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
