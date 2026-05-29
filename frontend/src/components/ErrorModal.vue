<script setup lang="ts">
defineProps<{
  modelValue: boolean
  error: { code: string, message: string, details?: string } | null
}>()

const emit = defineEmits<{
  'update:modelValue': [value: boolean]
}>()

function close() {
  emit('update:modelValue', false)
}
</script>

<template>
  <div class="modal" :class="{ 'is-active': modelValue }">
    <div class="modal-background" @click="close"></div>
    <div class="modal-card">
      <header class="modal-card-head">
        <p class="modal-card-title">
          <span class="icon-text has-text-danger">
            <span class="icon"><icon :icon="['fas', 'exclamation-circle']" /></span>
            <span>Error</span>
          </span>
        </p>
        <button class="delete" aria-label="close" @click="close"></button>
      </header>
      <section class="modal-card-body" v-if="error">
        <p class="has-text-weight-semibold">{{ error.message }}</p>
        <pre v-if="error.details" class="mt-3">{{ error.details }}</pre>
        <p class="mt-3 has-text-grey is-size-7">Code: {{ error.code }}</p>
      </section>
      <footer class="modal-card-foot">
        <button class="button" @click="close">Close</button>
      </footer>
    </div>
  </div>
</template>
