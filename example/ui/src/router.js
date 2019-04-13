import Vue from 'vue'
import Router from 'vue-router'
import Dashboard from '@/pages/Dashboard'
import Login from '@/pages/Login'
import ModelView from '@/pages/ModelView'
import Metrics from '@/pages/Metrics'

Vue.use(Router)

const router = new Router({
  base: '/ui',
  mode: 'history',
  routes: [
    {
      path: '/',
      component: Dashboard
    },
    {
      path: '/resource/:name',
      component: ModelView,
      props: (route) => ({ query: route.query.q })
    },
    {
      path: '/metrics',
      component: Metrics
    },
    {
      path: '/login',
      component: Login
    },
    {
      path: '*',
      redirect: '/'
    }
  ]
})

router.beforeEach((to, from, next) => {
  // ..
  next()
})

export default router
