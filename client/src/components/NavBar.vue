<template>
  <div>
    <nav class="navbar is-primary is-fixed-top" role="navigation" aria-label="main navigation">
      <div class="navbar-brand">
        <a class="navbar-item">
          <font-awesome-icon icon="beer" :style="{ color: 'black' }" size="lg" />
        </a>
      </div>

      <div id="navbarBasicExample" class="navbar-menu">
        <div class="navbar-start">
          <a class="navbar-item">
            <router-link v-bind:to="'/'">
              <a class="button is-white">Home</a>
            </router-link>
          </a>

          <a class="navbar-item">
            <router-link v-bind:to="'/booking'">
              <a class="button is-white">Booking</a>
            </router-link>
          </a>

          <div class="navbar-item">
            <router-link v-bind:to="'/statistics'">
              <a class="button is-white">Statistics</a>
            </router-link>
          </div>
        </div>

        <div class="navbar-end">
          <div class="navbar-item">
            <router-link v-bind:to="'/login'">
              <font-awesome-icon icon="user-secret" :style="{ color: 'black' }" size="lg" />
            </router-link>
          </div>
        </div>
      </div>
    </nav>
  </div>
</template>
