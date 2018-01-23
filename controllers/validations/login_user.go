package validations

import "github.com/astaxie/beego"

func CheckValidate(c *beego.Controller) {
	c.GetString("username")

}