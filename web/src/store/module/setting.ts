
export default {
    namespaced: true,
    state: {
        collapsed: false,
    },
    mutations: {
        setCollapsed (state:any, collapsed:any) {
            state.collapsed = collapsed;
        },
    },
    getters: {
        getCollapsed(state:any) {
            return state.collapsed
        },
    }
}

