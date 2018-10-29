import request from "../util/request";
import {Api} from "../config/api";

export async function FetchOrderList(userId) {
    let response = await request({
        url: Api.getOrderList,
        method: 'get',
        pathParams: {
            id: userId
        }
    });
    return await response.json()
}