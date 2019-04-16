<template>
  <div>
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
      allUsers: [],
      selectedUser: {}
    };
  },
  methods: {
    getUsers: function() {
      this.$http.get("/getUsers").then(response => {
        this.allUsers = [].concat.apply([], response.body);
        this.sortByName(this.allUsers);
      });
    },
    sortByName: function(array) {
      array.sort(function(a, b) {
        var bierNameA = a.BierName.toLowerCase(),
          bierNameB = b.BierName.toLowerCase();
        var firstNameA = a.FirstName.toLowerCase(),
          firstNameB = b.FirstName.toLowerCase();
        var lastNameA = a.LastName.toLowerCase(),
          lastNameB = b.LastName.toLowerCase();
        if (bierNameA < bierNameB) return -1;
        if (bierNameA > bierNameB) return 1;
        if (firstNameA < firstNameB) return -1;
        if (firstNameA > firstNameB) return 1;
        if (lastNameA < lastNameB) return -1;
        if (lastNameA > lastNameB) return 1;
        return 0;
      });
    },
    selectUser: function(user) {
      this.selectedUser = user;
      this.$emit("selectedUser", user);
    }
  },
  created() {
    this.$nextTick(this.getUsers());
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