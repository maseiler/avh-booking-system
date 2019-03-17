<template>
  <div>
    <h1>BOOKING SETTINGS</h1>
    <div class="columns">
      <div class="column is-2">
        <button class="button is-link" @click="showModifyBookEntryForm = true">Modify Book Entry</button>
        <ModifyBookEntryForm v-if="showModifyBookEntryForm" :entry="selectedEntry"></ModifyBookEntryForm>
      </div>
      <div class="column">
        <LastBookings
          :lastBookings="lastBookings"
          :users="users"
          :items="items"
          @selectEntry="selectEntry"
        ></LastBookings>
      </div>
    </div>
  </div>
</template>

<script src="./BookingSettings.js"></script>