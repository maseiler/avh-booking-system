<template>
  <div>
  <div class="columns search-input-area">
      <UserSearch />
      <button class="button is-link" @click="showAddUserForm = true">
        {{$t('booking.addUser.title')}}
      </button>
    <AddUserForm v-if="showAddUserForm" @close="showAddUserForm = false" />
  </div>
  <!-- <hr /> -->
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