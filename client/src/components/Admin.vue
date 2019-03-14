<template>
  <div>
    <h1>ADMIN</h1>
    <button class="button is-link" @click="showAddUserFormAdmin = true">Add User</button>
    <AddUserFormAdmin v-if="showAddUserFormAdmin" @close="showAddUserFormAdmin = false"></AddUserFormAdmin>
    <div class="columns">
      <div class="column">
        <div class="field">
          <div class="control has-icons-left">
            <input
              class="input"
              type="text"
              placeholder="Search user"
              v-model="search"
              v-on:keyup="searchUsers"
            >
            <span class="icon is-small is-left">
              <font-awesome-icon icon="search"/>
            </span>
          </div>
        </div>
      </div>
    </div>
    <div class="buttons" v-if="searchResults != []">
      <button
        class="button"
        v-for="user in searchResults"
        :key="user"
        @click="selectedUser = user"
        :class="[selectedUser === user ? 'is-link' : '']"
      >{{ displayName(user) }}</button>
    </div>
    <UserInfo :user="selectedUser"></UserInfo>
  </div>
</template>

<script src="./Admin.js"></script>