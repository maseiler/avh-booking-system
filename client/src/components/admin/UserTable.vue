<template>
  <div>
    <div class="table-container">
      <table class="table is-hoverable is-striped">
        <thead>
          <tr>
            <th>ID</th>
            <th>First Name</th>
            <th>Bier Name</th>
            <th>Last Name</th>
            <th>Boat Name</th>
            <th>Status</th>
            <th>Email</th>
            <th>Phone</th>
            <th>Balance</th>
            <th>Max Debt</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="user in users"
            :key="user"
            @click="selectUser(user)"
            :class="[selectedUser === user ? 'is-selected' : '']"
          >
            <th>{{user.ID}}</th>
            <td>{{user.FirstName}}</td>
            <td>{{user.BierName}}</td>
            <td>{{user.LastName}}</td>
            <td>{{user.BoatName}}</td>
            <td>{{user.Status}}</td>
            <td>{{user.Email}}</td>
            <td>{{user.Phone}}</td>
            <td>{{user.Balance}} €</td>
            <td>{{user.MaxDebt}} €</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  computed: {
    users() {
      return this.$store.state.users.sort(function(a, b) {
        if (a.ID < b.ID) return -1;
        if (a.ID > b.ID) return 1;
        return 0;
      });
    },
    selectedUser() {
      return this.$store.state.selectedUser;
    }
  },
  methods: {
    selectUser(user) {
      this.$store.commit("selectUser", user);
    }
  }
};
</script>
