<template>
  <div class="user-list">
    <div
      class="columns"
      v-for="(users, key) in userDict"
      :key="key"
      style="margin-bottom: 0cm"
    >
      <div class="column is-1">
        <h3 class="subtitle is-3 has-text-grey-light has-text-weight-bold">
          {{ key }}
        </h3>
      </div>
      <div class="column">
        <div class="buttons">
          <button
            class="button"
            v-for="user in users"
            :key="user"
            :class="buttonColor(selectedUser, user)"
            @click="selectUser(user)"
          >
            {{ displayUserName(user) }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    userDict: {},
  },
  computed: {
    selectedUser() {
      return this.$store.state.selectedUser;
    },
  },
};
</script>
