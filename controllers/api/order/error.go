package order

import (
	"net/http"
	ApiError "yougame.com/yougame-server/error"
)

var (
	NotSufficientFunds = ApiError.NewApiError(ApiError.APIError{
		Err:    "NotSufficientFunds",
		Detail: "Account balance not enough",
		Code:   "000004",
	}, http.StatusBadRequest)

	WrongOrderState = ApiError.NewApiError(ApiError.APIError{
		Err:    "WrongOrderState",
		Detail: "Wrong order state",
		Code:   "000005",
	}, http.StatusBadRequest)
)

