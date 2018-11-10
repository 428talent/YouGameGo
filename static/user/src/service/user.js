import fetch from 'dva/fetch';
import {Api} from "../config/api";
import request from "../util/request";

export async function FetchUser(userId) {
    const response = await fetch(Api.getUser.replace(":id", userId));
    const data = await response.json();
    return data
}

export async function ChangeUserProfile({nickname, userId}) {
    let response = await request({
        url: Api.changeProfile,
        method: 'put',
        pathParams: {
            id: userId
        },
        data: {
            nickname
        }
    });
    return await response.json()
}