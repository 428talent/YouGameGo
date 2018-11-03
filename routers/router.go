package routers

import (
	"github.com/astaxie/beego"
	cart2 "yougame.com/yougame-server/controllers/api/cart"
	"yougame.com/yougame-server/controllers/api/game"
	"yougame.com/yougame-server/controllers/api/order"
	"yougame.com/yougame-server/controllers/api/user"
	ApiWishlist "yougame.com/yougame-server/controllers/api/wishlist"
	"yougame.com/yougame-server/controllers/web"
	"yougame.com/yougame-server/controllers/web/admin"
	"yougame.com/yougame-server/controllers/web/cart"
	order2 "yougame.com/yougame-server/controllers/web/order"
	"yougame.com/yougame-server/controllers/web/wishlist"
)

func init() {
	beego.Router("/", &web.MainController{})
	beego.Router("/register", &web.RegisterController{})
	beego.Router("/login", &web.AuthorizationController{})
	beego.Router("/user", &web.AuthorizationController{}, "post:CreateUser")
	beego.Router("/signout", &web.AuthorizationController{}, "get:Logout")
	beego.Router("/login/auth", &web.AuthorizationController{}, "post:Login")
	beego.Router("/search/:key", &web.SearchController{})
	beego.Router("/game/:id", &web.DetailController{})
	beego.Router("/user/:id", &web.UserController{})
	beego.Router("/wishlist", &wishlist.WishListController{}, "post:SaveWishList")
	beego.Router("/cart", &cart.CartController{})
	beego.Router("/cart/:id/delete", &cart.CartController{}, "post:RemoveCartItem")
	beego.Router("/order/:id", &order2.OrderController{})
	beego.Router("/order", &order2.OrderController{}, "post:CreateOrder")
	beego.Router("/cart/delete", &cart.CartController{}, "post:ClearAll")
	beego.Router("/admin/dashboard", &admin.AdminDashboardController{})
	beego.Router("/api/game/:id/band", &game.GameController{}, "post:UploadGameBand")
	beego.Router("/api/game/:id", &game.GameController{}, "get:GetGame")
	beego.Router("/api/game/:id/preview/image", &game.GameController{}, "post:UploadGamePreviewImage")
	beego.Router("/api/game/:id/tags", &game.GameController{}, "post:AddTags")
	beego.Router("/api/game/:id/goods", &game.GameController{}, "post:AddGood")
	beego.Router("/api/user/:id/wishlist", &ApiWishlist.ApiWishListController{}, "get:GetWishList")
	beego.Router("/api/user/:id", &user.ApiUserController{}, "get:GetUser")
	beego.Router("/api/orders", &order.ApiOrderController{}, "post:CreateOrder")
	beego.Router("/api/order/:id/pay", &order.ApiOrderController{}, "post:PayOrder")
	beego.Router("/api/user/:id/orders", &order.ApiOrderController{}, "get:GetOrderList")
	beego.Router("/api/user/:id/carts", &cart2.ApiCartController{}, "get:GetCartList")
	beego.Router("/api/games", &game.GameController{})
	user.RegisterUserApiRouter()
}
