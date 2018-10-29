import {FetchWishListItems} from "../../service/wishlist";
import {FetchGame} from "../../service/game";
import {FetchCartList} from "../../service/cart";
import {ServerUrl} from "../../config/api";


export default {
    namespace: 'cartpage',
    state: {
        cartListItems: [],
        totalPrice:0
    },
    subscriptions: {
        setup({dispatch}) {
            dispatch({type: 'fetchCartlist', payload: {uid: 3}})
        },
    },
    effects: {
        * fetchCartlist(action, {put, call}) {
            let uid = action.payload.uid;
            const result = yield call(FetchCartList, uid);
            result.result.forEach(item =>{
               item.game.band = `${ServerUrl}/${item.game.band}`
            });
            const cartListItems = result.result;
            yield put({
                type: 'fetchCartListSucceed',
                cartListItems: cartListItems
            })

        },
    },
    reducers: {
        'fetchCartListSucceed'(state, {cartListItems}) {
            let totalPrice = 0
            cartListItems.forEach(item =>{
                totalPrice += item.good.price
            })
            return {
                ...state,
                cartListItems,
                totalPrice
            }
        }
    },
};