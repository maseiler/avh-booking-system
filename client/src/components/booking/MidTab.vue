<template>
  <div>
    <button class="button" @click="refreshPage">
      <icon class="icon is-small">
        <font-awesome-icon icon="redo" />
      </icon>
      <p>{{$t('generic.refresh')}}</p>
    </button>
    <br />
    <div
      v-if="
        Object.keys(this.$store.state.selectedUser).length !== 0 ||
        Object.keys(this.$store.state.selectedMultipleItems).length !== 0
      "
    >
      <Cart />
      <hr />
      <UserInfo></UserInfo>
    </div>
    <div v-else>
      <img class="logo" src="~@/assets/avh_logo.png" width="250px" />
      <hr />
      <LastNBookEntries />
    </div>
  </div>
</template>

<script>
import Cart from "./Cart.vue";
import UserInfo from "./UserInfo.vue";
import LastNBookEntries from "./LastNBookEntries.vue";

export default {
  components: {
    Cart,
    UserInfo,
    LastNBookEntries,
  },
  methods: {
    refreshPage() {
      this.$router.go();
    },
  },
};
</script>