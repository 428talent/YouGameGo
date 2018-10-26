import {FetchUser} from "../../../service/user";

export default {
    namespace: 'pagenav',
    state: {
        activeTab: "home"
    },
    subscriptions: {

    },
    effects: {

    },
    reducers: {
        'changePage'(state, {tag: tag}) {
            return {
                ...state,
                activeTab: tag
            }
        },
    },
};