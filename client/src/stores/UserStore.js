import { defineStore } from 'pinia'
import axios from 'axios'
import { printHTTPError } from '@/helper'

export const useUserStore = defineStore('user', {
  state: () => ({
    users: [],
  }),
  getters: {
    getUsersAH: (state) => {
      console.log('getUsersAH')
      return state.users.filter((user) => user['Status'] === 'AH')
    },
  },
  actions: {
    fetchUsers() {
      let self = this
      axios
        .get('getUsers')
        .then(function (response) {
          self.users = response.data
          // users = users.filter(Boolean)
          // helper.sortUsersByName(users) // TODO
        })
        .catch(function (error) {
          printHTTPError(error)
        })
    },
  },
})
