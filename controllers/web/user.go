package web

import "github.com/astaxie/beego"

type UserController struct {
	beego.Controller
}

func (c *UserController)Get()  {
	c.TplName = "user/index.html"
}