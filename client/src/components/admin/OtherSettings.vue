<template>
  <div>
    <br />
    <div class="tile is-ancestor is-narrow">
      <div class="tile is-parent is-vertical">
        <article class="tile is-child">
          <div class="box">
            <div class="columns">
              <div class="column is-2">
                <div class="buttons">
                  <button
                    class="button is-link is-fullwidth"
                    @click="showChangePasswordForm = true"
                  >
                  {{$t('admin.password.change')}}
                  </button>
                </div>
              </div>
              <div class="column is-3">
                <ChangePasswordForm v-if="showChangePasswordForm" />
              </div>
            </div>
          </div>
        </article>
      </div>
    </div>
  </div>
</template>

<script>
import ChangePasswordForm from "./ChangePasswordForm.vue";
export default {
  components: {
    ChangePasswordForm,
  },
  data() {
    return {
      showChangePasswordForm: false,
    };
  },
};
</script>
