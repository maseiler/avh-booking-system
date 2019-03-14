<template>
  <div>
    <div class="columns">
      <div class="column is-2">
        <div class="box">
          <aside class="menu">
            <p class="menu-label">Administration</p>
            <ul class="menu-list">
              <li>
                <a @click="showSetting('userSetting')" :class="[ showUserSettings ? 'is-active' : '']">User Settings</a>
              </li>
              <li>
                <a @click="showSetting('itemSetting')" :class="[ showItemSettings ? 'is-active' : '']">Item Settings</a>
              </li>
              <li>
                <a @click="showSetting('otherSetting')" :class="[ showOtherSettings ? 'is-active' : '']">Other Settings</a>
              </li>
            </ul>
          </aside>
        </div>
      </div>
      <div class="column">
        <UserSettings v-if="showUserSettings"></UserSettings>
        <ItemSettings v-if="showItemSettings"></ItemSettings>
        <OtherSettings v-if="showOtherSettings"></OtherSettings>
      </div>
    </div>
  </div>
</template>

<script src="./Admin.js"></script>