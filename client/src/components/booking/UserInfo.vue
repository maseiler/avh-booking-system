<template>
  <article class="message is-link" v-if="Object.keys(user).length !== 0">
    <div class="message-header">
      <p>{{displayUserName(user)}}</p>
    </div>
    <div class="message-body">
      <!-- TODO: favicon -->
      <div class="columns">
        <div class="column has-text-left">
          <h6 class="subtitle is-6">
            <div v-if="user.BierName !== ''">
              <font-awesome-icon icon="beer"/>
              &nbsp;Biername: {{user.BierName}}
            </div>
            <div v-if="user.FirstName !== ''">
              <font-awesome-icon icon="user"/>
              &nbsp;Firstname: {{user.FirstName}}
            </div>
            <div v-if="user.LastName !== ''">
              <font-awesome-icon icon="user"/>
              &nbsp;Lastname: {{user.LastName}}
            </div>
            <div v-if="user.Status !== ''">
              <font-awesome-icon icon="info-circle"/>
              &nbsp;Status: {{user.Status}}
            </div>
            <div v-if="user.Email !== ''">
              <font-awesome-icon icon="envelope"/>
              &nbsp;Email: {{user.Email}}
            </div>
            <div v-if="user.Phone !== ''">
              <font-awesome-icon icon="phone"/>
              &nbsp;Phone: {{user.Phone}}
            </div>
          </h6>
        </div>
        <div class="column has-text-centered">
          <h5 class="subtitle is-5" v-if="user.Balance !== ''">Balance:</h5>
          <h6 class="title is-5">{{user.Balance}} €</h6>
        </div>
      </div>
    </div>
  </article>
</template>

<script>
export default {
  props: {
    user: {}
  }
};
</script>