<template>
  <div>
    <br />
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
    <div class="tile is-ancestor is-narrow">
      <div class="tile is-parent is-vertical">
        <article class="tile is-child">
          <div class="box field is-grouped">
            <button
              class="button is-link is-1"
              @click="showAddUserFormAdmin = true"
            >
            {{$t('admin.userSettings.add')}}
            </button>
            &nbsp;
            <button
              class="button is-link is-1"
              @click="showModifyUserForm = true"
            >
            {{$t('admin.userSettings.modify')}}
            </button>
            &nbsp;
            <button
              class="button is-link is-1"
              @click="showDeleteUserForm = true"
            >
            {{$t('admin.userSettings.deleteUserTitle')}}
            </button>
          </div>
        </article>
        <article class="tile is-child">
          <div class="box" style="height: 17rem; overflow: auto">
            <UserSearch />
          </div>
        </article>
        <article class="tile is-child">
          <div
            class="box"
            style="height: calc(100vh - 26.25rem); overflow: auto"
          >
            <UserTable />
          </div>
        </article>
      </div>
    </div>
  </div>
</template>

<script>
import AddUserFormAdmin from "./AddUserFormAdmin.vue";
import ModifyUserForm from "./ModifyUserForm.vue";
import DeleteUserForm from "./DeleteUserForm.vue";
import UserSearch from "./UserSearch.vue";
import UserTable from "./UserTable.vue";

export default {
  components: {
    AddUserFormAdmin,
    ModifyUserForm,
    DeleteUserForm,
    UserSearch,
    UserTable,
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
