<template>
  <transition name="modal">
    <div class="modal-mask">
      <div class="modal-wrapper">
        <div class="modal-container">
          <div class="modal-header">
            <h1 class="title is-4">{{$t('admin.userSettings.deleteUserTitle')}}</h1>
            <hr />
          </div>

          <div class="modal-body">
            <article
              v-if="Object.keys(user).length === 0"
              class="message is-danger"
            >
              <div class="message-header">
                <div class="field is-grouped">
                  <p class="icon is-small is-left">
                    <font-awesome-icon icon="exclamation" size="lg" />
                  </p>
                  <p>{{$t('admin.userSettings.selectFirst')}}</p>
                </div>
              </div>
            </article>
            <div v-else>
              <p>
                {{$t('admin.userSettings.deleteShure')}}
                <b>{{ displayUserNameFull(user) }}</b
                >?
              </p>
            </div>
          </div>

          <div class="modal-footer">
            <div class="level">
              <div class="level-left">
                <div class="level-item">
                  <div class="control" v-if="Object.keys(user).length !== 0">
                    <button class="button is-link" @click="deleteUser">
                      {{$t('generic.delete')}}
                    </button>
                  </div>
                </div>
              </div>
              <div class="level-right">
                <div class="level-item">
                  <button class="button" @click="closeForm">{{$t('generic.cancel')}}</button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </transition>
</template>

<script>
export default {
  computed: {
    user() {
      return this.$store.state.selectedUser;
    },
  },
  methods: {
    deleteUser() {
      if (Object.keys(this.user).length === 0) {
        this.error = "Please select an user first.";
      } else {
        this.$http
          .post("deleteUser", this.user)
          .then(() => {
            let message = `${this.$t('messages.success.deleteUser')}: ${this.displayUserNameFull(this.user)}`;
            this.$store.commit("selectUser", {});
            this.$store.commit("getUsers");
            this.closeForm();
            this.$responseEventBus.$emit("successMessage", message);
          })
          .catch(() => {
            this.$responseEventBus.$emit(
              "failureMessage",
              this.$t('messages.failure.userDelete')
            );
          });
      }
    },
    closeForm() {
      this.$emit("close");
    },
  },
};
</script>
<style lang="scss">
@import "../../assets/modal.css";
</style>
