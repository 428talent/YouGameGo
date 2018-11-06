package web

import "github.com/astaxie/beego"

type AdminDashboardController struct {
	beego.Controller
}

func (c *AdminDashboardController) Get() {

	c.TplName = "admin/dashboard.html"
}