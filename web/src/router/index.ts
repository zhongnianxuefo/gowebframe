import { createRouter, createWebHashHistory, createWebHistory, RouteRecordRaw } from 'vue-router'
import request from "@/request";
import {useStore} from "vuex";

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    redirect: '/login'
  },
  {
    path: '/login',
    name: 'login',
    component: ()=>import('../views/LoginView.vue')

  },
  {
    path: '/home',
    name: 'home',
    component: ()=>import("../views/HomeView.vue"),
  },
  {
    path: '/list/:module/',
    name: 'list',
    component: ()=>import("../views/ListView.vue")
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
  }
]

const router = createRouter({
  // history: createWebHistory(process.env.BASE_URL),
  history: createWebHashHistory(),
  routes,
})

export default router



// import { createRouter, createWebHashHistory } from 'vue-router'


//
// const routes = [
//   {
//     path: '/',
//     redirect: '/login'
//   },
//   {
//     path: '/login',
//     name: 'login',
//     component: ()=>import("@/views/LoginPage")
//   },
//   {
//     path: '/home',
//     name: 'home',
//     component: ()=>import("@/views/HomePage")
//   },
//   {
//     path: '/list/:module/',
//     name: 'list',
//     component: ()=>import("@/views/ListPage")
//   },
//   {
//     path: '/test1/:mod/:id',
//     name: 'test1',
//     component: ()=>import("@/views/TestPage1")
//   },
// ]


