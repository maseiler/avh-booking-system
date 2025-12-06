<template>
<h1 class="title">Payment Settings</h1>
<h2 class="subtitle">Zahlungsdienstleister - Stripe</h2>
<div class="columns">
  <div class="column is-3">Aktivieren</div>
  <div class="column">
    <label class="checkbox">
      <input type="checkbox" @change="saveSettings">
      Stripe aktivieren
      <icon v-if="isSaving" :icon="['fas', 'spinner']" class="animateRight spinnerIcon" />
      <icon v-if="isSaved" :icon="['fas', 'check']" class="checkIcon" />
    </label>
  </div>
</div>

<div class="columns">
  <div class="column is-3">Stripe API Key</div>
  <div class="column">
    <p class="control has-icons-left has-icons-right">
      <input type="text" class="input" @blur="saveSettings">
      <span class="icon is-small is-left">
        <icon :icon="['fas', 'key']" />
      </span>
      <span class="icon is-right">
        <icon v-if="isSaving" :icon="['fas', 'spinner']" class="animateRight spinnerIcon" />
        <icon v-if="isSaved" :icon="['fas', 'check']" class="checkIcon" />
      </span>
    </p>
  </div>
</div>


</template>

<style scoped>
.checkIcon{
  animation: 2s linear mixOut forwards;
}

.animateRight{
  animation: 1s linear infinite turnRight;
}

@keyframes turnRight {
  from {
    transform:rotate(0deg);
  }
  to {
    transform:rotate(360deg);
  }
}

@keyframes mixOut{
  0% {
    opacity: 1;
  }
  50% {
    opacity: 1;
  }
  100% {
    opacity: 0;
  }
}
</style>

<script lang="ts">
  export default {
    data() {
      return {
        isSaving: false,
        isSaved: true,
      }
    },
    methods: {
      saveSettings(){
        this.isSaving = true;
        this.isSaved = false;
        setTimeout(()=>{
          this.isSaving = false;
          this.isSaved = true;
        }, 1500)
      }
    }
  }
</script>