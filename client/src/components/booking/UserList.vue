<template>
  <div>
    <div class="tabs">
      <ul>
        <li :class="[ activeTab === 'tab0' ? 'is-active' : '']">
          <a @click="activeTab='tab0'">Alle</a>
        </li>
        <li :class="[ activeTab === 'tab1' ? 'is-active' : '']">
          <a @click="activeTab='tab1'">Altherren</a>
        </li>
        <li :class="[ activeTab === 'tab2' ? 'is-active' : '']">
          <a @click="activeTab='tab2'">Aktive B</a>
        </li>
        <li :class="[ activeTab === 'tab3' ? 'is-active' : '']">
          <a @click="activeTab='tab3'">Aktive KA</a>
        </li>
        <li :class="[ activeTab === 'tab4' ? 'is-active' : '']">
          <a @click="activeTab='tab4'">Steganleger</a>
        </li>
        <li :class="[ activeTab === 'tab5' ? 'is-active' : '']">
          <a @click="activeTab='tab5'">Gäste</a>
        </li>
      </ul>
    </div>
    <div class="buttons" v-if="activeTab ==='tab0'">
      <button
        class="button"
        v-for="user in allUsers"
        :key="user"
        :class="buttonColor(user)"
        @click="selectUser(user)"
      >{{ displayUserName(user) }}</button>
    </div>
    <div class="buttons" v-if="activeTab ==='tab1'">
      <button
        class="button"
        v-for="user in usersAH"
        :key="user"
        :class="buttonColor(user)"
        @click="selectUser(user)"
      >{{ displayUserName(user) }}</button>
    </div>
    <div class="buttons" v-if="activeTab ==='tab2'">
      <button
        class="button"
        v-for="user in usersAktivB"
        :key="user"
        :class="buttonColor(user)"
        @click="selectUser(user)"
      >{{ displayUserName(user) }}</button>
    </div>
    <div class="buttons" v-if="activeTab ==='tab3'">
      <button
        class="button"
        v-for="user in usersAktivKA"
        :key="user"
        :class="buttonColor(user)"
        @click="selectUser(user)"
      >{{ displayUserName(user) }}</button>
    </div>
    <div class="buttons" v-if="activeTab ==='tab4'">
      <button
        class="button"
        v-for="user in usersSteganleger"
        :key="user"
        :class="buttonColor(user)"
        @click="selectUser(user)"
      >{{ displayUserName(user)}}</button>
    </div>
    <div class="buttons" v-if="activeTab ==='tab5'">
      <button
        class="button"
        v-for="user in usersGaeste"
        :key="user"
        :class="buttonColor(user)"
        @click="selectUser(user)"
      >{{ displayUserName(user)}}</button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      activeTab: "tab0",
      selectedUser: {}
    };
  },
  computed: {
    allUsers() {
      return this.$store.state.users;
    },
    usersAH() {
      return this.$store.getters.usersAH;
    },
    usersAktivB() {
      return this.$store.getters.usersAktivB;
    },
    usersAktivKA() {
      return this.$store.getters.usersAktivKA;
    },
    usersSteganleger() {
      return this.$store.getters.usersSteganleger;
    },
    usersGaeste() {
      return this.$store.getters.usersGast;
    }
  },
  methods: {
    buttonColor(user) {
      if (this.selectedUser === user) {
        return "is-link";
      } else if (user.Balance >= user.MaxDebt) {
        return "is-danger";
      } else if (user.MaxDebt - user.Balance <= user.MaxDebt * 0.1) {
        return "is-warning";
      } else {
        return "";
      }
    },
    selectUser(user) {
      if (this.selectedUser === user) {
        this.deselectUser();
        return;
      }
      this.selectedUser = user;
      this.$emit("selectUser", user);
      this.$userEventBus.$emit("selectUserToBus", user);
    },
    deselectUser(user) {
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
  created() {
    this.$userEventBus.$on("selectUserToBus", this.selectUserFromBus);
    this.$userEventBus.$on("deselectUserToBus", this.deselectUserFromBus);
  }
};
</script>