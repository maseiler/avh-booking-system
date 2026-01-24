<script setup lang="ts">



</script>

<template>
  <article class="message">
  <div class="message-header" v-if="$slots.header">
    <slot name="header"></slot>
  </div>
  <div class="message-body">
    <slot></slot>
  </div>
</article>
</template>