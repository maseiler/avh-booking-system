import { defineStore } from 'pinia'

export const useUserStore = defineStore('user', {
  state: () => ({
    users: [],
  }),
  getters: {
    getUsersAH: (state) => {
      return state.users.filter((user) => user['Status'] === 'AH')
    },
  },
  actions: {
    async fetchUsers() {
      this.users = ['ABC', 'DEF']
    },
  },
})
