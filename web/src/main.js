import { createApp } from 'vue'
import { createRouter, createWebHashHistory } from 'vue-router'

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import login from './components/UserLogin.vue'
import About from "./components/About.vue";
import App from './App.vue'

const routes = [
    { path: '/',  redirect: '/login' },
    { path: '/login', component: login },
    { path: '/about', component: About },
]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
})


const app = createApp(App)
app.use(router)
app.use(ElementPlus, { size: 'small', zIndex: 3000 })

app.mount('#app')

