<template>
  <div class="booking">
    <div class="columns is-gapless">
      <div class="column">
        <div class="box">
          <p class="title">Users</p>
          <button @click="boolAddUser = true">Add User</button>
          <AddUserForm v-if="boolAddUser" @close="boolAddUser = false"></AddUserForm>
        </div>

      </div>
      <div class="column is-narrow">
        <div class="box has-text-centered" style="width: 500px;">
          <img class="logo" src="~@/assets/avh_logo_2.png" width="250px">
        </div>
      </div>
      <div class="column">
        <div class="box">
          <ItemList></ItemList>
        </div>
      </div>
    </div>
  </div>
</template>

<script src="./Booking.js"></script>