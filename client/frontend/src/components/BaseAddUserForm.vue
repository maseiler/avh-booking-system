<template>
  <transition name="modal">
    <div class="modal-mask">
      <div class="modal-wrapper">
        <div class="modal-container">
          <div class="modal-header">
            <h1 class="title is-4">Add new user</h1>
            <hr />
          </div>

          <div class="modal-body">
            <div class="field">
              <label class="label">Status</label>
              <div class="control">
                <div class="select">

                  <select v-model="newUser.Status">
                    <option disabled value>Status</option>
                    <option>Aktiv B</option>
                    <option>Aktiv KA</option>
                    <option>AH</option>
                    <option>Steganleger</option>
                    <option>Gast</option>
                  </select>

                </div>
              </div>
            </div>

            <div v-if="
              newUser.Status !== 'Steganleger' &&
              newUser.Status !== 'Gast' &&
              newUser.Status !== ''
            " class="field">
              <label class="label">Biername</label>
              <div class="control has-icons-left">
                <input class="input" type="text" placeholder="Biername" v-model.lazy="newUser.BierName" />
                <span class="icon is-small is-left">
                  <font-awesome-icon icon="beer" />
                </span>
              </div>
            </div>

            <div v-if="newUser.Status !== ''" class="field">
              <label class="label">First name</label>
              <div class="control has-icons-left">
                <input class="input" type="text" placeholder="First name" v-model.lazy="newUser.FirstName" />
                <span class="icon is-small is-left">
                  <font-awesome-icon icon="user" />
                </span>
              </div>
            </div>

            <div v-if="newUser.Status !== ''" class="field">
              <label class="label">Last name</label>
              <div class="control has-icons-left">
                <input class="input" type="text" placeholder="Last name" v-model.lazy="newUser.LastName" />
                <span class="icon is-small is-left">
                  <font-awesome-icon icon="user" />
                </span>
              </div>
            </div>

            <div v-if="newUser.Status === 'Steganleger'" class="field">
              <label class="label">Boat Name</label>
              <div class="control has-icons-left">
                <input class="input" type="text" placeholder="Boat Name" v-model.lazy="newUser.BoatName" />
                <span class="icon is-small is-left">
                  <font-awesome-icon icon="anchor" />
                </span>
              </div>
            </div>

            <div v-if="
              newUser.Status === 'Steganleger' || newUser.Status === 'Gast'
            " class="field">
              <label class="label">Email</label>
              <div class="control has-icons-left">
                <input class="input" type="email" placeholder="Email Address" v-model.lazy="newUser.Email" />
                <span class="icon is-small is-left">
                  <font-awesome-icon icon="envelope" />
                </span>
              </div>
            </div>

            <div v-if="
              newUser.Status === 'Steganleger' || newUser.Status === 'Gast'
            " class="field">
              <label class="label">Phone Number</label>
              <div class="control has-icons-left">
                <input class="input" type="text" placeholder="Phone Number" v-model.lazy="newUser.Phone" />
                <span class="icon is-small is-left">
                  <font-awesome-icon icon="phone" />
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
                  <button class="button is-link" @click="submitUser">
                    Submit
                  </button>
                </div>
              </div>
              <div class="level-right">
                <div class="level-item">
                  <button class="button is-text" @click="resetAndCloseForm">
                    Cancel
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

<script lang="ts">
import { defineComponent } from 'vue'
import User, { nullUser } from '../models/User'

export default defineComponent({
  data() {
    return {
      newUser: nullUser as User,
      validationError: ""
    }
  },
  methods: {
    submitUser(): void {
      window.go.main.App.AddNewUser(this.newUser).then((response: boolean) => {
        // TODO notification
        if (response) {
          console.log("Added new user")
        } else {
          console.log("Failed to add new user")
        }
        this.resetAndCloseForm()
      })
        .catch((response: any) => { // TODO type
          console.log(response)
        })
      // TODO trigger fetching and rendering users
    },
    resetAndCloseForm(): void {
      this.validationError = ""
      this.$emit("close")
    }
  }
})
</script>

<style lang="scss">
@import "../assets/modal.css";
</style>
