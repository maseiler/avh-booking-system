<template>
  <div>
    <nav
      class="navbar is-primary is-fixed-top"
      role="navigation"
      aria-label="main navigation"
    >
      <div class="navbar-brand">
        <a class="navbar-item" @click="() => switchColorScheme(false, -1)">
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
          <div class="navbar-item darkLightSwitch">
            <font-awesome-icon icon="sun" />
            <input type="checkbox" class="darkSwitch" id="darkSwitch" @click="darkSwitchClick"/>
            <label for="darkSwitch"></label>
            <font-awesome-icon icon="moon" />
          </div>
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
  data: function () {
    return {
      timeInterval:null
    };
  },
  methods: {
    switchColorScheme(stored, nextItem){
      let body = document.getElementsByTagName("body")[0]
      //to add more themes, simply add them to this array.
      // TODO: Add the setting for Admin in UI to add themes and store the names in the database
      // TODO: Add color settings for Admin in UI to simply build own color themes in UI 
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
        this.setDarkSwitch(localStorage.getItem("colorScheme"));
        return
      }
      //rotate to next Theme, start over if out of bounds
      let nextThemeIndex = currentThemeIndex +1;
      if(nextItem != -1){
        nextThemeIndex = nextItem;
      }
      if(nextThemeIndex == themeList.length){
        nextThemeIndex = 0
      }
      //set next Theme
      localStorage.setItem("colorScheme", themeList[nextThemeIndex])
      body.classList.replace(themeList[currentThemeIndex], themeList[nextThemeIndex])
      this.setDarkSwitch(themeList[nextThemeIndex]);
    },
    darkSwitchClick(){
      let switchCB = document.getElementById('darkSwitch');
      if(switchCB.checked){
        this.switchColorScheme(false, 1);
        return;
      }
      this.switchColorScheme(false, 0);
    },
    setDarkSwitch(dark){
      let switchCB = document.getElementById('darkSwitch');
      switchCB.checked = false;
      if(dark == "dark"){
        switchCB.checked = true;
      }
    },
    timeHandler(){
      let currentTime = Date();
      let currentTimeInt = parseInt(currentTime.substr(16, 24).split(":").join());
      // TODO: either get the day/night Start Values from the Database or use the sunrise/sunset time directly from weather API
      // TODO: check in database if automatic theme switching is enabled
      let dayStart = 80000; // 8 Uhr in the morning
      let nightStart = 180000; // 6 in the evening
      //19 Second time period, so it must run once in the 10s Interval
      if(currentTimeInt < nightStart + 9 && currentTimeInt > nightStart - 10 ){
        //make it dark
        this.switchColorScheme(false, 1);
        return;
      }
      if(currentTimeInt < dayStart + 9 && currentTimeInt > dayStart - 10 ){
        //make it light
        this.switchColorScheme(false, 0);
        return;
      }
      return;
    }
  },
  mounted: function() {      
    this.switchColorScheme(localStorage.getItem("colorScheme") != null, -1);
    if(this.timeInterval == null){
      this.timeInterval = window.setInterval(this.timeHandler, 10000);
    }
  }
}
</script>