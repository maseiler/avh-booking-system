<script lang="ts" setup>

</script>

<template>
<details class="details">
  <summary class="summary">
    <icon :icon="['fas', 'caret-right']" class="marker-icon-closed"></icon>
    <icon :icon="['fas', 'caret-down']" class="marker-icon-open"></icon>
    <slot name="summary"></slot>
  </summary>
  <div class="details-content">
    <slot></slot>
  </div>
</details>
</template>

<style scoped>
.summary{
  display:flex;
  align-items: center;
  gap:1em;
  padding:1em 0.5em;
  border-top:var(--bulma-hr-height) solid var(--bulma-hr-background-color);
}
.details[open]{
  .marker-icon-closed{
    display:none;
  }
  .marker-icon-open{
    display: inline;
  }
}
.details{
  .marker-icon-open{
    display:none;
  }
}
.details-content{
  padding-bottom: 2em;
}
</style>