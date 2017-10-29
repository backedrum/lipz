// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import Home from '@/components/Home'
import Login from '@/components/Login'
import Signup from '@/components/Signup'
import SecretQuote from '@/components/SecretQuote'
import UserInfo from '@/components/UserInfo'
import Devices from '@/components/Devices'
import Devices2 from '@/components/Devices2'

import VueGoodTable from 'vue-good-table'
Vue.use(VueGoodTable)

// import css
import 'vue-easytable/libs/themes-base/index.css'

// import table and pagination comp
import {VTable, VPagination} from 'vue-easytable'

// Register to global
Vue.component(VTable.name, VTable)
Vue.component(VPagination.name, VPagination)

import VueResource from 'vue-resource'
Vue.use(VueResource)

import VueRouter from 'vue-router'
Vue.use(VueRouter)

Vue.config.productionTip = true

import auth from './auth'

function requireAuth (to, from, next) {
  if (!auth.isAuthenticated()) {
    this.$router.replace('/login')
  } else {
    next()
  }
}

const router = new VueRouter({
  mode: 'history',
  // base: __dirname,
  routes: [
    {
      path: '/',
      component: Home
    },
    {
      path: '/home',
      name: 'home',
      component: Home
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/signup',
      name: 'signup',
      component: Signup
    },
    {
      path: '/secretquote',
      name: 'secretquote',
      component: SecretQuote,
      beforeEnter: requireAuth
    },
    {
      path: '/userinfo',
      name: 'userinfo',
      component: UserInfo,
      beforeEnter: requireAuth
    },
    {
      path: '/devices',
      name: 'devices',
      component: Devices
    },
    {
      path: '/devices2',
      name: 'devices2',
      component: Devices2
    }
  ]
})

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: { App }
})
