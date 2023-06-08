<template>
  <transition name="modal">
    <div class="modal-mask">
      <div class="modal-wrapper">
        <div class="modal-container">
          <div class="modal-header">
            <h1 class="title is-4">{{$t('admin.addUser.title')}}</h1>
            <hr />
          </div>

          <div class="modal-body">
            <div class="field">
              <label class="label">{{$t('user.nickname')}}</label>
              <div class="control has-icons-left">
                <input
                  class="input"
                  type="text"
                  v-bind:placeholder="$t('user.nickname')"
                  v-model.lazy="newUser.BierName"
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
                  v-model.lazy="newUser.FirstName"
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
                  v-model.lazy="newUser.LastName"
                />
                <span class="icon is-small is-left">
                  <font-awesome-icon icon="user" />
                </span>
              </div>
            </div>

            <div v-if="newUser.Status == 'Steganleger'" class="field">
              <label class="label">{{$t('user.boatName')}}</label>
              <div class="control has-icons-left">
                <input
                  class="input"
                  type="text"
                  v-bind:placeholder="$t('user.boatName')"
                  v-model.lazy="newUser.BoatName"
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
                  v-model.lazy="newUser.Email"
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
                  v-model.lazy="newUser.Phone"
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
                  <select v-model="newUser.Status">
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
                v-model.number="newUser.Balance"
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
                v-model.number="newUser.MaxDebt"
              />
              <span class="icon is-small is-left">
                <font-awesome-icon icon="money-bill" />
              </span>
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
                  <button class="button is-link" @click="submitUser">
                    {{$t('generic.submit')}}
                  </button>
                </div>
              </div>
              <div class="level-right">
                <div class="level-item">
                  <button class="button is-text" @click="resetAndCloseForm">
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
  data() {
    return {
      newUser: {
        BierName: "",
        FirstName: "",
        LastName: "",
        BoatName: "",
        Email: "",
        Phone: "",
        Status: "",
        Balance: 0,
        MaxDebt: 0,
      },
      validationError: "",
    };
  },
  methods: {
    submitUser() {
      this.$http
        .post("addUser", this.newUser)
        .then(() => {
          let message = `${this.$t('messages.success.newUserAdded')}: ${this.displayUserNameFull(this.newUser)}`;
          this.resetAndCloseForm();
          this.$store.commit("getUsers");
          this.$responseEventBus.$emit("successMessage", message);
        })
        .catch((response) => {
          //ToDo: internationalize this message
          this.validationError = response.data;
        });
    },
    resetAndCloseForm() {
      this.newUser.BierName = "";
      this.newUser.FirstName = "";
      this.newUser.LastName = "";
      this.newUser.BoatName = "";
      this.newUser.Status = "";
      this.newUser.Email = "";
      this.newUser.Phone = "";
      this.newUser.Balance = "";
      this.newUser.MaxDebt = "";
      this.validationError = "";
      this.$emit("close");
    },
  },
};
</script>
<style lang="scss">
@import "../../assets/modal.css";
</style>
