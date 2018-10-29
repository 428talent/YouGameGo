const BaseAuthUrl = "http://localhost:8080";
export const ServerUrl = "http://localhost:8888";
export const Api = {
    getUser: BaseAuthUrl + "/v1/user/:id",
    getWishListItems: ServerUrl + "/api/user/:id/wishlist",
    getGame: ServerUrl + "/api/game/:id",
    getCartList: ServerUrl + "/api/user/:id/cart"
};
