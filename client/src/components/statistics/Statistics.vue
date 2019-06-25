<template>
  <div>
    <br>
    <img class="logo" src="~@/assets/WIP.png">
    <br>
    <!-- {{users[0].FirstName}} -->
    {{usersAH[0].BierName}}
  </div>
</template>

<script>
export default {
  computed: {
    users() {
      return this.$store.state.users;
    },
    usersAH(){
      return this.$store.getters.usersAH
    }
  }
};
</script>
