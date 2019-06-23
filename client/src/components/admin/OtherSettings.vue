<template>
  <div>
    <br>
    <div class="columns">
      <div class="column is-2">
        <div class="buttons">
          <button
            class="button is-link is-fullwidth"
            @click="showChangePasswordForm = true"
          >Change Password</button>
        </div>
      </div>
      <div class="column is-3">
        <ChangePasswordForm v-if="showChangePasswordForm"/>
      </div>
    </div>
  </div>
</template>

<script>
import ChangePasswordForm from "./ChangePasswordForm.vue";
export default {
  components: {
    ChangePasswordForm
  },
  data() {
    return {
      showChangePasswordForm: false
    };
  }
};
</script>
