package parser

import (
	"encoding/json"
)

type CreateUserRequestStruct struct {
	Username string `json:"username" valid:"Required;MinSize(4);MaxSize(16)"`
	Password string `json:"password" valid:"Required;MinSize(4);MaxSize(16)"`
}

func (r *CreateUserRequestStruct) Parse(body []byte) error {
	err := json.Unmarshal(body, r)
	return err
}


type GetTokenRequestStruct struct {
	LoginName string `json:"login_name"`
	Password  string `json:"password"`
}

func (r *GetTokenRequestStruct)ParseGetTokenRequestBody(body []byte) (*GetTokenRequestStruct, error) {
	var result GetTokenRequestStruct
	err := json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}