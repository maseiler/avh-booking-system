<template>
  <div>
    <nav class="navbar is-primary" role="navigation" aria-label="main navigation">
      <div class="navbar-brand">
        <a class="navbar-item">
          <font-awesome-icon icon="euro-sign" :style="{ color: 'black' }" size="lg"/>
          <!-- TODO: logo or  -->
          <!-- <img src="~@/assets/Web-Wap.png"> -->
          <!-- <img src="https://bulma.io/images/bulma-logo.png" width="112" height="28"> -->
        </a>

        <!-- <a role="button" class="navbar-burger burger" aria-label="menu" aria-expanded="false" data-target="navbarBasicExample">
      <span aria-hidden="true"></span>
      <span aria-hidden="true"></span>
      <span aria-hidden="true"></span>
        </a>-->
      </div>

      <div id="navbarBasicExample" class="navbar-menu">
        <div class="navbar-start">
          <a class="navbar-item">
            <router-link v-bind:to="'/'">
              <!-- <font color="black">Home</font> -->
              <a class="button is-white">Home</a>
            </router-link>
          </a>

          <a class="navbar-item">
            <router-link v-bind:to="'/booking'">
              <!-- <font color="black">Booking</font>  -->
              <a class="button is-white">Booking</a>
            </router-link>
          </a>

          <div class="navbar-item">
            <router-link v-bind:to="'/statistics'">
              <!-- <font color="black">Statistics</font>  -->
              <a class="button is-white">Statistics</a>
            </router-link>
          </div>
        </div>

        <div class="navbar-end">
          <div class="navbar-item">
            <!-- <router-link v-bind:to="'/admin'"> -->
              <font-awesome-icon
                icon="user-secret"
                @click="showLogin = true"
                :style="{ color: 'black' }"
                size="lg"
              />
            <!-- </router-link> -->
          </div>
        </div>
      </div>
    </nav>
    <AdminLogin v-if="showLogin" @close="showLogin = false"/>
  </div>
</template>

<script>
import AdminLogin from "../components/admin/Login.vue";
export default {
  components: {
    AdminLogin
  },
  data: function() {
    return {
      showLogin: false
    };
  },
  methods: {}
};
</script>
<style lang="scss">
@import "../assets/modal.css";
</style>

