import {FetchUser} from "../service/user";

export default {
    namespace: 'app',
    state: {
        user: null,
        activeTab: "Home",
        isLoadingModalShow: false
    },
    subscriptions: {
        setup({dispatch}) {
            dispatch({type: 'fetchUser', payload: {uid: 3}})
        },
    },
    effects: {
        * fetchUser(action, {put, call}) {
            let uid = action.payload.uid;
            const result = yield call(FetchUser, uid);
            console.log("fetched user");
            console.log(result);
            yield put({
                type: 'fetchUserSucceed',
                result: result
            })
        }
    },
    reducers: {
        'fetchUserSucceed'(state, {result: result}) {
            console.log({
                ...state,
                user: result
            });
            return {
                ...state,
                user: result
            }
        },
        'changeTab'(state, {activeTab}) {
            return {
                ...state,
                activeTab: activeTab
            }
        },
        'setLoadingModalShow'(state, {isShow}) {
            return {
                ...state,
                isLoadingModalShow: isShow
            }
        }
    },
};