<template>
  <div>
    <nav
      class="navbar is-primary is-fixed-top"
      role="navigation"
      aria-label="main navigation"
    >
      <div class="navbar-brand">
        <a class="navbar-item">
          <font-awesome-icon
            icon="beer"
            :style="{ color: 'black' }"
            size="lg"
          />
        </a>
      </div>

      <div id="navbarBasicExample" class="navbar-menu">
        <div class="navbar-start">
          <div class="navbar-item">
            <router-link class="button is-white" v-bind:to="'/'">
              Home
            </router-link>
          </div>

          <div class="navbar-item">
            <router-link class="button is-white" v-bind:to="'/booking'">
              Booking
            </router-link>
          </div>

          <div class="navbar-item">
            <router-link class="button is-white" v-bind:to="'/statistics'">
              Statistics
            </router-link>
          </div>
        </div>

        <div class="navbar-end">
          <div class="navbar-item">
            <router-link v-bind:to="'/login'">
              <font-awesome-icon
                icon="user-secret"
                :style="{ color: 'black' }"
                size="lg"
              />
            </router-link>
          </div>
        </div>
      </div>
    </nav>
  </div>
</template>
