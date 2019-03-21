<template>
  <div>
    <p class="title">Users</p>
    <div class="columns">
      <div class="column">
        <UserSearch :allUsers="allUsers" @selectUser="selectUser"/>
      </div>
      <div class="column is-2">
        <button class="button is-link" @click="showAddUserForm = true">Add User</button>
      </div>
      <AddUserForm v-if="showAddUserForm" @close="showAddUserForm = false"/>
    </div>
    <hr>
    <UserList :allUsers="allUsers" @selectUser="selectUser"/>
  </div>
</template>

<script src="./UserTab.js"></script>