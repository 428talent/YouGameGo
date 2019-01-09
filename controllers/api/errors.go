package api

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	ApiError "yougame.com/yougame-server/application"
	"yougame.com/yougame-server/security"
	"yougame.com/yougame-server/service"
	"yougame.com/yougame-server/validate"
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

	ResourceNoFoundError = ApiError.NewApiError(ApiError.APIError{
		Err:    "ResourceNoFoundError",
		Detail: "resource  not found",
		Code:   "100004",
	}, http.StatusNotFound)

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

	DuplicateResourceApiError = ApiError.NewApiError(ApiError.APIError{
		Err:    "DuplicateResourceApiError",
		Detail: "resource already exist",
		Code:   "100005",
	}, http.StatusConflict)
)
var (
	ParseJsonDataError     = errors.New("cannot parse json request")
	ClaimsNoFoundError     = errors.New("claims not found")
	ResourceNotFoundError  = errors.New("resource not found")
	DuplicateResourceError = errors.New("resource already exist")
)

func HandleApiError(controller beego.Controller, err error) {

	switch err {
	case service.NoAuthError:
		AuthFailedError.ServerError(controller)
		return
	case ParseJsonDataError:
		ParseRequestDataError.ServerError(controller)
		return
	case security.ReadAuthorizationFailed:
		AuthFailedError.ServerError(controller)
		return
	case PermissionDeniedError:
		PermissionNotAllowError.ServerError(controller)
		return
	case orm.ErrNoRows:
		ResourceNoFoundError.ServerError(controller)
		return
	case ResourceNotFoundError:
		ResourceNoFoundError.ServerError(controller)
		return
	case DuplicateResourceError:
		DuplicateResourceApiError.ServerError(controller)
		return
	}

	if _, isJWTValidateError := err.(*jwt.ValidationError); isJWTValidateError {
		AuthFailedError.ServerError(controller)
		return
	}
	if validateError, ok := err.(*validate.ValidateError); ok {
		ApiError.NewApiError(ApiError.APIError{
			Err:    "InvalidateRequestDataError",
			Detail: validateError.Error(),
			Code:   "100005",
		}, http.StatusBadRequest).ServerError(controller)
		return
	}
	ServerError.ServerError(controller)
}
func CheckError(errorHandle func(error)) {
	troubleMaker := recover()
	if troubleMaker != nil {
		err := troubleMaker.(error)
		errorHandle(err)
	}
}
