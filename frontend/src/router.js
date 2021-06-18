import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

import Welcome from './components/Welcome'
import Home from './components/Home'
import ChatRoom from './components/ChatRoom'
import About from './components/About'

export const router = new VueRouter({
    mode: 'history',
    routes: [
        {
            path: '/welcome',
            component: Welcome,
            name: 'welcome'
        },
        {
            path: '/',
            component: Home,
            name: 'home',
            children: [
                {
                    path: 'chat',
                    component: ChatRoom,
                    name: 'chatroom'
                },
                {
                    path: 'about',
                    component: About,
                    name: 'about',
                },
            ],
            redirect: '/chat',
        },

        {
            path: '*',
            redirect: '/welcome'
        }
    ]
})

router.beforeEach((to, from, next) => {
    if (localStorage.getItem('token') == null && from.name != '/welcome') {
        next('/welcome')
        return
    }

    next()
})