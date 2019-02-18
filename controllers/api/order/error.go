package order

import (
	"net/http"
	ApiError "yougame.com/yougame-server/application"
	"yougame.com/yougame-server/controllers/api"
	"yougame.com/yougame-server/service"
)

func init() {
	api.RegisterErrors(map[error]*ApiError.APIErrorResponse{
		service.WrongOrderStateError:    WrongOrderState,
		service.NotSufficientFundsError: NotSufficientFunds,
	})
}

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
