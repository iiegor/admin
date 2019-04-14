import Vue from 'vue'
import Vuex from 'vuex'
import * as api from '@/api'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    meta: null,
    loggedIn: false,
    user: null
  },
  mutations: {
    updateMeta (state, data) {
      state.meta = data
    },

    updateUser (state, data) {
      state.user = data

      if (state.user) {
        state.loggedIn = true
      } else {
        state.loggedIn = false
      }
    }
  },
  actions: {
    login ({ commit }, { email, password }) {
      api.login(email, password)
        .then(data => {
          console.log('data:', data)

          commit('updateUser', data)
        })
        .catch(err => {
          console.error('error:', err)
        })
    },

    logout ({ commit }) {
      commit('updateUser', null)
    },

    fetchMeta ({ commit }) {
      api.fetchMeta()
        .then(data => {
          commit('updateMeta', data)
        })
        .catch(err => {
          console.error('fetchMeta', err)

          commit('updateMeta', null)
        })
    }
  }
})
