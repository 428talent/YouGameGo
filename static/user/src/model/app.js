import {FetchUser} from "../service/user";

export default {
    namespace: 'app',
    state: {
        user: null
    },
    subscriptions: {
        setup({dispatch}) {
            dispatch({type: 'fetchUser', payload: {uid: 2}})
        },
    },
    effects: {
        * fetchUser(action, {put, call}) {
            let uid = action.payload.uid;
            const result = yield call(FetchUser, uid);
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
    },
};