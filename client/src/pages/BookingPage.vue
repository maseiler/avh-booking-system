<template>
  <div class="bookingPage">
    <h1 class="title is-1">TODO</h1>
    {{ users }}
  </div>
</template>
<script setup>
  import { storeToRefs } from 'pinia'
  import { useUserStore } from '@/stores/UserStore'

  const userStore = useUserStore()
  userStore.fetchUsers()

  const { users } = storeToRefs(userStore)
</script>
