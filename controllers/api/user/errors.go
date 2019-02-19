package user

import (
	"net/http"
	ApiError "yougame.com/yougame-server/application"
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/service"
)

func init() {
	api.RegisterErrors(map[error]*ApiError.APIErrorResponse{
		service.UserExistError:    ExistError,
		service.EmailAlreadyExist: EmailExistError,
	})
}

var (
	ExistError = ApiError.NewApiError(ApiError.APIError{
		Err:    "UserExistError",
		Detail: "User already exist!",
		Code:   api.UserExistCode,
	}, http.StatusBadRequest)

	EmailExistError = ApiError.NewApiError(ApiError.APIError{
		Err:    "EmailAlreadyExistError",
		Detail: "email already exist",
		Code:   api.UserEmailExistCode,
	}, http.StatusBadRequest)
)
