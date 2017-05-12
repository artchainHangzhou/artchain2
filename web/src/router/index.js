import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

import Hello from 'components/Hello'
import Index from 'views/index'
import ifcont from 'views/userconten'
import tradefrom from 'views/tradefrom'
import entery from 'views/entery'

export default new Router({
  routes: [
    {
      path: '/',
      redirect:'/index',
    },
    {
      path: '/index',
      component: Index
    },
    {
      path: '/user',
      component: ifcont
    },
    {
      path: '/tradefrom',
      component: tradefrom
    },
    {
      path: '/entery',
      component: entery
    }
  ]
})
