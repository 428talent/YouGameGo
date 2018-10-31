const BaseAuthUrl = "http://localhost:8080";
export const ServerUrl = "http://localhost:8888";
export const Api = {
    getUser: ServerUrl + "/api/user/:id",
    getWishListItems: ServerUrl + "/api/user/:id/wishlist",
    getGame: ServerUrl + "/api/game/:id",
    getCartList: ServerUrl + "/api/user/:id/carts",
    getOrderList: ServerUrl + "/api/user/:id/orders",
    payOrder: ServerUrl + "/api/order/:id/pay"
};
