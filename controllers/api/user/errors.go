package user

import (
	"net/http"
	ApiError "yougame.com/yougame-server/error"
)

var (
	UserExistError = ApiError.NewApiError(ApiError.APIError{
		Err:    "UserExistError",
		Detail: "User already exist!",
		Code:   "200001",
	}, http.StatusBadRequest)
)