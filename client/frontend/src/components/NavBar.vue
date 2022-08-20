<template>
  <nav class="navbar is-huette is-fixed-top" role="navigation">
    <div class="navbar-brand">
      <a class="navbar-item">
        <font-awesome-icon icon="beer" :style="{ color: 'black' }" size="lg" />
      </a>
      <a role="button" id="burger" class="navbar-burger" :class="active ? 'is-active' : ''" @click="showMobileMenu">
        <span></span>
        <span></span>
        <span></span>
      </a>
    </div>

    <div class="navbar-menu" :class="active ? 'is-active' : ''">
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
          <router-link v-bind:to="'/feedback'">
            <a class="button is-white">Feedback</a>
          </router-link>
        </div>

        <div class="navbar-item">
          <router-link v-bind:to="'/login'">
            <a class="button">
              <icon class="icon is-small is-left">
                <font-awesome-icon icon="gear" />
              </icon>
            </a>
          </router-link>
        </div>
      </div>
    </div>
  </nav>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  data() {
    return {
      isActive: '',
      active: false
    }
  },
  methods: {
    showMobileMenu() {
      this.active = !this.active;
    }
  }
})
</script>