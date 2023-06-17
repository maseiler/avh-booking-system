<template>
  <div class="bookingPage">
    <div class="columns is-gapless">
      <div class="column">
        <div class="box max-height-scroll-y">
          <UserTab />
        </div>
      </div>
      <div class="column is-narrow">
        <div class="box max-height-scroll-y" style="width: 500px">
          <!--<MidTab />-->
          <h1>TODO</h1>
        </div>
      </div>
      <div class="column">
        <div class="box max-height-scroll-y">
          <!--<ItemTab />-->
          <h1>TODO</h1>
        </div>
      </div>
    </div>
    <h1 class="title is-1">TODO</h1>
    {{ users }}
  </div>
</template>
<script setup>
  import UserTab from '../components/booking/UserTab.vue'

  import { storeToRefs } from 'pinia'
  import { useUserStore } from '@/stores/UserStore'

  const userStore = useUserStore()
  userStore.fetchUsers()

  const { users } = storeToRefs(userStore)
</script>
