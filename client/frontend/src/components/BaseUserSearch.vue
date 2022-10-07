<template>
    <div>
        <!-- search bar -->
        <div class="columns">
            <div class="column">
                <div class="field">
                    <div class="control has-icons-left">
                        <input class="input" type="text" placeholder="Search user" v-model="search"
                            v-on:keyup="searchUsers" />
                        <span class="icon is-small is-left">
                            <font-awesome-icon icon="search" />
                        </span>
                    </div>
                </div>
            </div>
        </div>
        <!-- search results -->
        <!--<div class="buttons" v-if="searchResults !== null">-->
        <div class="buttons" v-if="searchResults.length > 0">
            <button class="button" v-for="user in searchResults" :key="user" :class="buttonColor(selectedUser, user)"
                @click="selectUser(user)">
                {{ displayUserName(user) }}
            </button>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { displayUserName, buttonColor, selectUser } from '../helper'
import User from './../models/User'
import { useUserStore } from '../stores/userStore'

const userStore = useUserStore()

export default defineComponent({
    data() {
        return {
            search: "",
            searchResults: [] as User[]
        }
    },
    computed: {
        allUsers(): User[] {
            return userStore.getAllUsers
        },
        selectedUser(): User {
            return userStore.getSelectedUser
        },
    },
    methods: {
        displayUserName,
        buttonColor,
        selectUser,
        searchUsers(): void {
            if (this.search != "") {
                var tmpSearch = this.search.toLowerCase()
                this.searchResults = this.allUsers.filter(
                    (user: User) =>
                        user.BierName.toLowerCase().includes(tmpSearch) ||
                        user.FirstName.toLowerCase().includes(tmpSearch) ||
                        user.LastName.toLowerCase().includes(tmpSearch) ||
                        user.Email.toLowerCase().includes(tmpSearch) ||
                        user.Phone.toLowerCase().includes(tmpSearch) ||
                        user.BoatName.toLowerCase().includes(tmpSearch)
                )
            } else {
                this.searchResults = [] as User[]
            }
        }
    }
})
</script>