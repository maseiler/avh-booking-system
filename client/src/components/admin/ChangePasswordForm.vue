<template>
  <div>
    <div class="field">
      <div class="control">
        <input
          class="input"
          type="password"
          v-bind:placeholder="$t('admin.password.old')"
          v-model="oldPassword"
        />
      </div>
    </div>
    <div class="field">
      <div class="control">
        <input
          class="input"
          type="password"
          v-bind:placeholder="$t('admin.password.new')"
          v-model="newPassword1"
        />
      </div>
    </div>
    <div class="field">
      <div class="control">
        <input
          class="input"
          type="password"
          v-bind:placeholder="$t('admin.password.confirm')"
          v-model="newPassword2"
        />
      </div>
    </div>
    <article v-if="validationError !== ''" class="message is-danger">
      <div class="message-header">
        <div class="field is-grouped">
          <p class="icon is-small is-left">
            <font-awesome-icon icon="exclamation" size="lg" />
          </p>
          <p>{{ validationError }}</p>
        </div>
      </div>
    </article>
    <button class="button is-link" @click="changePassword">{{$t('generic.change')}}</button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      oldPassword: "",
      newPassword1: "",
      newPassword2: "",
      validationError: "",
    };
  },
  methods: {
    changePassword() {
      if (this.newPassword1 === this.newPassword2) {
        this.$http
          .post("changeAdminPassword", [this.oldPassword, this.newPassword1])
          .then(() => {
            this.$responseEventBus.$emit(
              "successMessage",
              this.$t('messages.success.passwordChangedAdmin')
            );
          })
          .catch((response) => {
            //ToDo: Internationalize this failure Message
            this.$responseEventBus.$emit("failureMessage", response.body);
          });
      } else {
        this.validationError = this.$t('messages.failure.passwordNoMatch');
      }
    },
  },
};
</script>

