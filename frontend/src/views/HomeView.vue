<script setup lang="ts">
import { computed } from 'vue';
import { useSettingStore } from '../store/SettingStore';

const setting$ = useSettingStore();
const compTitle = computed(() => {
  let cT = setting$.get("compTitle");
  if (cT == -1){
    return "Set a Company Title in the General Settings > Company Settings"
  }
  return cT.value;
});

const compSlogan = computed(() => {
  let cT = setting$.get("compSlogan");
  if (cT == -1){
    return "Set a Slogan in the General Settings > Company Settings"
  }
  return cT.value;
})

const companyLogo = computed(() => {
  if (setting$.get("compLogo") == -1){
      return null;
    }
    return setting$.get("compLogo").value;
})



</script>

<template>
  <div class="grid">
    <div class="cell">
      <h1 class="title is-1">{{compTitle}}</h1>
    </div>
    <div class="cell">
      <img v-if="companyLogo != null" :src="companyLogo">
      <p v-if="companyLogo == null">Set a Company Logo in the General Settings > Company Settings</p>
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
