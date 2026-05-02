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
        <input type="text" class="input" v-model="account.nickname">
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
            <option v-for="category in category$.accountCategories" :value="category.id" class="has-icons-left">
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
        <input type="number" class="input" v-model="maxDebt" step="1">
        <span class="icon is-small is-left">
          <icon :icon="['fas', 'usd']" />
        </span>
      </p>
    </div>
  </div>

  <hr class="divider"></hr>
  <div class="columns">
    <div class="column is-3">Show Advanced Settings</div>
    <div class="column">
      <ToggleSwitch v-model="showAdvanced"> </ToggleSwitch>
    </div>
  </div>

  <div class="columns" v-if="showAdvanced">
    <div class="column is-3" title="Useful if you want an account to show up in one category, but it should not be restricted by the categories visibilities.">Always show all Products</div>
    <div class="column" >
      <ToggleSwitch>Not yet implemented</ToggleSwitch>
    </div>
  </div>

  <hr class="divider" v-if="showAdvanced"></hr>
  <div class="columns">
    <div class="column is-3"></div>
    <div class="column">
      <Buttons>
        <Button :fa-icon="['fas', 'times']" icon-position="left" @click="$router.go(-1)">
          Cancel
        </Button>

        <Button class="is-primary" @click="actionButtonClicked" :fa-icon="['fas', 'save']" icon-position="right">
          {{ actionButton }}
        </Button>
      </Buttons>      
    </div>
  </div>

</template>

<style scoped>
.dropdown-item.is-active {
  background-color: transparent;
  color: inherit;
  font-weight: 600;
  border-left: 3px solid hsl(var(--bulma-primary-h), var(--bulma-primary-s), var(--bulma-primary-l));
  padding-left: calc(1rem - 3px);
}
.columns{
  align-items: center;
}
</style>

<script lang="ts">
import { useAccountStore } from '../../store/AccountStore';
import { useCategoryStore } from '../../store/CategoryStore';
import { Account } from '../../composables/account';
import Buttons from '../../composables/elements/Buttons.vue';
import Button from '../../composables/elements/Button.vue';
import { useSocketStore } from '../../store/socketStore';
import ToggleSwitch from '../../composables/elements/ToggleSwitch.vue';

export default {
  data() {
    return {
      account$: useAccountStore(),
      category$: useCategoryStore(),
      socket$: useSocketStore(),
      account: {} as Account,
      doneMounting: false,
      showAdvanced: false
    }
  },
  mounted() {
    if(this.isEdit){
      this.account = this.account$.byId(parseInt(this.$route.params.accountId.toString())).copy()
    } else {
      let newAccount = {} as Account;
      this.account = new Account(newAccount);
    }
    this.doneMounting = true;
  },
  components: {
    Buttons,
    Button,
    ToggleSwitch
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
    },
    maxDebt: {
      get() {
        return this.account.maxDebt / 100;
      },
      set(newValue: number) {
        this.account.maxDebt = newValue * 100;
      }
    }
  },
  methods: {
    actionButtonClicked(){
      if(this.isEdit){
        // Update current User
        // this.account$.byId(parseInt(this.$route.params.accountId.toString()))?.update(this.account);
        this.socket$.updateAccount(this.account);
        this.$router.push({name:'AccountSettings'});
        return;
      }
      this.account.maxDebt = Math.floor(this.account?.maxDebt);
      this.account.enabled = false;
      this.account.balance = 0;
      this.socket$.addAccount(this.account as Account);
      this.$router.push({name:'AccountSettings'});
    }
  }
}
</script>