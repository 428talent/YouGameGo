package routers

import (
	"github.com/astaxie/beego"
	"yougame.com/yougame-server/controllers/api/cart"
	"yougame.com/yougame-server/controllers/api/comment"
	"yougame.com/yougame-server/controllers/api/game"
	"yougame.com/yougame-server/controllers/api/good"
	"yougame.com/yougame-server/controllers/api/image"
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
				beego.NSRouter("/profile", &user.ApiUserController{}, "get:GetUserProfile"),

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
				beego.NSRouter("/", &game.GameController{}, "get:GetGame;put:UpdateGame;patch:UpdateGame;delete:DeleteGame"),
				beego.NSRouter("/band", &game.GameController{}, "get:GetGameBand;put:UploadGameBand;post:UploadGameBand"),
				beego.NSRouter("/preview", &game.GameController{}, "get:GetGamePreview;post:UploadGamePreviewImage"),
				beego.NSRouter("/tags", &game.GameController{}, "get:GetTags;post:AddTags"),
				beego.NSRouter("/goods", &game.GameController{}, "post:AddGood;get:GetGood"),
			),
		),
		beego.NSNamespace("/games",

			beego.NSRouter("/", &game.GameController{}, "post:CreateGame;get:GetGameList"),
		),
		beego.NSNamespace("wishlist",
			beego.NSRouter("/", &wishlist.ApiWishListController{}, "get:GetWishList;delete:DeleteWishListItems;post:Create"),
			beego.NSRouter("/:id", &wishlist.ApiWishListController{}, "delete:DeleteItem"),
		),
		beego.NSNamespace("carts",
			beego.NSRouter("/", &cart.ApiCartController{}, "post:Create;get:GetCartList"),
		),
		beego.NSNamespace("cart",
			beego.NSRouter("/:id", &cart.ApiCartController{}, "delete:DeleteItem"),
		),
		beego.NSNamespace("/orders",

			beego.NSRouter("/", &order.ApiOrderController{}, "get:GetOrderList"),
			beego.NSNamespace("/:id",
				beego.NSRouter("/goods", &order.ApiOrderController{}, "get:GetOrderGoodsWithOrder"),
			),
		),
		beego.NSNamespace("good",
			beego.NSNamespace("/:id",
				beego.NSRouter("/", &good.Controller{}, "put:UpdateGood;patch:UpdateGood;get:GetGood;delete:DeleteGood"),
				beego.NSRouter("/comments", &comment.ApiCommentController{}, "post:CreateComment"),
			),
		),
		beego.NSNamespace("image",
			beego.NSNamespace("/:id",
				beego.NSRouter("/", &image.Controller{}, "put:UpdateImage;patch:UpdateImage;delete:DeleteImage"),
			),
		),
		beego.NSNamespace("goods",
			beego.NSRouter("/", &good.Controller{}, "get:GetGoods;post:CreateGood"),
		),
		beego.NSRouter("/ordergood", &order.ApiOrderController{}, "get:GetOrderGoods"),
		beego.NSRouter("/comments", &comment.ApiCommentController{}, "get:GetCommentList"),
	))
}
