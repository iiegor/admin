import Vue from 'vue'
import Router from 'vue-router'
import Dashboard from '@/pages/Dashboard'
import Login from '@/pages/Login'
import Resource from '@/pages/Resource'
import Metrics from '@/pages/Metrics'
import store from './store'

Vue.use(Router)

const router = new Router({
  base: '/ui',
  mode: 'history',
  routes: [
    {
      path: '/resource/:name',
      component: Resource
    },
    {
      path: '/metrics',
      component: Metrics
    },
    {
      // Para identificar fácilmente
      // la ruta en los guards
      name: 'login',
      path: '/login',
      component: Login
    },
    {
      path: '/',
      component: Dashboard
    },
    {
      path: '*',
      redirect: '/'
    }
  ]
})

router.beforeEach((to, from, next) => {
  // Redireccionar al login
  // cuando la ruta sea otra y no esté logeado.
  if (!store.state.loggedIn && to.name !== 'login') {
    return next('/login')
  }

  next()
})

export default router
