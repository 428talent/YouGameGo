package validate

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"net/http"
	"strings"
	AppError "yougame.com/yougame-server/error"
)

type ValidateError struct {
	Errors []*validation.Error
}

func (e ValidateError) Error() string {
	var errorStringList []string
	for _, validateError := range e.Errors {
		errorStringList = append(errorStringList, fmt.Sprintf("[%s : %s]", validateError.Key, validateError.Message))
	}
	return strings.Join(errorStringList, "   ")
}

func (e *ValidateError) BuildResponse() *AppError.APIErrorResponse {
	return &AppError.APIErrorResponse{
		Success: false,
		Err:     "Validate error",
		Detail:  e.Error(),
		Code:    AppError.ValidateError,
		StatusCode:http.StatusBadRequest,
	}
}
func ValidateData(r interface{}) error {
	valid := validation.Validation{}
	_, err := valid.Valid(r)
	if err != nil {
		return err
	}
	valid.HasErrors()
	beego.Debug(valid.HasErrors())
	if !valid.HasErrors() {
		return nil
	} else {
		return &ValidateError{Errors: valid.Errors}
	}
}
