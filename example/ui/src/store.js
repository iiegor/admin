import Vue from 'vue'
import Vuex from 'vuex'
import * as api from '@/api'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    meta: null,
    user: null,
    loggedIn: false
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
    auth ({ commit }, { username, password }) {
      api.auth(username, password)
        .then(data => {
          console.log('data:', data)

          commit('updateUser', data)
        })
        .catch(err => {
          console.error('actions.auth:', err)

          commit('updateUser', null)
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
          console.error('actions.fetchMeta', err)

          commit('updateMeta', null)
        })
    }
  }
})
