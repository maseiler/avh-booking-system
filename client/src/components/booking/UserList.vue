<template>
  <div>
    <div class="tabs">
      <ul>
        <li :class="[activeTab === 'tab0' ? 'is-active' : '']">
          <a @click="activeTab = 'tab0'">Alle</a>
        </li>
        <li :class="[activeTab === 'tab1' ? 'is-active' : '']">
          <a @click="activeTab = 'tab1'">Altherren</a>
        </li>
        <li :class="[activeTab === 'tab2' ? 'is-active' : '']">
          <a @click="activeTab = 'tab2'">Aktive B</a>
        </li>
        <li :class="[activeTab === 'tab3' ? 'is-active' : '']">
          <a @click="activeTab = 'tab3'">Aktive KA</a>
        </li>
        <li :class="[activeTab === 'tab4' ? 'is-active' : '']">
          <a @click="activeTab = 'tab4'">Steganleger</a>
        </li>
        <li :class="[activeTab === 'tab5' ? 'is-active' : '']">
          <a @click="activeTab = 'tab5'">Gäste</a>
        </li>
      </ul>
    </div>
    <div v-if="activeTab === 'tab0'">
      <UserDictionary :userDict="allUsers">
    </div>
    <div v-if="activeTab === 'tab1'">
      <UserDictionary :userDict="usersAH">
    </div>
    <div v-if="activeTab === 'tab2'">
    <UserDictionary :userDict="usersAktivB">
    </div>
    <div v-if="activeTab === 'tab3'">
      <UserDictionary :userDict="usersAktivKA">
    </div>
    <div v-if="activeTab === 'tab4'">
      <UserDictionary :userDict="usersSteganleger">
    </div>
    <div v-if="activeTab === 'tab5'">
      <UserDictionary :userDict="usersGaeste">
    </div>
  </div>
</template>

<script>
import helper from "../../helper.js";
import UserDictionary from "./UserDictionary.vue";

export default {
  components: {
    UserDictionary,
  },
  data() {
    return {
      activeTab: "tab0",
    };
  },
  created() {
    this.$root.$on("set-user-tab-to-all", () => {
      this.activeTab = "tab0";
    });
  },
  computed: {
    allUsers() {
      return helper.getUsersAsDict(this.$store.state.users);
    },
    usersAH() {
      return helper.getUsersAsDict(this.$store.getters.usersAH);
    },
    usersAktivB() {
      return helper.getUsersAsDict(this.$store.getters.usersAktivB);
    },
    usersAktivKA() {
      return helper.getUsersAsDict(this.$store.getters.usersAktivKA);
    },
    usersSteganleger() {
      return helper.getUsersAsDict(this.$store.getters.usersSteganleger);
    },
    usersGaeste() {
      return helper.getUsersAsDict(this.$store.getters.usersGast);
    },
  },
};
</script>