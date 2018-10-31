import {ServerUrl} from "../../config/api";
import {FetchOrderList, PayOrder} from "../../service/order";

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
        orders: [],
        orderModal: {
            isShow: false,
            order: {},
            isPaying: false
        }
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
            console.log(result);
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
        * payOrder(action, {put, call, select}) {
            let orderId = action.payload.orderId;
            yield put({
                type: 'setPaying',
                isPaying: true
            });
            const result = yield call(PayOrder, orderId);
            if (result.success) {
                yield put({
                    type: 'app/setLoadingModalShow',
                    isShow: false
                });
            }
            yield put({
                type: 'setPaying',
                isPaying: false
            });


        }
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
        'setOrderModel'(state, {isShow, order}) {
            return {
                ...state,
                orderModal: {
                    isShow, order
                }
            }
        },
        'fetchOrderListSuccess'(state, {result}) {
            return {
                ...state,
                orders: result.result
            }
        },
        'setPaying'(state, {isPaying}) {
            return {
                ...state,
                orderModal: {
                    isPaying,
                    ...state.orderModal
                }
            }
        },
        'onPaySucceed'(state, {orderId}) {
            state.orders.forEach(order => {
                order.state = "Done"
            });
            return {
                ...state,
                orders: state.orders
            }
        }
    },
};