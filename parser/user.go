package parser

import (
	"encoding/json"
)

type CreateUserRequestStruct struct {
	Username string `json:"username" valid:"Required;MinSize(4);MaxSize(16)"`
	Password string `json:"password" valid:"Required;MinSize(4);MaxSize(16)"`
	Email    string `json:"email" valid:"Required;"`
}

func (r *CreateUserRequestStruct) Parse(body []byte) error {
	err := json.Unmarshal(body, r)
	return err
}

type AuthTokenRequestBody struct {
	LoginName string `json:"login_name"`
	Password  string `json:"password"`
}

func (r *AuthTokenRequestBody) ParseGetTokenRequestBody(body []byte) (*AuthTokenRequestBody, error) {
	var result AuthTokenRequestBody
	err := json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type UploadUserAvatarRequestStruct struct {
	Avatar    string `json:"avatar"`
	ImageType string `json:"image_type"`
}

func (r *UploadUserAvatarRequestStruct) Parse(body []byte) error {
	err := json.Unmarshal(body, r)
	if err != nil {
		return err
	}
	return nil
}

type ChangeProfileRequestStruct struct {
	Nickname string `json:"nickname"`
}

func (r *ChangeProfileRequestStruct) Parse(body []byte) error {
	err := json.Unmarshal(body, r)
	if err != nil {
		return err
	}
	return nil
}

type ResetUserPasswordRequestStruct struct {
	Username string `json:"username"`
}

func (r *ResetUserPasswordRequestStruct) Parse(body []byte) error {
	err := json.Unmarshal(body, r)
	if err != nil {
		return err
	}
	return nil
}

type RecoveryPasswordRequestStruct struct {
	Code     int `json:"code"`
	Password string `json:"password"`
}

func (r *RecoveryPasswordRequestStruct) Parse(body []byte) error {
	err := json.Unmarshal(body, r)
	if err != nil {
		return err
	}
	return nil
}
