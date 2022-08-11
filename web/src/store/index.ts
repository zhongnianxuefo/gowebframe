import { createStore } from 'vuex'

export default createStore({
  state: {
    user: undefined,
    token: "",
    allMenus: [{title:"首页", icon:"home-outlined", page:"home", menuKey:"home", hasSubmenu:false}],
    collapsed: false,
  },
  getters: {
  },
  mutations: {
    setToken(state, token) {
      state.token = token;
    },
    setAllMenus (state, allMenus) {
      state.allMenus = allMenus;
      // localStorage.setItem(process.env.VUE_APP_USER_KEY, JSON.stringify(user))
    },
    setCollapsed(state, collapsed){
      state.collapsed = collapsed;
    }
  },
  actions: {

  },
  modules: {
  }
})
