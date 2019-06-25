<template>
  <div>
    <div class="columns">
      <div class="column">
        <UserSearch @selectUser="selectUser"/>
      </div>
      <div class="column is-2">
        <button class="button is-link" @click="showAddUserForm = true">Add User</button>
      </div>
      <AddUserForm v-if="showAddUserForm" @close="showAddUserForm = false"/>
    </div>
    <hr>
    <UserList @selectUser="selectUser"/>
  </div>
</template>

<script>
import UserSearch from "./UserSearch.vue";
import UserList from "./UserList.vue";
import AddUserForm from "./AddUserForm.vue";
import Vue from "vue";

export default {
  components: {
    UserSearch,
    UserList,
    AddUserForm
  },
  data: function() {
    return {
      showAddUserForm: false,
      selectedUser: {}
    };
  },
  methods: {
    selectUser: function(user) {
      this.selectedUser = user;
      this.$emit("selectedUser", user);
    }
  }
};

var UserEventBus = new Vue();
Object.defineProperties(Vue.prototype, {
  $userEventBus: {
    get: function() {
      return UserEventBus;
    }
  }
});
</script>