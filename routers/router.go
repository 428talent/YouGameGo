package routers

import (
	"github.com/astaxie/beego"
	"yougame.com/yougame-server/controllers/api/cart"
	"yougame.com/yougame-server/controllers/api/comment"
	"yougame.com/yougame-server/controllers/api/game"
	"yougame.com/yougame-server/controllers/api/order"
	"yougame.com/yougame-server/controllers/api/user"
	"yougame.com/yougame-server/controllers/api/wishlist"
)

func init() {
	beego.AddNamespace(beego.NewNamespace("/api",
		beego.NSRouter("users", &user.ApiUserController{}, "post:CreateUser"),
		beego.NSNamespace("/user",
			beego.NSNamespace("/:id",
				beego.NSRouter("/", &user.ApiUserController{}, "get:GetUser"),
				beego.NSRouter("/orders", &order.ApiOrderController{}, "get:GetOrderList"),

				beego.NSNamespace("/avatar",
					beego.NSRouter("/upload", &user.ApiUserController{}, "post:UploadAvatar"),
					beego.NSRouter("/", &user.ApiUserController{}, "put:UploadJsonAvatar"),
				),
				beego.NSNamespace("/profile",
					beego.NSRouter("/", &user.ApiUserController{}, "put:ChangeUserProfile"),
				),
				beego.NSNamespace("/carts",
					beego.NSRouter("/", &cart.ApiCartController{}, "get:GetCartList"),
				),
				beego.NSNamespace("/wishlist",
					beego.NSRouter("/", &user.ApiUserController{}, "get:GetUserWishList"),
				),
				beego.NSNamespace("/orders",
					beego.NSRouter("/", &user.ApiUserController{}, "get:GetOrderList"),
				),
			),
			beego.NSRouter("/auth", &user.ApiUserController{}, "post:UserLogin"),
		),
		beego.NSNamespace("/game",
			beego.NSNamespace("/:id",
				beego.NSRouter("/", &game.GameController{}, "get:GetGame"),
				beego.NSRouter("/band", &game.GameController{}, "get:GetGameBand;put:UploadGameBand"),
				beego.NSRouter("/preview", &game.GameController{}, "get:GetGamePreview"),
				beego.NSRouter("/tags", &game.GameController{}, "get:GetTags;post:AddTags"),
				beego.NSRouter("/goods", &game.GameController{}, "post:AddGood"),
			),
		),
		beego.NSNamespace("/games",

			beego.NSRouter("/", &game.GameController{}, "post:CreateGame"),

		),
		beego.NSNamespace("wishlist",
			beego.NSRouter("/", &wishlist.ApiWishListController{}, "get:GetWishList;delete:DeleteWishListItems"),
		),
		beego.NSNamespace("/orders",

			beego.NSRouter("/", &order.ApiOrderController{}, "get:GetOrderList"),
			beego.NSNamespace("/:id",
				beego.NSRouter("/goods", &order.ApiOrderController{}, "get:GetOrderGoodsWithOrder"),
			),
		),
		beego.NSNamespace("/good",
			beego.NSNamespace("/:id",
				beego.NSRouter("/comments", &comment.ApiCommentController{}, "post:CreateComment"),
				beego.NSRouter("/", &game.GameController{}, "get:GetGood"),
			),
		),
		beego.NSRouter("/ordergood", &order.ApiOrderController{}, "get:GetOrderGoods"),
		beego.NSRouter("/comments", &comment.ApiCommentController{}, "get:GetCommentList"),

	))
}
