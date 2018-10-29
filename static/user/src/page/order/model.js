import {ServerUrl} from "../../config/api";
import {FetchOrderList} from "../../service/order";

export default {
    namespace: 'orderpage',
    state: {
        filters: [{
            name: "未付款",
            active: false
        },
            {
                name: "已付款",
                active: false
            }],
        orders: []
    },
    subscriptions: {
        setup({dispatch}) {
            dispatch({type: 'fetchOrderList', payload: {uid: 3}})
        },
    },
    effects: {
        * fetchOrderList(action, {put, call}) {
            let uid = action.payload.uid;
            const result = yield call(FetchOrderList, uid);
            console.log(result)
            result.result.forEach(order => {
                order.goods.forEach(good => {
                    good.band_pic = `${ServerUrl}/${good.band_pic}`
                });
            });
            yield put({
                type: 'fetchOrderListSuccess',
                result
            })

        },
    },
    reducers: {
        'setFilter'(state, {name, active}) {
            let newFilter = state.filters.map(filter => {
                if (name === filter.name) {
                    return {
                        name: filter.name,
                        active: active
                    }
                }
                return {...filter}
            });
            return {
                ...state,
                filters: newFilter
            }
        },
        'fetchOrderListSuccess'(state, {result}) {
            return {
                ...state,
                orders: result.result
            }
        }
    },
};