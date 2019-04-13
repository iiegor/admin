import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    counter: 0,
    meta: null
  },
  mutations: {
    increment (state) {
      state.counter++
    },

    meta (state, data) {
      state.meta = data
    }
  },
  actions: {
    increment ({ commit }) {
      commit('increment')
    }
  }
})
