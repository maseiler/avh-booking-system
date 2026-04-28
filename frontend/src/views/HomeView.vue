<script setup lang="ts">
import { computed } from 'vue';
import { useSettingStore } from '../store/SettingStore';

const setting$ = useSettingStore();
const compTitle = computed(() => {
  let cT = setting$.get("compTitle");
  if (cT == -1){
    return "Set a Company Title in the Admin Settings"
  }
  return cT.value;
});

const compSlogan = computed(() => {
  let cT = setting$.get("compSlogan");
  if (cT == -1){
    return "Set a Slogan in the Admin Settings"
  }
  return cT.value;
})




</script>

<template>
  <div class="grid">
    <div class="cell">
      <h1 class="title">{{compTitle}}</h1>
    </div>
    <div class="cell">
      <img src="./" alt="">
    </div>
    <div class="cell">
      <span class="title is-4">{{compSlogan}}</span>
    </div>

  </div>
  
</template>

<style scoped>
.grid{
  align-items:center;
  justify-items: center;
  grid-template-columns: auto;
  grid-template-rows: 1fr auto 1fr;
}
</style>
