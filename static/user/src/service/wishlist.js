import fetch from "dva/fetch";
import {Api} from "../config/api";
import request from "../util/request";

export async function FetchWishListItems(userId) {
    let response = await request({
        url: Api.getWishListItems,
        method: 'get',
        queryParams: {
            page: 1,
            pageSize: 10
        },
        pathParams: {
            id: userId
        }
    });
    return await response.json()
}