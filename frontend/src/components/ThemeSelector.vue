<template>
  <div class="navbar-item has-dropdown is-hoverable">
    <a class="navbar-link">{{ theme$.label }}</a>
    <div class="navbar-dropdown">
      <a v-for="t in themes" :key="t.key"
         class="navbar-item" :class="{ 'is-active': theme$.currentTheme === t.key }"
         @click="theme$.applyTheme(t.key)">
        {{ t.label }}
      </a>
    </div>
  </div>
</template>

<style scoped>
.navbar-item.is-active {
  background-color: transparent !important;
  color: inherit !important;
  font-weight: 600;
  border-left: 3px solid hsl(var(--bulma-primary-h), var(--bulma-primary-s), var(--bulma-primary-l));
  padding-left: calc(1rem - 3px);
}
</style>

<script lang="ts">
import { useThemeStore } from '../store/themeStore'

export default {
  data() {
    return {
      theme$: useThemeStore(),
      themes: [
        { key: 'catppuccin-latte',     label: 'Latte' },
        { key: 'catppuccin-frappe',    label: 'Frappé' },
        { key: 'catppuccin-macchiato', label: 'Macchiato' },
        { key: 'catppuccin-mocha',     label: 'Mocha' },
      ]
    }
  }
}
</script>
