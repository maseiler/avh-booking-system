<template>
  <section class="hero is-fullheight">
    <div class="hero-body">
      <div class="container">
        <div class="columns is-centered">
          <div class="column is-4">
            <div class="box">
              <h1 class="title">Client Setup</h1>
              <p class="subtitle">Set a name for this terminal to identify its location.</p>
              <div class="field">
                <label class="label">Location Name</label>
                <div class="control">
                  <input
                    class="input"
                    type="text"
                    v-model="locationName"
                    placeholder="e.g. Bar, Kitchen"
                    @keyup.enter="save"
                  />
                </div>
              </div>
              <div class="field">
                <div class="control">
                  <button class="button is-primary" :disabled="!locationName.trim()" @click="save">
                    Save
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script lang="ts">
export default {
  emits: ['setup-complete'],
  data() {
    return {
      locationName: '',
    };
  },
  methods: {
    save() {
      const name = this.locationName.trim();
      if (!name) return;
      localStorage.setItem('avhbs_client_id', name);
      this.$emit('setup-complete');
    },
  },
};
</script>
