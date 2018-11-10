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
            yield put({
                type: 'fetchUserSucceed',
                result: result
            })
        },
        * refreshUserProfile(action, {put, call, select}) {
            const { app } = yield select();
            if (app.user == null) {
                return
            }
            const result = yield call(FetchUser, app.user.id);
            console.log(result);
            yield put({
                type: 'onRefreshUserProfileSucceed',
                user: result
            })
        }
    },
    reducers: {
        'fetchUserSucceed'(state, {result: result}) {
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
        },
        onRefreshUserProfileSucceed(state, {user}) {
            let stateUser = state.user;
            stateUser.profile = user.profile;
            console.log(user);
            return {
                ...state,
                user
            }
        }

    },
};