package parser

import "encoding/json"

type CreateGameCollectionRequestBody struct {
	Name   string `json:"name"`
	Title  string `json:"title"`
	Enable bool   `json:"enable"`
}

type AddGameRequestBody struct {
	Games []int `json:"games"`
}

func (r *AddGameRequestBody) Parse(body []byte) error {
	err := json.Unmarshal(body, r)
	return err
}

