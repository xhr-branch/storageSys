import { createStore } from 'vuex'

export default createStore({
  state: {
    user: null,
    isAuthenticated: false
  },
  mutations: {
    setUser(state, user) {
      state.user = user
      state.isAuthenticated = !!user
    },
    logout(state) {
      state.user = null
      state.isAuthenticated = false
      localStorage.removeItem('user')
    }
  },
  actions: {
    login({ commit }, user) {
      commit('setUser', user)
      localStorage.setItem('user', JSON.stringify(user))
    },
    logout({ commit }) {
      commit('logout')
    },
    initializeAuth({ commit }) {
      const user = JSON.parse(localStorage.getItem('user'))
      if (user) {
        commit('setUser', user)
      }
    }
  },
  getters: {
    isAuthenticated: state => state.isAuthenticated,
    currentUser: state => state.user
  }
}) 