<template>
  <section class="section">
    <h1 class="title">Client Settings</h1>
    <div class="field">
      <label class="label">Location Name</label>
      <div class="control">
        <input
          class="input"
          type="text"
          v-model="locationName"
          placeholder="e.g. Bar, Kitchen"
        />
      </div>
      <p class="help">Identifies this terminal to the server. Changes take effect after saving and reloading.</p>
    </div>
    <div class="field">
      <div class="control">
        <button class="button is-primary" :disabled="!locationName.trim()" @click="save">
          Save &amp; Reload
        </button>
      </div>
    </div>
  </section>
</template>

<script lang="ts">
export default {
  data() {
    return {
      locationName: localStorage.getItem('avhbs_client_id') ?? '',
    };
  },
  methods: {
    save() {
      const name = this.locationName.trim();
      if (!name) return;
      localStorage.setItem('avhbs_client_id', name);
      location.reload();
    },
  },
};
</script>
