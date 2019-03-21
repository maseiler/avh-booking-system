<template>
  <div>
    <!-- search bar -->
    <div class="columns">
      <div class="column">
        <div class="field">
          <div class="control has-icons-left">
            <input
              class="input"
              type="text"
              placeholder="Search user"
              v-model="search"
              v-on:keyup="searchUsers"
            >
            <span class="icon is-small is-left">
              <font-awesome-icon icon="search"/>
            </span>
          </div>
        </div>
      </div>
    </div>
    <!-- search results -->
    <div class="buttons" v-if="searchResults !== []">
      <button
        class="button"
        v-for="user in searchResults"
        :key="user"
        :class="[selectedUser === user ? 'is-link' : '']"
        @click="selectUser(user)"
      >{{ displayName(user) }}</button>
    </div>
  </div>
</template>

<script src="./UserSearch.js"/>