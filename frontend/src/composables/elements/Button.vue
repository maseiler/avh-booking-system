<script setup lang="ts">

const props = defineProps<{
  faIcon?: string[],
  iconPosition?: string,
}>()

</script>

<template>
  <button class="button">
    <span v-if="props.faIcon !== undefined && props.iconPosition == 'left'" class="icon" :class="props.iconPosition">
      <icon :icon="props.faIcon" />
    </span>
    <span><slot></slot></span>
    <span v-if="props.faIcon !== undefined && props.iconPosition == 'right'" class="icon" :class="props.iconPosition">
      <icon :icon="props.faIcon" />
    </span>
  </button>
</template>