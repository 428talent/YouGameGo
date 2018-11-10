import {ChangeUserProfile} from "../../service/user";
import {message} from "antd"

export default {
    namespace: 'settingpage',
    state: {
        profile: {
            avatar: {
                dialog: {
                    isShow: false,
                    scale: 1
                }
            },
        }
    },
    subscriptions: {},
    effects: {
        * 'changeUserProfile'(action, {put, call, select, take}) {
            const {app} = yield select();
            if (app.user == null) {
                return
            }
            console.log(action);
            const result = yield call(ChangeUserProfile, {userId: app.user.id, nickname: action.payload.nickname});
            yield put({ type:'app/refreshUserProfile'});
            message.success("更改昵称成功",3)
        }
    },
    reducers: {},
};