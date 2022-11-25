<template>
  <div>
    <nav
      class="navbar is-primary is-fixed-top"
      role="navigation"
      aria-label="main navigation"
    >
      <div class="navbar-brand">
        <a class="navbar-item" @click="() => switchColorScheme(false)">
          <font-awesome-icon
            icon="beer"
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
                size="lg"
              />
            </router-link>
          </div>
        </div>
      </div>
    </nav>
  </div>
</template>
<script>
export default {
  methods: {
    switchColorScheme(stored){
      let body = document.getElementsByTagName("body")[0]
      //posibility to add more themes
      let themeList = ["light", "dark"]
      let currentThemeIndex = 0
      //get current Theme
      for(let i = 0; i < themeList.length; i++){
        if(body.classList.contains(themeList[i])){
          currentThemeIndex = i
        }
      }
      if(stored){
        body.classList.replace(themeList[currentThemeIndex], localStorage.getItem("colorScheme"))
        return
      }
      //rotate to next Theme, start over if out of bounds
      let nextThemeIndex = currentThemeIndex +1
      if(nextThemeIndex == themeList.length){
        nextThemeIndex = 0
      }
      //set next Theme
      localStorage.setItem("colorScheme", themeList[nextThemeIndex])
      body.classList.replace(themeList[currentThemeIndex], themeList[nextThemeIndex])
    }
  },
  mounted: function() {      
    this.switchColorScheme(localStorage.getItem("colorScheme") != null)
  }
}
</script>