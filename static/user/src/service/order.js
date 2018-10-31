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

export async function PayOrder(orderId) {
    let response = await request({
        url: Api.payOrder,
        method: 'post',
        pathParams: {
            id: orderId
        }
    });
    return await response.json()
}