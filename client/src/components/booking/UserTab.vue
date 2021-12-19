<template>
  <div>
    <div class="columns">
      <div class="column">
        <UserSearch />
      </div>
      <div class="column is-2">
        <button class="button is-link" @click="showAddUserForm = true">
          Add User
        </button>
      </div>
      <AddUserForm v-if="showAddUserForm" @close="showAddUserForm = false" />
    </div>
    <hr />
    <UserList />
  </div>
</template>

<script>
import UserSearch from "./UserSearch.vue";
import UserList from "./UserList.vue";
import AddUserForm from "./AddUserForm.vue";

export default {
  components: {
    UserSearch,
    UserList,
    AddUserForm,
  },
  data() {
    return {
      showAddUserForm: false,
    };
  },
};
</script>