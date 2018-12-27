package application

import (
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
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
	Success    bool   `json:"success"`
	Err        string `json:"error"`
	Detail     string `json:"detail"`
	Code       string `json:"code"`
	StatusCode int    `json:"-"`
}

func NewApiError(err APIError, statusCode int) *APIErrorResponse {
	return &APIErrorResponse{
		Success:    false,
		Err:        err.Err,
		Detail:     err.Detail,
		Code:       err.Code,
		StatusCode: statusCode,
	}
}

func (r *APIErrorResponse) ServerError(c beego.Controller) {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.Ctx.ResponseWriter.WriteHeader(r.StatusCode)
	c.Data["json"] = *r
	c.ServeJSON()
}

func ServerNoAuthError(c beego.Controller) {
	c.Data["json"] = &APIErrorResponse{
		Success: false,
		Err:     "Authorization failed",
		Detail:  "Authorization failed",
		Code:    "000003",
	}
	c.Ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
	c.ServeJSON()
}
