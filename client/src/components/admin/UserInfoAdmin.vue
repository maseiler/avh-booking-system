<template>
  <article class="message is-link" v-if="Object.keys(user).length !== 0">
    <div class="message-header">
      <p>{{displayUserName(user)}}</p>
    </div>
    <div class="message-body">
      <div class="columns">
        <div class="column has-text-left">
          <h6 class="subtitle is-6">
            <div>
              <font-awesome-icon icon="fingerprint" />
              &nbsp;{{user.ID}}
            </div>
            <div>
              <font-awesome-icon icon="beer" />
              &nbsp;{{user.BierName}}
            </div>
            <div>
              <font-awesome-icon icon="user" />
              &nbsp;{{user.FirstName}}
            </div>
            <div>
              <font-awesome-icon icon="user-circle" />
              &nbsp;{{user.LastName}}
            </div>
            <div>
              <font-awesome-icon icon="info-circle" />
              &nbsp;{{user.Status}}
            </div>
          </h6>
        </div>
        <div class="column has-text-left">
          <h6 class="subtitle is-6">
            <div>
              <font-awesome-icon icon="anchor" />
              &nbsp;{{user.BoatName}}
            </div>
            <div>
              <font-awesome-icon icon="envelope" />
              &nbsp;{{user.Email}}
            </div>
            <div>
              <font-awesome-icon icon="phone" />
              &nbsp;{{user.Phone}}
            </div>
            <div>
              &nbsp;
              <font-awesome-icon icon="euro-sign" />
              &nbsp;&nbsp;{{user.Balance}}&nbsp;€
            </div>
            <div>
              <font-awesome-icon icon="credit-card" />
              &nbsp;{{user.MaxDebt}}&nbsp;€
            </div>
          </h6>
        </div>
      </div>
    </div>
  </article>
</template>

<script>
export default {
  computed: {
    user() {
      return this.$store.state.selectedUser;
    }
  }
};
</script>