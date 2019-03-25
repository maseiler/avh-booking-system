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
      >{{ displayName(user) }}</button>
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
    displayName: function(user) {
      if (user.BierName != "") {
        return user.BierName;
      } else {
        if (user.LastName != "") {
          return user.FirstName + " " + user.LastName[0] + ".";
        }
        return user.FirstName;
      }
    },
    selectUser: function(user) {
      this.selectedUser = user;
      this.$emit("selectUser", user);
      this.$userEventBus.$emit("sendToBus", user);
    },
    receiveFromEventBus(user) {
      this.selectedUser = user;
    }
  },
  created: function() {
    this.$userEventBus.$on("sendToBus", this.receiveFromEventBus);
  }
};
</script>