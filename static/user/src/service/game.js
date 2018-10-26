import request from "../util/request";
import {Api} from "../config/api";

export async function FetchGame(gameId) {
    let response = await request({
        url: Api.getGame,
        method: 'get',
        pathParams: {
            id: gameId
        }
    });
    return await response.json()
}