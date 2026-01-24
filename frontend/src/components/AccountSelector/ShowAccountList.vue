<template>
  <div class="table-container">
    <table class="table is-striped is-hoverable">
      <tbody>
        <tr>
          <th @click="sortFor('id')">
            ID
            {{ sortedTo == "id" ? "⯆" : "" }}
          </th>
          <th @click="sortFor('enabled')">
            Enabled
            {{ sortedTo == "enabled" ? "⯆" : "" }}
          </th>
          <th @click="sortFor('fn')">
            First Name
            {{ sortedTo == "fn" ? "⯆" : "" }}
          </th>
          <th @click="sortFor('nn')">
            Nickname
          {{ sortedTo == "nn" ? "⯆" : "" }}</th>
          <th @click="sortFor('ln')">
            Last Name
            {{ sortedTo == "ln" ? "⯆" : "" }}
          </th>
          <th @click="sortFor('mail')">
            E-Mail
          {{ sortedTo == "mail" ? "⯆" : "" }}
          </th>
          <th @click="sortFor('phone')">
            Phone
          {{ sortedTo == "phone" ? "⯆" : "" }}</th>
          <th @click="sortFor('bal')" class="has-text-right">
            Balance
          {{ sortedTo == "bal" ? "⯆" : "" }}</th>
          <th @click="sortFor('md')" class="has-text-right">
            MaxDebt
          {{ sortedTo == "md" ? "⯆" : "" }}</th>
          <th @click="sortFor('cat')">
            Category
          {{ sortedTo == "cat" ? "⯆" : "" }}</th>
          <th @click="sortFor('create')">
            Created At
          {{ sortedTo == "create" ? "⯆" : "" }}</th>
          <th v-show="hasEditAccountRights">Edit</th>
        </tr>
        <tr :class="account$.selected.includes(account) ? 'is-primary' : ''" v-for="account in accountsSorted" @click="account$.select(account)">
          <td>{{ account.id }}</td>
          <td><ToggleSwitch v-model="account.enabled" :disabled="false" /></td>
          <td>{{ account.firstName }}</td>
          <td>{{ account.nickname }}</td>
          <td>{{ account.lastName }}</td>
          <td class="has-copy-btn">{{ account.email }} <span class="icon is-small" @click="copyText(account.email)"><icon :icon="['fas', 'copy']" /></span></td>
          <td class="has-copy-btn">{{ account.phone }} <span class="icon is-small"><icon :icon="['fas', 'copy']" /></span></td>
          <td class="has-text-right">{{ $n(account.balance / 100, 'currency', 'de-DE') }}</td>
          <td class="has-text-right">{{ $n(account.maxDebt / 100, 'currency', 'de-DE') }}</td>
          <td>
            <button class="tag" :class="category$.byId(account.category) == undefined? 'is-skeleton' : ''">
              <span class="icon"><icon :icon="category$.byId(account.category)?.icon" /></span>
              <span>{{ category$.byId(account.category)?.title }}</span>
            </button>  
          </td>
          <td>{{ account.createdAt }}</td>
          <td v-show="hasEditAccountRights">
            <button class="button">
              <router-link :to="{ name: 'AccountSettingsSingle', params: { accountId: account.id } }">
                <span class="icon"><icon :icon="['fas', 'pen']"/></span>
              </router-link>
              </button>
          </td>
        </tr>

      </tbody>
    </table>
  </div>
</template>

<script lang="ts">
import type { Account } from '../../composables/account';
import { useAccountStore } from '../../store/AccountStore';
import type { PropType } from 'vue';
import { useCategoryStore } from '../../store/CategoryStore';
import ToggleSwitch from '../../composables/elements/ToggleSwitch.vue';

export default {
  data(){
    return {
      account$: useAccountStore(),
      category$: useCategoryStore(),
      dev: false,
      accountsSorted: [] as Account[],
      sortedTo: "",
    }
  },
  props: {
    accounts: {
      type: Array as PropType<Account[]>
    }
  },
  methods: {
    copyText(txt: string){
      navigator.clipboard.writeText(txt);
    },
    sortFor(sortParam: string){
      this.sortedTo = sortParam;
      switch (sortParam) {
        case("id"):{
          this.accountsSorted = this.accounts.sort((a, b) => {
            return a.id - b.id;
          }) as Account[];
          return;
        } 
        case("enabled"): {
            this.accountsSorted = this.accounts.sort((a, b) => {
            return (b.enabled ? 1 : 0) - (a.enabled ? 1 : 0);
          }) as Account[];
          return;
        }
        case("fn"): {
            this.accountsSorted = this.accounts.sort((a, b) => {
            return a.firstName.localeCompare(b.firstName);
          }) as Account[];
          return;
        }
        case("nn"): {
            this.accountsSorted = this.accounts.sort((a, b) => {
            return a.nickname.localeCompare(b.nickname);
          }) as Account[];
          return;
        }
        case("ln"): {
            this.accountsSorted = this.accounts.sort((a, b) => {
            return a.lastName.localeCompare(b.lastName);
          }) as Account[];
          return;
        }
        case("mail"): {
            this.accountsSorted = this.accounts.sort((a, b) => {
            return a.email.localeCompare(b.email);
          }) as Account[];
          return;
        }
        case("phone"): {
            this.accountsSorted = this.accounts.sort((a, b) => {
            return a.phone.localeCompare(b.phone);
          }) as Account[];
          return;
        }
        case("bal"): {
            this.accountsSorted = this.accounts.sort((a, b) => {
            return a.balance - b.balance;
          }) as Account[];
          return;
        }
        case("md"): {
            this.accountsSorted = this.accounts.sort((a, b) => {
            return a.maxDebt - b.maxDebt;
          }) as Account[];
          return;
        }
        case("cat"): {
            this.accountsSorted = this.accounts.sort((a, b) => {
            return a.category - b.category;
          }) as Account[];
          return;
        }
        case("create"): {
            this.accountsSorted = this.accounts.sort((a, b) => {
              let aDate = new Date(a.createdAt as string);
              let bDate = new Date(b.createdAt as string);
              if (aDate > bDate) { return -1} else { return 1};
          }) as Account[];
          return;
        }
      }
    }
  },
  components: {
    ToggleSwitch
  },
  computed: {
    hasEditAccountRights(){
      if(this.dev){
        console.warn("User has elevated privileges to edit products because you are running this in development environment")
        return true;
      }
      //ToDo Check if currently loged in user is allowed to edit products
      return false
    }
  },
  mounted() {
    this.dev = import.meta.env.DEV;
    this.accountsSorted = this.accounts as Account[];
  },
  watch: {
    accounts(newList, oldList){
      this.accountsSorted = newList;
      this.sortFor(this.sortedTo);
    }
  }
}
</script>

<style scoped>
  .has-copy-btn{
    position:relative;

    .icon {
      position:absolute;
      left:100%;
      z-index:10;
      padding:.8em;
      background-color:rgba(0,0,0,.3);
      border-radius:var(--bulma-radius);
      visibility:hidden;
      pointer-events: none;
      cursor:pointer;
    }
    &:hover .icon {
      visibility: visible;
      pointer-events: all;
    }
  }
  th{
    cursor: pointer;
    &:hover{
      background-color: rgba(255, 255, 255, .2);
    }
  }
  td{
    vertical-align: middle;
  }
</style>