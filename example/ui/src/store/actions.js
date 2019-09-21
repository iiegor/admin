import * as api from '../api'

export default {
  login ({ commit }, { username, password }) {
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
