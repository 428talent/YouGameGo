package comment

import(
	"net/http"
	ApiError "yougame.com/yougame-server/error"
)

var (
	CommentAlreadyExistError = ApiError.NewApiError(ApiError.APIError{
		Err:    "CommentAlreadyExistError",
		Detail: "comment already exist !",
		Code:   "110001",
	}, http.StatusBadRequest)

)