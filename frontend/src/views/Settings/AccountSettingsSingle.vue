<template>
  <h1 class="title" v-if="doneMounting">Edit Account: {{ account.getFullName() }}</h1>
  <div class="columns">
    <div class="column is-3">ID:</div>
    <div class="column">
      <input type="text" class="input" :value="account.id" disabled>
    </div>
  </div>

  <div class="columns">
    <div class="column is-3">First Name:</div>
    <div class="column">
      <input type="text" class="input" v-model="account.firstName">
    </div>
  </div>

  <div class="columns">
    <div class="column is-3">Last Name:</div>
    <div class="column">
      <input type="text" class="input" v-model="account.lastName">
    </div>
  </div>

  <div class="columns">
    <div class="column is-3">Nickname:</div>
    <div class="column">
      <input type="text" class="input" v-model="account.nickName">
    </div>
  </div>

  <div class="columns">
    <div class="column is-3">E-Mail:</div>
    <div class="column">
      <input type="email" class="input" v-model="account.email">
    </div>
  </div>

  <div class="columns">
    <div class="column is-3">Phone Number:</div>
    <div class="column">
      <input type="tel" class="input" v-model="account.phone">
    </div>
  </div>


  <div class="columns">
    <div class="column is-3">Category:</div>
    <div class="column">
      <div class="control has-icons-left">
        <div class="select">
          <select v-model="account.category">
            <option v-for="category in category$.accountCategorys" :value="category.id" class="has-icons-left">
              {{ category.title }}
            </option>
          </select>
        </div>
        <div class="icon is-small is-left">
          <icon :icon="categoryIcon"/>
        </div>
      </div>
    </div>
  </div>

</template>

<script lang="ts">
import { useAccountStore } from '../../store/AccountStore';
import { useCategoryStore } from '../../store/CategoryStore';
import type { Account } from '../../composables/account';

export default {
  data() {
    return {
      account$: useAccountStore(),
      category$: useCategoryStore(),
      account: {} as Account,
      doneMounting: false
    }
  },
  mounted() {
    this.account = this.account$.byId(parseInt(this.$route.params.accountId.toString()))
    this.doneMounting = true;
  },
  computed: {
    categoryIcon(){
      return this.category$.byId(this.account.category)?.icon
    }
  },
  methods: {
  }
}
</script>