<template>
  <div>
    <br>
    <div class="columns">
      <div class="column is-1">
        <div class="buttons">
          <button class="button is-link is-fullwidth" @click="showAddUserFormAdmin = true">Add User</button>
          <button class="button is-link is-fullwidth" @click="showModifyUserForm = true">Modify User</button>
          <button class="button is-link is-fullwidth" @click="showDeleteUserForm = true">Delete User</button>
        </div>
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
      </div>
      <div class="column is-4">
        <UserSearch :allUsers="users" @selectUser="selectUser"/>
      </div>
      <div class="column is-4">
        <UserInfoAdmin :user="selectedUser"/>
      </div>
    </div>
  </div>
</template>

<script>
import AddUserFormAdmin from "./AddUserFormAdmin.vue";
import ModifyUserForm from "./ModifyUserForm.vue";
import DeleteUserForm from "./DeleteUserForm.vue";
import UserSearch from "../booking/UserSearch.vue";
import UserInfoAdmin from "./UserInfoAdmin.vue";

export default {
  components: {
    AddUserFormAdmin,
    ModifyUserForm,
    DeleteUserForm,
    UserSearch,
    UserInfoAdmin
  },
  props: {
    users: []
  },
  data: function() {
    return {
      showAddUserFormAdmin: false,
      showModifyUserForm: false,
      showDeleteUserForm: false,
      selectedUser: {}
    };
  },
  methods: {
    selectUserFromBus: function(user) {
      this.selectedUser = user;
    },
    deselectUserFromBus: function() {
      this.selectUserFromBus = {};
    }
  },
  created: function() {
    this.$userEventBus.$on("selectUserToBus", this.selectUserFromBus);
    this.$userEventBus.$on("deselectUserToBus", this.deselectUserFromBus);
  }
};
</script>