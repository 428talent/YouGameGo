import {FetchWishListItems} from "../../service/wishlist";
import {FetchGame} from "../../service/game";

export default {
    namespace: 'homepage',
    state: {
        wishlist: [],
        wishGameList:[]
    },
    subscriptions: {
        setup({dispatch}) {
            dispatch({type: 'fetchWishlist', payload: {uid: 3}})
        },
    },
    effects: {
        * fetchWishlist(action, {put, call}) {
            let uid = action.payload.uid;
            const result = yield call(FetchWishListItems, uid);
            const wishlistItems = result.result;
            for (let idx = 0; idx < wishlistItems.length; idx++) {
                const gameResult = yield call(FetchGame, wishlistItems[idx].game_id);
                yield put({
                    type: 'fetchWishlistGame',
                    game: gameResult
                })
            }
        },

        * getGame(action, {put, call}) {
            let gameId = action.payload.gameId;
        }
    },
    reducers: {
        'fetchWishlistSucceed'(state, {result: result}) {
            return {
                ...state,
                wishlist: result
            }
        },
        'fetchWishlistGame'(state, {game}) {
            console.log(state);
            return {
                ...state,
                wishlist: [game,...state.wishlist],
                wishGameList: [game,...state.wishGameList]
            }
        }
    },
};