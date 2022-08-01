import { createRouter, createWebHashHistory } from 'vue-router'
import login from "@/views/LoginPage";
import home from "@/views/HomePage";


const routes = [
    { path: '/',  redirect: '/login' },
    { path: '/login', component: login },
    { path: '/home', component: home },
]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
})

export default router
