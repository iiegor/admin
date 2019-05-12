import Vue from 'vue'
import Router from 'vue-router'
import Dashboard from '@/pages/Dashboard'
import Auth from '@/pages/Auth'
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
      name: 'auth',
      path: '/auth',
      component: Auth
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
  if (!store.state.loggedIn && to.name !== 'auth') {
    return next('/auth')
  }

  next()
})

export default router
