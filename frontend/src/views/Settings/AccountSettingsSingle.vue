<template>
  <h1 class="title" v-if="doneMounting && isEdit">
    Edit Account:
    {{ account.getFullName() }}
  </h1>
  <h1 class="title" v-if="doneMounting && !isEdit">
    Neuen Account Erstellen
  </h1>
  <div class="columns" v-if="isEdit">
    <div class="column is-3">ID:</div>
    <div class="column">
      <input type="text" class="input" :value="account.id" disabled>
    </div>
  </div>

  <div class="columns">
    <div class="column is-3">First Name:</div>
    <div class="column">
      <p class="control has-icons-left">
        <input type="text" class="input" v-model="account.firstName">
        <span class="icon is-small is-left">
          <icon :icon="['fas', 'user']" />
        </span>
      </p>
    </div>
  </div>

  <div class="columns">
    <div class="column is-3">Last Name:</div>
    <div class="column">
      <p class="control has-icons-left">
        <input type="text" class="input" v-model="account.lastName">
        <span class="icon is-small is-left">
        <icon :icon="['fas', 'id-card']" />
        </span>
      </p>
    </div>
  </div>

  <div class="columns">
    <div class="column is-3">Nickname:</div>
    <div class="column">
      <p class="control has-icons-left">
        <input type="text" class="input" v-model="account.nickName">
        <span class="icon is-small is-left">
          <icon :icon="['fas', 'user-secret']" />
        </span>
      </p>
    </div>
  </div>

  <div class="columns">
    <div class="column is-3">E-Mail:</div>
    <div class="column">
      <p class="control has-icons-left">
        <input type="email" class="input" v-model="account.email">
        <span class="icon is-small is-left">
          <icon :icon="['fas', 'envelope']" />
        </span>
      </p>
    </div>
  </div>

  <div class="columns">
    <div class="column is-3">Phone Number:</div>
    <div class="column">
      <p class="control has-icons-left">
        <input type="tel" class="input" v-model="account.phone">
        <span class="icon is-small is-left">
          <icon :icon="['fas', 'phone']" />
        </span>
      </p>
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

  <div class="columns">
    <div class="column is-3">Max Debt Allowance:</div>
    <div class="column">
      <p class="control has-icons-left">
        <input type="number" class="input" v-model="account.maxDebt">
        <span class="icon is-small is-left">
          <icon :icon="['fas', 'usd']" />
        </span>
      </p>
    </div>
  </div>

  <div class="columns">
    <div class="column is-3"></div>
    <div class="column">
      <button class="button is-primary" @click="actionButtonClicked">{{ actionButton }}</button>
    </div>
  </div>

</template>

<script lang="ts">
import { useAccountStore } from '../../store/AccountStore';
import { useCategoryStore } from '../../store/CategoryStore';
import { Account } from '../../composables/account';

export default {
  data() {
    return {
      account$: useAccountStore(),
      category$: useCategoryStore(),
      account: {} as Account,
      doneMounting: false,
    }
  },
  mounted() {
    if(this.isEdit){
      this.account = this.account$.byId(parseInt(this.$route.params.accountId.toString()))
    } else {
      let newAccount = {} as Account;
      this.account = new Account(newAccount);
    }
    this.doneMounting = true;
  },
  computed: {
    categoryIcon(){
      return this.category$.byId(this.account.category)?.icon
    },
    isEdit(){
      return this.$route.params.accountId?.toString().length > 0;
    },
    actionButton(){
      return this.isEdit ? "Speichern" : "Neu Erstellen";
    }
  },
  methods: {
    actionButtonClicked(){
      if(this.isEdit){
        // Update current User
        return;
      }
      this.account.maxDebt = Math.floor(this.account.maxDebt);
      this.account.enabled = false;
      this.account.balance = 0;
      this.account.id = Math.ceil((1 + Math.random()) * 100 );
      this.account$.addAccount(this.account as Account);
      this.$router.push({name:'AccountSettings'})
    }
  }
}
</script>