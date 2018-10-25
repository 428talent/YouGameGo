package error

import (
	"fmt"
	"github.com/astaxie/beego"
)

type APIError struct {
	Err    string
	Detail string
	Code   string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API Error %s \n error:%s \n detail:%s", e.Code, e.Err, e.Code)
}

type APIErrorResponse struct {
	Success bool `json:"success"`
	Err     string `json:"error"`
	Detail  string `json:"detail"`
	Code    string `json:"code"`
}

func NewApiError(err APIError) *APIErrorResponse {
	return &APIErrorResponse{
		Success: false,
		Err:     err.Err,
		Detail:  err.Detail,
		Code:    err.Code,
	}
}

func (r *APIErrorResponse) ServerError(c beego.Controller,statusCode int) {
	c.Data["json"] = *r
	c.Ctx.ResponseWriter.WriteHeader(statusCode)
	c.ServeJSON()
}