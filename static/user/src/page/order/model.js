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
            }]
    },
    subscriptions: {},
    effects: {},
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
        }
    },
};