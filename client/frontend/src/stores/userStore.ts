import { defineStore } from 'pinia'
import User, { nullUser } from '../models/User'
import { sortUserList } from '../helper'

export const useUserStore = defineStore('userStore', {
    state: () => ({
        userList: [] as User[],
        selectedUser: nullUser as User
    }),
    getters: {
        fetchAllUsers(): void {
            window.go.main.App.GetAllUsers().then((response: User[]) => {
                this.userList = sortUserList(response)
            })
        },
        getAllUsers(): User[] {
            return this.userList
        },
        getUsersAH(): User[] {
            return this.userList.filter((user: User) => user.Status === "AH")
        },
        getUsersAktivB(): User[] {
            return this.userList.filter((user: User) => user.Status === "Aktiv B")
        },
        getUsersAktivKA(): User[] {
            return this.userList.filter((user: User) => user.Status === "Aktiv KA")
        },
        getUsersAktivSteganleger(): User[] {
            return this.userList.filter((user: User) => user.Status === "Steganleger")
        },
        getUsersGast(): User[] {
            return this.userList.filter((user: User) => user.Status === "Gast")
        },
        getSelectedUser(): User {
            return this.selectedUser
        }
    },
    actions: {
        selectUser(user: User) {
            this.selectedUser = user
        },
    },
})
