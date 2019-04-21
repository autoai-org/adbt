import Vue from 'vue'
import Router from 'vue-router'
import Main from '@/pages/Main'
import Home from '@/pages/Home'
import SignIn from '@/pages/auth/SignIn'
Vue.use(Router)

export default new Router({
    routes: [
    
        {
            path: '/',
            component: Main,
            children:[
                {
                    path: '',
                    component: Home
                },
                {
                    path: 'signin',
                    component: SignIn
                }
            ]
        },
    ]
})
