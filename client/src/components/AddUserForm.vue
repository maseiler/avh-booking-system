<template>
  <div>
    <!-- TODO: do as modal (~> pop up) -->
    <!-- <div class="modal">
    <div class="modal-background"></div>
    <div class="modal-card">
      <header class="modal-card-head">
        <p class="modal-card-title">Modal title</p>
        <button class="delete" aria-label="close"></button>
      </header>
    <section class="modal-card-body">-->
    <div class="field">
      <label class="label">Biername</label>
      <div class="control has-icons-left">
        <input class="input" type="text" placeholder="Biername" v-model.lazy="newUser.bierName">
        <span class="icon is-small is-left">
          <font-awesome-icon icon="beer"/>
        </span>
      </div>
    </div>

    <div class="field">
      <label class="label">First name</label>
      <div class="control has-icons-left">
        <input class="input" type="text" placeholder="First name" v-model.lazy="newUser.firstName">
        <span class="icon is-small is-left">
          <font-awesome-icon icon="user"/>
        </span>
    </div>

    <div class="field">
      <label class="label">Last name</label>
      <div class="control has-icons-left">
        <input class="input" type="text" placeholder="Last name" v-model.lazy="newUser.lastName">
        <span class="icon is-small is-left">
          <font-awesome-icon icon="user"/>
        </span>
    </div>

    <div class="field">
      <label class="label">Email</label>
      <div class="control has-icons-left">
        <input class="input" type="email" placeholder="Email Address" v-model.lazy="newUser.email">
        <span class="icon is-small is-left">
          <font-awesome-icon icon="envelope"/>
        </span>
    </div>
    <div class="field">
      <label class="label">Phone Number</label>
      <div class="control has-icons-left">
        <input class="input" type="text" placeholder="Phone Number" v-model.lazy="newUser.phone">
        <span class="icon is-small is-left">
          <font-awesome-icon icon="phone"/>
        </span>
    </div>

    <div class="field">
      <label class="label">Status</label>
      <div class="control">
        <div class="select">
          <select v-model="newUser.status">
            <option disabled value="">Status</option>
            <option>Aktiv B</option>
            <option>Aktiv KA</option>
            <option>AH</option>
            <option>Gast</option>
          </select>
        </div>
      </div>
    </div>

    <div class="field is-grouped">
          <div class="control">
            <button class="button is-link" @click="submitUser">Submit</button>
          </div>
          <div class="control">
            <button class="button is-text" @click="cancelSubmission">Cancel</button>
          </div>
        </div>
      <!-- </section>
      <footer class="modal-card-foot">
        <button class="button is-success">Save changes</button>
        <button class="button">Cancel</button>
      </footer>
    </div>
    </div> -->
  </div>
</template>

 <script src="./AddUserForm.js"></script>