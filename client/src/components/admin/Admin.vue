<template>
  <div>
    <div class="columns">
      <div class="column is-2">
        <div class="box">
          <aside class="menu">
            <p class="menu-label">Administration</p>
            <ul class="menu-list">
              <li>
                <a @click="showSetting('userSettings')" :class="[ showUserSettings ? 'is-active' : '']">User Settings</a>
              </li>
              <li>
                <a @click="showSetting('itemSettings')" :class="[ showItemSettings ? 'is-active' : '']">Item Settings</a>
              </li>
              <li>
                <a @click="showSetting('bookingSettings')" :class="[ showBookingSettings ? 'is-active' : '']">Booking Settings</a>
              </li>
              <li>
                <a @click="showSetting('otherSettings')" :class="[ showOtherSettings ? 'is-active' : '']">Other Settings</a>
              </li>
            </ul>
          </aside>
        </div>
      </div>
      <div class="column">
        <UserSettings v-if="showUserSettings" :users="users"></UserSettings>
        <ItemSettings v-if="showItemSettings" :items="items"></ItemSettings>
        <BookingSettings v-if="showBookingSettings" :lastBookings="lastBookings" :users="users" :items="items"></BookingSettings>
        <OtherSettings v-if="showOtherSettings"></OtherSettings>
      </div>
    </div>
  </div>
</template>

<script src="./Admin.js"></script>