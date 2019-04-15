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
          >{{ displayUserName(user) }}</button>
        </div>
      </div>
      <div class="column is-narrow">
        <button class="button is-link" @click="showAddUserFormAdmin = true">Add User</button>
        <button class="button is-link" @click="showModifyUserForm = true">Modify User</button>
        <button class="button is-link" @click="showDeleteUserForm = true">Delete User</button>
        <AddUserFormAdmin v-if="showAddUserFormAdmin" @close="showAddUserFormAdmin = false"/>
        <ModifyUserForm
          :user="selectedUser"
          v-if="showModifyUserForm"
          @close="showModifyUserForm = false"
        />
        <DeleteUserForm
          :user="selectedUser"
          v-if="showDeleteUserForm"
          @close="showDeleteUserForm = false"
        />
        <UserInfo :user="selectedUser"/>
      </div>
    </div>
  </div>
</template>

<script>
import AddUserFormAdmin from "./AddUserFormAdmin.vue";
import ModifyUserForm from "./ModifyUserForm.vue";
import DeleteUserForm from "./DeleteUserForm.vue";
import UserInfo from "./UserInfo.vue";

export default {
  components: {
    AddUserFormAdmin,
    ModifyUserForm,
    DeleteUserForm,
    UserInfo
  },
  props: {
    users: []
  },
  data: function() {
    return {
      showAddUserFormAdmin: false,
      showModifyUserForm: false,
      showDeleteUserForm: false,
      search: "",
      searchResults: [],
      selectedUser: {}
    };
  },
  methods: {
    searchUsers: function() {
      if (this.search != "") {
        var tmpSearch = this.search.toLowerCase();
        this.searchResults = this.users.filter(
          user =>
            user["BierName"].toLowerCase().includes(tmpSearch) |
            user["FirstName"].toLowerCase().includes(tmpSearch) |
            user["LastName"].toLowerCase().includes(tmpSearch)
        );
      } else {
        this.searchResults = [];
      }
    }
  }
};
</script>