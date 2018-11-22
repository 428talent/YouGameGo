package util

import (
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

func GetSiteAndPortUrl(c beego.Controller) string  {
	if c.Ctx.Input.Port() == 80{
		return c.Ctx.Input.Site()
	}else{
		return fmt.Sprintf("%s:%d",c.Ctx.Input.Site(),c.Ctx.Input.Port())
	}
}

func BuildParamString(params map[string]string) string{
	rawParamList := make([]string,0)
	for name, value := range params {
		rawParamList = append(rawParamList, fmt.Sprintf("%s=%s",name,value))
	}
	return strings.Join(rawParamList,"&")
}