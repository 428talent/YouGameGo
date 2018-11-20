package util

import (
	"fmt"
	"github.com/astaxie/beego"
)

func GetSiteAndPortUrl(c beego.Controller) string  {
	if c.Ctx.Input.Port() == 80{
		return c.Ctx.Input.Site()
	}else{
		return fmt.Sprintf("%s:%d",c.Ctx.Input.Site(),c.Ctx.Input.Port())
	}
}