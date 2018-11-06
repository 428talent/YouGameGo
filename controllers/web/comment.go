package web

import "github.com/astaxie/beego"

type CommentController struct {
	beego.Controller
}

func (c *CommentController)WriteComment(){
	c.TplName = "comment/write.html"
}

