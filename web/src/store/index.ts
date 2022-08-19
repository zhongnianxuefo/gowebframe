import { createStore } from 'vuex'
import VuexPersistence from 'vuex-persist'

import account from './module/account'
import setting from './module/setting'

const vuexLocal = new VuexPersistence({
    storage: window.localStorage,
})

export default createStore({
    modules: {
        account,
        setting,
    },
    plugins: [vuexLocal.plugin]
})


