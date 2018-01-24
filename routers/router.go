package routers

import (
	"github.com/astaxie/beego"
	"you_game_go/controllers"
	"you_game_go/controllers/api/admin/game"
	"you_game_go/controllers/api/admin/user"
	"you_game_go/controllers/api/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/api/web/user/create", &api_web.CreateUserController{})
	beego.Router("/api/web/user/login", &api_web.UserLoginController{})

	beego.Router("/api/admin/user/login", &api_admin_user.UserLoginController{})
	beego.Router("/api/admin/game", &api_admin_game.GameController{})
}
