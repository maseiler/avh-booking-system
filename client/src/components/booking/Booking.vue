<template>
  <div class="booking">
    <div class="columns is-gapless">
      <div class="column">
        <div class="box">
          <UserTab></UserTab>
        </div>
        
      </div>
      <div class="column is-narrow">
        <div class="box has-text-centered" style="width: 500px;">
          <img class="logo" src="~@/assets/avh_logo.png" width="250px">
        </div>
      </div>
      <div class="column">
        <div class="box">
          <ItemTab @selectedItem="getItem"></ItemTab>
        </div>
      </div>
    </div>
  </div>
</template>

<script src="./Booking.js"></script>