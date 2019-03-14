<template>
  <transition name="modal">
    <div class="modal-mask">
      <div class="modal-wrapper">
        <div class="modal-container">
          <div class="modal-header">
            <slot name="header">
              <h1>Modify user</h1>
            </slot>
          </div>

          <div class="modal-body">
            <slot name="body">
              <div class="field">
                <label class="label">Biername</label>
                <div class="control has-icons-left">
                  <input
                    class="input"
                    type="text"
                    placeholder="Biername"
                    v-model.lazy="user.BierName"
                  >
                  <span class="icon is-small is-left">
                    <font-awesome-icon icon="beer"/>
                  </span>
                </div>
              </div>

              <div class="field">
                <label class="label">First name</label>
                <div class="control has-icons-left">
                  <input
                    class="input"
                    type="text"
                    placeholder="First name"
                    v-model.lazy="user.FirstName"
                  >
                  <span class="icon is-small is-left">
                    <font-awesome-icon icon="user"/>
                  </span>
                </div>
              </div>

              <div class="field">
                <label class="label">Last name</label>
                <div class="control has-icons-left">
                  <input
                    class="input"
                    type="text"
                    placeholder="Last name"
                    v-model.lazy="user.LastName"
                  >
                  <span class="icon is-small is-left">
                    <font-awesome-icon icon="user"/>
                  </span>
                </div>
              </div>

              <div class="field">
                <label class="label">Email</label>
                <div class="control has-icons-left">
                  <input
                    class="input"
                    type="email"
                    placeholder="Email Address"
                    v-model.lazy="user.Email"
                  >
                  <span class="icon is-small is-left">
                    <font-awesome-icon icon="envelope"/>
                  </span>
                </div>
              </div>

              <div class="field">
                <label class="label">Phone Number</label>
                <div class="control has-icons-left">
                  <input
                    class="input"
                    type="text"
                    placeholder="Phone Number"
                    v-model.lazy="user.Phone"
                  >
                  <span class="icon is-small is-left">
                    <font-awesome-icon icon="phone"/>
                  </span>
                </div>
              </div>

              <div class="field">
                <label class="label">Status</label>
                <div class="control">
                <div class="select">
                <select v-model="selectedStatus">
                    <option
                      v-for="option in statusOptions"
                      :key="option"
                      :value="option"
                    >{{option}}</option>
                </select>
                </div>
                </div>
              </div>
            </slot>
          </div>

          <div class="field">
            <label class="label">Balance</label>
            <div class="control has-icons-left">
              <input class="input" type="text" placeholder="Balance" v-model.lazy="user.Balance">
              <span class="icon is-small is-left">
                <font-awesome-icon icon="euro-sign"/>
              </span>
            </div>
          </div>

          <div class="field">
            <label class="label">Max Debt</label>
            <div class="control has-icons-left">
              <input class="input" type="text" placeholder="Max Debt" v-model.lazy="user.MaxDebt">
              <span class="icon is-small is-left">
                <font-awesome-icon icon="money-bill"/>
              </span>
            </div>
          </div>

          <div class="modal-footer">
            <slot name="footer">
              <article v-if="validationError !==''" class="message is-danger">
                <div class="message-header">
                  <div class="field is-grouped">
                    <p class="icon is-small is-left">
                      <font-awesome-icon icon="exclamation" size="lg"/>
                    </p>
                    <p>{{validationError}}</p>
                  </div>
                </div>
              </article>
              <div class="field is-grouped">
                <div class="control">
                  <button class="button is-link" @click="modifyUser">Modify</button>
                </div>
                <div class="control">
                  <button class="button is-text" @click="cancelSubmission">Cancel</button>
                </div>
              </div>
            </slot>
          </div>
        </div>
      </div>
    </div>
  </transition>
</template>

<script src="./ModifyUserForm.js"></script>
<style lang="scss">
@import "../assets/modal.css";
</style>