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
            >
            <span class="icon is-small is-left">
              <font-awesome-icon icon="search"/>
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
        :class="[selectedUser === user ? 'is-link' : '']"
        @click="selectUser(user)"
      >{{ displayUserName(user) }}</button>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    allUsers: []
  },
  data: function() {
    return {
      selectedUser: {},
      search: "",
      searchResults: []
    };
  },
  methods: {
    searchUsers: function() {
      if (this.search != "") {
        var tmpSearch = this.search.toLowerCase();
        this.searchResults = this.allUsers.filter(
          user =>
            user["BierName"].toLowerCase().includes(tmpSearch) |
            user["FirstName"].toLowerCase().includes(tmpSearch) |
            user["LastName"].toLowerCase().includes(tmpSearch)
        );
      } else {
        this.searchResults = [];
      }
    },
    selectUser: function(user) {
      if (this.selectedUser === user) {
        this.deselectUser();
        return;
      }
      this.selectedUser = user;
      this.$emit("selectUser", user);
      this.$userEventBus.$emit("selectUserToBus", user);
    },
    deselectUser: function(user) {
      this.selectedUser = {};
      this.$emit("selectUser", this.selectedUser);
      this.$userEventBus.$emit("deselectUserToBus");
    },
    deselectUserFromBus() {
      this.selectedUser = {};
    },
    selectUserFromBus(user) {
      this.selectedUser = user;
    }
  },
  created: function() {
    this.$userEventBus.$on("selectUserToBus", this.selectUserFromBus);
    this.$userEventBus.$on("deselectUserToBus", this.deselectUserFromBus);
  }
};
</script>