<template>
  <div>
    <h1>USER SETTINGS</h1>
    <div class="columns">
      <div class="column is-4">
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
      </div>
      <div class="column is-narrow">
        <button class="button is-link" @click="showAddUserFormAdmin = true">Add User</button>
        <button class="button is-link" @click="showModifyUserForm = true">Modify User</button>
        <button class="button is-link" @click="showDeleteUserForm = true">Delete User</button>
        <AddUserFormAdmin v-if="showAddUserFormAdmin" @close="showAddUserFormAdmin = false"></AddUserFormAdmin>
        <ModifyUserForm
          :user="selectedUser"
          v-if="showModifyUserForm"
          @close="showModifyUserForm = false"
        ></ModifyUserForm>
        <DeleteUserForm
          :user="selectedUser"
          v-if="showDeleteUserForm"
          @close="showDeleteUserForm = false"
        ></DeleteUserForm>
        <UserInfo :user="selectedUser"></UserInfo>
      </div>
    </div>
  </div>
</template>

<script src="./UserSettings.js"></script>