package routers

import (
	"github.com/astaxie/beego"
	"yougame.com/yougame-server/controllers/api/cart"
	"yougame.com/yougame-server/controllers/api/collection"
	"yougame.com/yougame-server/controllers/api/comment"
	"yougame.com/yougame-server/controllers/api/game"
	"yougame.com/yougame-server/controllers/api/good"
	"yougame.com/yougame-server/controllers/api/image"
	"yougame.com/yougame-server/controllers/api/inventory"
	"yougame.com/yougame-server/controllers/api/order"
	"yougame.com/yougame-server/controllers/api/permission"
	"yougame.com/yougame-server/controllers/api/profile"
	"yougame.com/yougame-server/controllers/api/tag"
	"yougame.com/yougame-server/controllers/api/transaction"
	"yougame.com/yougame-server/controllers/api/user"
	"yougame.com/yougame-server/controllers/api/wallet"
	"yougame.com/yougame-server/controllers/api/wishlist"
)

func init() {
	beego.AddNamespace(beego.NewNamespace("/api",
		beego.NSRouter("users", &user.ApiUserController{}, "post:CreateUser;get:List"),
		beego.NSRouter("usergroups", &user.ApiUserController{}, "get:UserGroupList;post:CreateUserGroup"),
		beego.NSNamespace("usergroup",
			beego.NSNamespace("/:id",
				beego.NSRouter("/", &user.ApiUserController{}, "get:UserGroup"),
				beego.NSRouter("/permissions", &user.ApiUserController{}, "post:AddPermission;delete:RemovePermission"),
				beego.NSRouter("/users", &user.ApiUserController{}, "post:AddUserGroupUser;delete:RemoveUserGroupUser"),
			),
		),
		beego.NSNamespace("/user",
			beego.NSNamespace("/:id",
				beego.NSRouter("/", &user.ApiUserController{}, "get:GetUser"),
				beego.NSRouter("/orders", &order.ApiOrderController{}, "get:GetOrderList"),
				beego.NSRouter("/profile", &user.ApiUserController{}, "get:GetUserProfile"),
				beego.NSRouter("/inventory/game", &user.ApiUserController{}, "get:GetInventoryGame"),

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
				beego.NSNamespace("/wallet",
					beego.NSRouter("/", &wallet.Controller{}, "get:GetWallet"),
				),
				beego.NSNamespace("/transactions",
					beego.NSRouter("/", &transaction.Controller{}, "get:GetTransactionList"),
				),
			),
			beego.NSRouter("/auth", &user.ApiUserController{}, "post:UserLogin"),
			beego.NSRouter("/reset", &user.ApiUserController{}, "post:SendResetPasswordEmail"),
			beego.NSRouter("/password", &user.ApiUserController{}, "post:RecoveryPassword"),
		),
		beego.NSNamespace("/game",
			beego.NSNamespace("/:id",
				beego.NSRouter("/", &game.Controller{}, "get:GetGame;put:UpdateGame;patch:UpdateGame;delete:DeleteGame"),
				beego.NSRouter("/band", &game.Controller{}, "get:GetGameBand;put:UploadGameBand;post:UploadGameBand"),
				beego.NSRouter("/preview", &game.Controller{}, "get:GetGamePreview;post:UploadGamePreviewImage"),
				beego.NSRouter("/tags", &game.Controller{}, "get:GetTags;post:AddTags;delete:DeleteTags"),
				beego.NSRouter("/goods", &game.Controller{}, "post:AddGood;get:GetGood"),
				beego.NSNamespace("/comments",
					beego.NSRouter("/summary", &comment.ApiCommentController{}, "get:GetCommentSummary"),
				),
			),
		),
		beego.NSNamespace("tags",
			beego.NSRouter("/", &tag.Controller{}, "post:CreateTag;get:List"),
		), beego.NSNamespace("permissions",
			beego.NSRouter("/", &permission.Controller{}, "get:List"),
		),
		beego.NSNamespace("tag",
			beego.NSNamespace("/:id",
				beego.NSRouter("/", &tag.Controller{}, "patch:Update;put:Update;delete:DeleteTag"),
			),
		),
		beego.NSNamespace("/collections",
			beego.NSRouter("/", &collection.Controller{}, "get:GetGameCollectionList;post:Create;put:UpdateBulkCollection"),
		),
		beego.NSNamespace("/collection",
			beego.NSNamespace("/:id",
				beego.NSRouter("/", &collection.Controller{}, "delete:DeleteGameCollection;patch:Update;put:Update;get:GetObject"),
				beego.NSRouter("/games", &collection.Controller{}, "post:AddGame;delete:DeleteGame"),
			),
		),
		beego.NSNamespace("/games",
			beego.NSRouter("/", &game.Controller{}, "post:CreateGame;get:GetGameList"),
		),
		beego.NSNamespace("/profile",

			beego.NSRouter("/", &profile.Controller{}, "get:GetProfileList"),
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
			beego.NSRouter("/", &order.ApiOrderController{}, "get:GetOrderList;post:CreateOrder"),
			beego.NSNamespace("/:id",
				beego.NSRouter("/goods", &order.ApiOrderController{}, "get:GetOrderGoodsWithOrder"),
			),
		),
		beego.NSNamespace("/order",
			beego.NSNamespace("/:id",
				beego.NSRouter("/goods", &order.ApiOrderController{}, "get:GetOrderGoodsWithOrder"),
				beego.NSRouter("/", &order.ApiOrderController{}, "get:GetOrder"),
				beego.NSRouter("/pay", &order.ApiOrderController{}, "post:PayOrder"),
			),
		),
		beego.NSNamespace("good",
			beego.NSNamespace("/:id",
				beego.NSRouter("/", &good.Controller{}, "put:UpdateGood;patch:UpdateGood;get:GetGood;delete:DeleteGood"),
				beego.NSRouter("/comments", &comment.ApiCommentController{}, "post:CreateComment"),
			),
		),
		beego.NSNamespace("/inventors",
			beego.NSRouter("/", &inventory.Controller{}, "get:GetInventoryList"),
		),
		beego.NSNamespace("image",
			beego.NSNamespace("/:id",
				beego.NSRouter("/", &image.Controller{}, "put:UpdateImage;patch:UpdateImage;delete:DeleteImage"),
			),
		),
		beego.NSNamespace("images",
			beego.NSRouter("/", &image.Controller{}, "get:GetImageList"),
		),
		beego.NSNamespace("goods",
			beego.NSRouter("/", &good.Controller{}, "get:GetGoods;post:CreateGood;delete:DeleteBulkGood;put:UpdateBulkGood"),
		),
		beego.NSRouter("/ordergood", &order.ApiOrderController{}, "get:GetOrderGoods"),
		beego.NSRouter("/comments", &comment.ApiCommentController{}, "get:GetCommentList;post:CreateComment;delete:DeleteComments;put:UpdateComments"),
		beego.NSNamespace("/comment",
			beego.NSNamespace("/:id",
				beego.NSRouter("/", &comment.ApiCommentController{}, "put:Update;patch:Update")),
		),
	))
}
