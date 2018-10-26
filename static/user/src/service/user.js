import fetch from 'dva/fetch';
import {Api} from "../config/api";

export async function FetchUser(userId) {
    const response = await fetch(Api.getUser.replace(":id",userId));
    const data = await response.json();
    return data
}