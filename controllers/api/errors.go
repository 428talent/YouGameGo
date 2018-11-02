package api

import (
	"errors"
	"github.com/astaxie/beego"
	"net/http"
	ApiError "yougame.com/yougame-server/error"
	"yougame.com/yougame-server/security"
	"yougame.com/yougame-server/service"
)

var (
	ServerError = ApiError.NewApiError(ApiError.APIError{
		Err:    "Server error",
		Detail: "Server error",
		Code:   "000003",
	}, http.StatusInternalServerError)

	AuthFailedError = ApiError.NewApiError(ApiError.APIError{
		Err:    "Authorization failed",
		Detail: "Authorization failed",
		Code:   "100000",
	}, http.StatusUnauthorized)

	ParseRequestDataError = ApiError.NewApiError(ApiError.APIError{
		Err:    "ParseRequestDataError",
		Detail: "Parse request data error,check request data",
		Code:   "100001",
	}, http.StatusBadRequest)

	PermissionNotAllowError = ApiError.NewApiError(ApiError.APIError{
		Err:    "PermissionNotAllowError",
		Detail: "permission denied",
		Code:   "100002",
	}, http.StatusForbidden)
)
var (
	ParseJsonDataError = errors.New("cannot parse json request")
)

func HandleApiError(controller beego.Controller, err error) {
	switch err {
	case service.NoAuthError:
		AuthFailedError.ServerError(controller)
	case ParseJsonDataError:
		ParseRequestDataError.ServerError(controller)
	case security.ReadAuthorizationFailed:
		AuthFailedError.ServerError(controller)
	case PermissionDeniedError:
		PermissionNotAllowError.ServerError(controller)
	default:
		ServerError.ServerError(controller)
	}
}
func CheckError(errorHandle func(error)) {
	troubleMaker := recover()
	if troubleMaker != nil {
		err := troubleMaker.(error)
		errorHandle(err)
	}
}
