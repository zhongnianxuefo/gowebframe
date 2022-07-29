import { createRouter, createWebHashHistory } from 'vue-router'
import login from "@/components/LoginPage";
import home from "@/components/HomePage";


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
