package routers

import (
	"github.com/astaxie/beego"
	"yougame.com/yougame-server/controllers/api/game"
	"yougame.com/yougame-server/controllers/api/order"
	"yougame.com/yougame-server/controllers/api/user"
	"yougame.com/yougame-server/controllers/web"
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
	beego.Router("/wishlist", &web.WishListController{}, "post:SaveWishList")
	beego.Router("/game/:id/good/:goodId/comment/write", &web.CommentController{}, "get:WriteComment")
	beego.Router("/cart", &web.CartController{})
	beego.Router("/cart/:id/delete", &web.CartController{}, "post:RemoveCartItem")
	beego.Router("/order/:id", &web.OrderController{})
	beego.Router("/order", &web.OrderController{}, "post:CreateOrder")
	beego.Router("/order/pay", &web.OrderController{}, "post:PayOrder")
	beego.Router("/cart/delete", &web.CartController{}, "post:ClearAll")
	beego.Router("/admin/dashboard", &web.AdminDashboardController{})
	beego.Router("/comments/create", &web.CommentController{},"post:SaveComment")
	//beego.Router("/api/game/:id/band", &game.GameController{}, "post:UploadGameBand")
	//beego.Router("/api/game/:id", &game.GameController{}, "get:GetGame")
	//beego.Router("/api/game/:id/preview/image", &game.GameController{}, "post:UploadGamePreviewImage")
	//beego.Router("/api/game/:id/tags", &game.GameController{}, "post:AddTags")
	//beego.Router("/api/game/:id/goods", &game.GameController{}, "post:AddGood")
	//beego.Router("/api/user/:id/wishlist", &ApiWishlist.ApiWishListController{}, "get:GetWishList")
	//beego.Router("/api/user/:id", &user.ApiUserController{}, "get:GetUser")
	//beego.Router("/api/orders", &order.ApiOrderController{}, "post:CreateOrder")
	//beego.Router("/api/order/:id/pay", &order.ApiOrderController{}, "post:PayOrder")
	//beego.Router("/api/user/:id/orders", &order.ApiOrderController{}, "get:GetOrderList")
	//beego.Router("/api/user/:id/carts", &cart2.ApiCartController{}, "get:GetCartList")
	beego.Router("/api/games", &game.GameController{})
	registerApiRouter()
	//user.RegisterUserApiRouter()
}

func registerApiRouter() {
	ns := beego.NewNamespace("/api",
		beego.NSRouter("users", &user.ApiUserController{}, "post:CreateUser"),
		beego.NSNamespace("/user",
			beego.NSNamespace("/:id",
				beego.NSRouter("/orders", &order.ApiOrderController{}, "get:GetOrderList"),
			),
			beego.NSRouter("/auth", &user.ApiUserController{}, "post:UserLogin"),
		),

	)
	beego.AddNamespace(ns)
}
