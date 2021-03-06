// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import Devices from '@/components/Devices'
import Capture from '@/components/Capture'

import VModal from 'vue-js-modal'
Vue.use(VModal)

import VueGoodTable from 'vue-good-table'
Vue.use(VueGoodTable)

import 'vue-good-table/dist/vue-good-table.css'

import VueResource from 'vue-resource'
Vue.use(VueResource)

import VueRouter from 'vue-router'
Vue.use(VueRouter)

Vue.config.productionTip = true

Vue.use(VModal, { dialog: true })

const router = new VueRouter({
  mode: 'history',
  base: __dirname,
  routes: [
    {
      path: '/devices',
      name: 'devices',
      component: Devices
    },
    {
      path: '/capture/:interfaceName/:duration/:filename',
      name: 'capture',
      component: Capture
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
