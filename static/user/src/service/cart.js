import request from "../util/request";
import {Api} from "../config/api";

export async function FetchCartList(userId) {
    let response = await request({
        url: Api.getCartList,
        method: 'get',
        pathParams: {
            id: userId
        }
    });
    return await response.json()
}