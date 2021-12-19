<template>
  <div>
    <br />
    <div class="columns">
      <div class="column is-1">
        <div class="buttons">
          <button
            class="button is-link is-fullwidth"
            @click="showAddUserFormAdmin = true"
          >
            Add User
          </button>
          <button
            class="button is-link is-fullwidth"
            @click="showModifyUserForm = true"
          >
            Modify User
          </button>
          <button
            class="button is-link is-fullwidth"
            @click="showDeleteUserForm = true"
          >
            Delete User
          </button>
        </div>
        <AddUserFormAdmin
          v-if="showAddUserFormAdmin"
          @close="showAddUserFormAdmin = false"
        />
        <ModifyUserForm
          v-if="showModifyUserForm"
          @close="showModifyUserForm = false"
        />
        <DeleteUserForm
          v-if="showDeleteUserForm"
          @close="showDeleteUserForm = false"
        />
      </div>
      <div class="column is-4">
        <UserInfoAdmin />
        <UserSearch />
      </div>
      <div class="column is-7">
        <div class="box" style="height: 88vh; overflow: auto">
          <UserTable />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import AddUserFormAdmin from "./AddUserFormAdmin.vue";
import ModifyUserForm from "./ModifyUserForm.vue";
import DeleteUserForm from "./DeleteUserForm.vue";
import UserSearch from "../booking/UserSearch.vue";
import UserTable from "./UserTable.vue";
import UserInfoAdmin from "./UserInfoAdmin.vue";

export default {
  components: {
    AddUserFormAdmin,
    ModifyUserForm,
    DeleteUserForm,
    UserSearch,
    UserTable,
    UserInfoAdmin,
  },
  data() {
    return {
      showAddUserFormAdmin: false,
      showModifyUserForm: false,
      showDeleteUserForm: false,
    };
  },
};
</script>