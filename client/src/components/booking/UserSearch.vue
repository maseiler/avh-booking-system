<template>
  <div>
    <!-- search bar -->
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
            />
            <span class="icon is-small is-left">
              <font-awesome-icon icon="search" />
            </span>
          </div>
        </div>
      </div>
    </div>
    <!-- search results -->
    <div class="buttons" v-if="searchResults !== []">
      <button
        class="button"
        v-for="user in searchResults"
        :key="user"
        :class="buttonColor(selectedUser, user)"
        @click="selectUser(user)"
      >{{ displayUserName(user) }}</button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      search: "",
      searchResults: []
    };
  },
  computed: {
    allUsers() {
      return this.$store.state.users;
    },
    selectedUser() {
      return this.$store.state.selectedUser;
    }
  },
  methods: {
    searchUsers() {
      if (this.search != "") {
        var tmpSearch = this.search.toLowerCase();
        this.searchResults = this.allUsers.filter(
          user =>
            user["BierName"].toLowerCase().includes(tmpSearch) |
            user["FirstName"].toLowerCase().includes(tmpSearch) |
            user["LastName"].toLowerCase().includes(tmpSearch) |
            user["Email"].toLowerCase().includes(tmpSearch) |
            user["Phone"].toLowerCase().includes(tmpSearch) |
            user["BoatName"].toLowerCase().includes(tmpSearch)
        );
      } else {
        this.searchResults = [];
      }
    },
    selectUser(user) {
      this.$store.commit("selectUser", user);
    }
  }
};
</script>