package routers

import (
	"you_game_go/controllers"
	"github.com/astaxie/beego"
	"you_game_go/controllers/api/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/api/web/user/create", &api_web.CreateUserController{})
}
