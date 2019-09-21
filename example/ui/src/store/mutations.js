export default {
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
}
