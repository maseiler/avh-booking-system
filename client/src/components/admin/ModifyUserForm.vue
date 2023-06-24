<template>
  <transition name="modal">
    <div class="modal-mask">
      <div class="modal-wrapper">
        <div class="modal-container">
          <div class="modal-header">
            <h1 class="title is-4">{{$t('admin.userSettings.modify')}}</h1>
            <hr />
          </div>

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
            <div class="modal-body">
              <div class="field">
                <label class="label">{{$t('user.nickname')}}</label>
                <div class="control has-icons-left">
                  <input
                    class="input"
                    type="text"
                    v-bind:placeholder="$t('user.nickname')"
                    v-model.lazy="user.BierName"
                  />
                  <span class="icon is-small is-left">
                    <font-awesome-icon icon="beer" />
                  </span>
                </div>
              </div>

              <div class="field">
                <label class="label">{{$t('user.firstName')}}</label>
                <div class="control has-icons-left">
                  <input
                    class="input"
                    type="text"
                    v-bind:placeholder="$t('user.firstName')"
                    v-model.lazy="user.FirstName"
                  />
                  <span class="icon is-small is-left">
                    <font-awesome-icon icon="user" />
                  </span>
                </div>
              </div>

              <div class="field">
                <label class="label">{{$t('user.lastName')}}</label>
                <div class="control has-icons-left">
                  <input
                    class="input"
                    type="text"
                    v-bind:placeholder="$t('user.lastName')"
                    v-model.lazy="user.LastName"
                  />
                  <span class="icon is-small is-left">
                    <font-awesome-icon icon="user" />
                  </span>
                </div>
              </div>

              <div v-if="user.Status == 'Steganleger'" class="field">
                <label class="label">{{$t('user.boatName')}}</label>
                <div class="control has-icons-left">
                  <input
                    class="input"
                    type="text"
                    v-bind:placeholder="$t('user.boatName')"
                    v-model.lazy="user.BoatName"
                  />
                  <span class="icon is-small is-left">
                    <font-awesome-icon icon="anchor" />
                  </span>
                </div>
              </div>

              <div class="field">
                <label class="label">{{$t('user.eMail')}}</label>
                <div class="control has-icons-left">
                  <input
                    class="input"
                    type="email"
                    v-bind:placeholder="$t('user.eMail')"
                    v-model.lazy="user.Email"
                  />
                  <span class="icon is-small is-left">
                    <font-awesome-icon icon="envelope" />
                  </span>
                </div>
              </div>

              <div class="field">
                <label class="label">{{$t('user.phoneNumber')}}</label>
                <div class="control has-icons-left">
                  <input
                    class="input"
                    type="text"
                    v-bind:placeholder="$t('user.phoneNumber')"
                    v-model.lazy="user.Phone"
                  />
                  <span class="icon is-small is-left">
                    <font-awesome-icon icon="phone" />
                  </span>
                </div>
              </div>

              <div class="field">
                <label class="label">{{$t('user.status')}}</label>
                <div class="control">
                  <div class="select">
                    <select v-model="user.Status">
                      <option disabled value>{{$t('user.status')}}</option>
                      <!-- ToDo: Fetch Status from DB instead of hard code -->
                      <option>Aktiv B</option>
                      <option>Aktiv KA</option>
                      <option>AH</option>
                      <option>Steganleger</option>
                      <option>Gast</option>
                    </select>
                  </div>
                </div>
              </div>
            </div>

            <div class="field">
              <label class="label">{{$t('generic.balance')}}</label>
              <div class="control has-icons-left">
                <input
                  class="input"
                  type="text"
                  v-bind:placeholder="$t('generic.balance')"
                  v-model.number="user.Balance"
                />
                <span class="icon is-small is-left">
                  <font-awesome-icon icon="euro-sign" />
                </span>
              </div>
            </div>

            <div class="field">
              <label class="label">{{$t('generic.max')}} {{$t('generic.debt')}}</label>
              <div class="control has-icons-left">
                <input
                  class="input"
                  type="text"
                  v-bind:placeholder="`${$t('generic.max')} ${$t('generic.debt')}`"
                  v-model.number="user.MaxDebt"
                />
                <span class="icon is-small is-left">
                  <font-awesome-icon icon="money-bill" />
                </span>
              </div>
            </div>
          </div>

          <div class="modal-footer">
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
            <div class="level">
              <div class="level-left">
                <div class="level-item">
                  <button class="button is-link" @click="modifyUser">
                    {{$t('generic.modify')}}
                  </button>
                </div>
              </div>
              <div class="level-right">
                <div class="level-item">
                  <button class="button is-text" @click="closeForm">
                    {{$t('generic.cancel')}}
                  </button>
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
  data() {
    return {
      validationError: "",
    };
  },
  methods: {
    modifyUser() {
      if (Object.keys(this.user).length === 0) {
        this.validationError = this.$t('admin.userSettings.selectFirst');
      } else {
        this.$http
          .post("modifyUser", this.user)
          .then(() => {
            let message = `${this.$t('messages.success.modifiedUser')}: ${this.displayUserNameFull(this.user)}`;
            this.$store.commit("selectUser", {});
            this.$store.commit("getUsers");
            this.closeForm();
            this.$responseEventBus.$emit("successMessage", message);
          })
          .catch((response) => {
            //ToDo: internationalize this failure message
            this.validationError = response.body;
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
