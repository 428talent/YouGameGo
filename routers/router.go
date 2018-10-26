package routers

import (
	"github.com/astaxie/beego"
	"yougame.com/yougame-server/controllers/api/admin/game"
	"yougame.com/yougame-server/controllers/api/admin/order"
	ApiWishlist "yougame.com/yougame-server/controllers/api/admin/wishlist"
	"yougame.com/yougame-server/controllers/api/web"
	"yougame.com/yougame-server/controllers/web"
	"yougame.com/yougame-server/controllers/web/admin"
	"yougame.com/yougame-server/controllers/web/cart"
	"yougame.com/yougame-server/controllers/web/wishlist"
)

func init() {
	beego.Router("/", &web.MainController{})
	beego.Router("/register", &web.RegisterController{})
	beego.Router("/login", &web.LoginController{})
	beego.Router("/signout", &web.LoginController{}, "get:Logout")
	beego.Router("/search/:key", &web.SearchController{})
	beego.Router("/game/:id", &web.DetailController{})
	beego.Router("/user/:id", &web.UserController{})
	beego.Router("/wishlist", &wishlist.WishListController{}, "post:SaveWishList")
	beego.Router("/cart", &cart.CartController{})
	beego.Router("/cart/:id/delete", &cart.CartController{}, "post:RemoveCartItem")
	beego.Router("/admin/dashboard", &admin.AdminDashboardController{})
	beego.Router("/api/web/user/create", &api_web.CreateUserController{})
	beego.Router("/api/web/user/login", &api_web.UserLoginController{})
	beego.Router("/api/game/:id/band", &api_admin_game.GameController{}, "post:UploadGameBand")
	beego.Router("/api/game/:id", &api_admin_game.GameController{}, "get:GetGame")
	beego.Router("/api/game/:id/preview/image", &api_admin_game.GameController{}, "post:UploadGamePreviewImage")
	beego.Router("/api/game/:id/tags", &api_admin_game.GameController{}, "post:AddTags")
	beego.Router("/api/game/:id/goods", &api_admin_game.GameController{}, "post:AddGood")
	beego.Router("/api/user/:id/wishlist", &ApiWishlist.ApiWishListController{}, "get:GetWishList")
	beego.Router("/api/orders", &order.ApiOrderController{}, "post:CreateOrder")
	beego.Router("/api/admin/game", &api_admin_game.GameController{})
}
