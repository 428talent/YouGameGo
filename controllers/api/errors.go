package api

import (
	"errors"
	"github.com/astaxie/beego"
	"net/http"
	ApiError "yougame.com/yougame-server/error"
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
	default:
		ServerError.ServerError(controller)
	}
}
