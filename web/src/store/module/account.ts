
export default {
    namespaced: true,
    state: {
        token:"",
        allMenus: [],
    },
    mutations: {
        setToken(state:any, value:any) {
            state.token = value
        },
        setAllMenus (state:any, allMenus:any) {
            state.allMenus = allMenus;
        },
    },
    getters: {
        getToken(state:any) {
            return state.token
        },
        getAllMenus(state:any) {
            return state.allMenus
        },

        // getAllMenus: async function (state:any) {
        //     // var that = this
        //     // var code = Store.fetchYqm();
        //     // let params = {
        //     //     inviteCode: code
        //     // }
        //     const response = await http.post(params,api.getCode)
        //     // var resJson = response.data;
        //     return state.allMenus
        //
        // }


    }
}

