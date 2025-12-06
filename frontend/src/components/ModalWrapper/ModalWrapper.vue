<template>
  <div class="modal" :class="modalIsActiveClass">
    <div class="modal-background" @click="modalIsActive = false"></div>
    <div class="modal-card">
      <header class="modal-card-head">
        <p class="modal-card-title">{{ modalTitle }}</p>
        <button class="delete" aria-label="close" @click="modalIsActive = false"></button>
      </header>
      <section class="modal-card-body">
        <slot name="header"></slot>
      </section>
      <footer class="modal-card-foot">
        <div class="buttons">
          <button class="button" @click="modalIsActive = false">Close</button>
          <slot name="footer"></slot>
        </div>
      </footer>
    </div>
  </div>
</template>

<style scoped>

</style>

<script lang="ts">

export default {
  data() {
    return {
      
    }
  },
  props: {
    modalTitle: String,
    modalIsActive: Boolean
  },
  computed: {
    modalIsActiveClass() {
      return this.modalIsActive ? "is-active" : "";
    }
  }
}
</script>